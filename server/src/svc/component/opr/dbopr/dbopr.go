package dbopr

import (
	"core/dbfw/mysql"
	"core/netfw"
	"errors"
	"fmt"
	"log"
	"protos/coredef"
	"strings"
)

var dbInfo []*coredef.OperateDB

var dbMap map[string]dbfw.IDBExecutor

var dbMapInfo map[string]string = map[string]string{
	"db_td2cn":    "td2game:td2mysqlpass@tcp(192.168.1.245:3306)/db_td2cn_dev2",
	"db_td2ios":   "td2game:td2mysqlpass@tcp(192.168.1.245:3306)/db_td2ios_dev2",
	"db_td2wp":    "td2game:td2mysqlpass@tcp(192.168.1.245:3306)/db_td2wp_dev2",
	"db_icloud":   "td2game:td2mysqlpass@tcp(192.168.1.245:3306)/db_icloud_dev2",
	"db_discount": "td2game:td2mysqlpass@tcp(192.168.1.245:3306)/db_td2cn_statistics_dev2",
}

func IterateShowDB(callback func(name, showname string)) {
	for i := 0; i < len(dbInfo); i++ {
		if dbInfo[i].GetShowName() != "" {
			callback(dbInfo[i].GetName(), dbInfo[i].GetShowName())
		}
	}
}

func ConnectDB() {

	dbMap = make(map[string]dbfw.IDBExecutor)

	var config coredef.OperateConfig
	netfw.GetConfig("OperateConfig", &config)

	for _, v := range config.GetDB() {
		dbExec := dbfw.NewDBExecutor()
		dbConfig := netfw.SvcConfig.GetDB()
		log.Println("connect:", v.GetAddr())
		err := dbExec.Open(v.GetAddr(), dbConfig.GetShowOperate(), int(dbConfig.GetConnCount()))
		if err != nil {
			log.Println(err)
		}
		dbMap[v.GetName()] = dbExec
		dbInfo = append(dbInfo, v)
	}

}

func CreateDBExecutor(db string, call func(dbfw.IDBExecutor)) error {

	dbExec, ok := dbMap[db]
	if !ok {
		log.Println("find db failed:", db)
		return errors.New("no such " + db + " db server")
	}

	call(dbExec)

	return nil

}

func ExecuteSQL(db string, s interface{}, cmd string, args ...interface{}) (res []interface{}, err error) {

	erx := CreateDBExecutor(db, func(dbExec dbfw.IDBExecutor) {

		switch strings.ToLower(cmd[:6]) {
		case "select":
			res, err = dbExec.Query(s, cmd, args...)
			log.Println("query length:", len(res), "err:", err)
		case "update":
			err = dbExec.Update(s, cmd, args...)
		case "insert":
			err = dbExec.Insert(s, cmd, args...)
		default:
			err = dbExec.Exec(s, cmd, args...)
		}

	})

	if erx != nil {
		res, err = nil, erx
	}

	return

}

type tableExist struct {
	Table_name string `db:"q"`
}

func IsTableExist(db string, table string) bool {

	res, err := ExecuteSQL(db, &tableExist{}, fmt.Sprintf("select $FIELD_NAME$ from information_schema.tables where table_name='%s'", table))
	if err != nil {
		log.Println(err)
		return false
	}

	return len(res) != 0

}

var PageRecordCount int32 = 100

type tableCount struct {
	Count int32 `db:"q"`
}

func GetRecordCount(db, tb string) int32 {

	res, err := ExecuteSQL(db, &tableCount{}, fmt.Sprintf("select count(*) as count from %s", tb))
	if err != nil {
		return 0
	}

	return res[0].(tableCount).Count

}

func GetRecordPageCount(db, tb string) int32 {

	count := GetRecordCount(db, tb)

	if count%PageRecordCount == 0 {
		return count / PageRecordCount
	}

	return count/PageRecordCount + 1

}
