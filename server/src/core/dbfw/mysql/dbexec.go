package dbfw

import (
	"errors"
	"fmt"
	"github.com/ziutek/mymysql/autorc"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
	"log"
	"os"
	"reflect"
	"time"
	//_ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

type IDBExecutor interface {
	// username:password@tcp(address:port)/dbname
	Open(dsn string, showOperate bool, connCount int) error

	Close()

	// 查询, cmd中可以添加$FIELD_NAME$, 根据s的字段, 合成语句,  struct tag包含q
	Query(s interface{}, cmd string, args ...interface{}) ([]interface{}, error)

	// 执行, cmd中可以添加$FIELD_NAME$ $FIELD_VALUE$ $FIELD_EQUN$, 根据s的字段, 合成语句, struct tag包含e
	Exec(s interface{}, cmd string, args ...interface{}) error

	// 更新, cmd中可以添加$FIELD_EQUN$, 根据s的字段, 合成语句, struct tag包含u
	Update(s interface{}, cmd string, args ...interface{}) error

	// 插入, cmd中可以添加$FIELD_NAME$ $FIELD_VALUE$, 根据s的字段, 合成语句, struct tag包含i
	Insert(s interface{}, cmd string, args ...interface{}) error
}

type dbExecutor struct {
	showOperate bool // 显示所有操作日志

	connChan chan *autorc.Conn
}

func NewDBExecutor() IDBExecutor {
	return &dbExecutor{
		showOperate: true,
	}
}

// username:password@tcp(address:port)/dbname
func (self *dbExecutor) Open(dsn string, showOperate bool, connCount int) error {
	self.showOperate = showOperate

	self.connChan = make(chan *autorc.Conn, connCount)

	var conn *autorc.Conn

	for i := 0; i < connCount; i++ {

		if usr, passwd, nettype, addr, dbname, err := parseDSN(dsn); err == nil {

			//func New(proto, laddr, raddr, user, passwd string, db ...string) *Conn {
			conn = autorc.New(nettype, "", addr, usr, passwd, dbname)
		} else {
			return err
		}

		self.connChan <- conn
	}

	return nil
}

func (self *dbExecutor) backConn(conn *autorc.Conn) {
	self.connChan <- conn
}

func (self *dbExecutor) fetchConn() *autorc.Conn {

	return <-self.connChan
}

func (self *dbExecutor) Close() {

	for i := 0; i < cap(self.connChan); i++ {
		conn := <-self.connChan
		conn.Raw.Close()
	}

}

func fillStruct(row mysql.Row, sample reflect.Value, rowIndex int) {

	rowSize := len(row)
	// 遍历输入的结构体, 将每个类型的值缓冲地址存到scanparam中
	for i := 0; i < sample.NumField(); i++ {

		v := sample.Field(i)

		if i >= rowSize {
			log.Printf("out of row size, %d, rowsize: %d", i, rowSize)
			continue
		}

		//log.Println(i, rowIndex, row.Str(rowIndex))

		switch v.Kind() {
		case reflect.Int32, reflect.Uint32, reflect.Int64:
			v.SetInt(row.ForceInt64(rowIndex))
		case reflect.String:
			v.SetString(row.Str(rowIndex))
		case reflect.Struct:
			fillStruct(row, v, rowIndex)
		default:
			log.Printf("dbfw.makeScanParam unsupport type: %s", v.Type())
		}

		rowIndex++

	}
}

// 查询
func (self *dbExecutor) Query(s interface{}, cmd string, args ...interface{}) ([]interface{}, error) {

	if s == nil {
		return nil, errors.New("struct can not be nil")
	}

	sqlcmd := reflectCompose(s, cmd, Token_FIELD_NAME, "q", useFieldName)

	self.DBLog("#DBQuery: %s", fmt.Sprintf(sqlcmd, args...))

	dbConn := self.fetchConn()

	defer self.backConn(dbConn)

	rows, res, err := dbConn.Query(sqlcmd, args...)

	if err != nil {
		return nil, err
	}

	defer eatResult(res)

	vSample := reflect.ValueOf(s).Elem()

	outBuffer := make([]interface{}, len(rows))

	for col, row := range rows {

		if err != nil {
			log.Printf("%s", err)
			continue
		}

		fillStruct(row, vSample, 0)

		outBuffer[col] = vSample.Interface()

	}

	return outBuffer, nil
}

// 执行, 无返回, s只做格式化cmd用
func (self *dbExecutor) Exec(s interface{}, cmd string, args ...interface{}) error {

	sqlcmd := reflectCompose(s, cmd, Token_FIELD_NAME, "e", useFieldName)
	sqlcmd = reflectCompose(s, sqlcmd, Token_FIELD_VALUE, "e", useFieldValue)
	sqlcmd = reflectCompose(s, sqlcmd, Token_FIELD_EQUN, "e", useEquationField)

	dbConn := self.fetchConn()

	defer self.backConn(dbConn)

	logstr := fmt.Sprintf(sqlcmd, args...)
	self.DBLog("#DBExec: %s", logstr)
	logFile(logstr)
	_, _, err := dbConn.Query(sqlcmd, args...)

	if err != nil {
		return err
	}

	return nil
}

// 更新, s格式化输出为等式
func (self *dbExecutor) Update(s interface{}, cmd string, args ...interface{}) error {

	sqlcmd := reflectCompose(s, cmd, Token_FIELD_EQUN, "u", useEquationField)

	dbConn := self.fetchConn()

	defer self.backConn(dbConn)

	logstr := fmt.Sprintf(sqlcmd, args...)
	self.DBLog("#DBUpdate: %s", logstr)
	logFile(logstr)
	_, _, err := dbConn.Query(sqlcmd, args...)

	if err != nil {
		return err
	}

	return nil
}

// 插入, s格式化输出为等式
func (self *dbExecutor) Insert(s interface{}, cmd string, args ...interface{}) error {

	sqlcmd := reflectCompose(s, cmd, Token_FIELD_NAME, "i", useFieldName)
	sqlcmd = reflectCompose(s, sqlcmd, Token_FIELD_VALUE, "i", useFieldValue)

	dbConn := self.fetchConn()

	defer self.backConn(dbConn)

	logstr := fmt.Sprintf(sqlcmd, args...)
	self.DBLog("#DBInsert: %s", logstr)
	logFile(logstr)
	_, _, err := dbConn.Query(sqlcmd, args...)

	if err != nil {
		return err
	}

	return nil
}

func (self *dbExecutor) DBLog(format string, v ...interface{}) {
	if self.showOperate {
		log.Printf(format, v...)
	}
}

func eatResult(res mysql.Result) {
	for {
		res, err := res.NextResult()

		if err != nil {
			log.Printf("%s", err)
			continue
		}

		if res == nil {
			break
		}

		if res.StatusOnly() {
			break
		}
	}
}

func logFile(str string) {
	f, e := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND, 0777)
	if e != nil {
		log.Println(e)
	}
	defer f.Close()
	f.WriteString(time.Now().Format("2006-01-02 15:04:05") + "	" + str + "\n")
}
