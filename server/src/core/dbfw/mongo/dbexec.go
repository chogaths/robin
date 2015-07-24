package dbfw

import (
	//"errors"
	//"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"log"
)

type IDBExecutor interface {
	// url = mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
	Open(url string, showOperate bool, connCount int) error

	Close()

	// 查询一条，result为struct地址
	FindOne(collName string, query interface{}, result interface{}) error

	// 查询多条，result为切片
	FindAll(collName string, result interface{}) error

	// 插入
	Insert(collName string, docs ...interface{}) error

	// 更新
	Update(collName string, selector interface{}, update interface{}) error
}

type dbExecutor struct {
	showOperate bool // 显示所有操作日志

	sesChan chan *mgo.Session
}

func NewDBExecutor() IDBExecutor {
	return &dbExecutor{
		showOperate: true,
	}
}

// username:password@tcp(address:port)/dbname
func (self *dbExecutor) Open(url string, showOperate bool, connCount int) error {
	self.showOperate = showOperate

	// 到admin库验证账号
	url = url + "?authSource=admin"

	self.sesChan = make(chan *mgo.Session, connCount)

	for i := 0; i < connCount; i++ {

		ses, err := mgo.Dial(url)
		if err != nil {
			return err
		}

		ses.SetMode(mgo.Monotonic, true)

		self.sesChan <- ses
	}

	return nil
}

func (self *dbExecutor) Insert(collName string, docs ...interface{}) error {

	ses := self.fetchSes()
	defer self.backSes(ses)

	return ses.DB("").C(collName).Insert(docs)

}

func (self *dbExecutor) Update(collName string, selector interface{}, update interface{}) error {
	ses := self.fetchSes()
	defer self.backSes(ses)

	return ses.DB("").C(collName).Update(selector, update)

}

func (self *dbExecutor) FindOne(collName string, query interface{}, result interface{}) error {
	ses := self.fetchSes()
	defer self.backSes(ses)

	return ses.DB("").C(collName).Find(query).One(result)

}

func (self *dbExecutor) FindAll(collName string, result interface{}) error {
	ses := self.fetchSes()
	defer self.backSes(ses)

	return ses.DB("").C(collName).Find(bson.M{}).All(result)
}

func (self *dbExecutor) backSes(ses *mgo.Session) {
	self.sesChan <- ses
}

func (self *dbExecutor) fetchSes() *mgo.Session {

	return <-self.sesChan
}

func (self *dbExecutor) Close() {

	for i := 0; i < cap(self.sesChan); i++ {
		ses := <-self.sesChan
		ses.Close()
	}

}
