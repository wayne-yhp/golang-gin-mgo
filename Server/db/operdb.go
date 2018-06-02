package db

import (
	"gopkg.in/mgo.v2"
	"sync"
	"Server/utils"
)

/*
数据库操作类
*/
var obj *Operdb
type Operdb struct {
	mgo_session *mgo.Session	//连接对象
	mgo_db *mgo.Database		//数据库对象
	mgo_c *mgo.Collection		//集合对象
	bulk *mgo.Bulk				//批处理对象
	rwlock *sync.RWMutex			//读写锁
}

func init()  {
	infomap :=utils.Resolve()
	obj = new(Operdb)
	obj.rwlock = new(sync.RWMutex)
	obj.connect(infomap)
}

func Newoperdb() *Operdb {
	return obj
}

//数据库连接,对象初始化
func (oper *Operdb)connect(info map[string]string){
	 ip  := info["ip"]
	 port  := info["port"]
	url := ip + ":" + port
	oper.mgo_session, _ = mgo.Dial(url)
	oper.mgo_db = oper.mgo_session.DB(info["databasename"])
	oper.mgo_db.Login(info["username"], info["password"])
	oper.mgo_c = oper.mgo_db.C(info["colname"])
	oper.bulk = oper.mgo_c.Bulk()
}

//关闭连接
func (oper *Operdb)close()  {
	oper.mgo_session.Close()
}

//批量插入
func (oper *Operdb)BulkInset(arg interface{})  {
	oper.bulk.Insert(arg)		//批量记录
}

//批量执行
func (oper *Operdb)BulkRun()  {
	oper.rwlock.Lock()			//写锁
	oper.bulk.Run()				//批量插入
	oper.bulk = oper.mgo_c.Bulk()
	oper.rwlock.Unlock()
}