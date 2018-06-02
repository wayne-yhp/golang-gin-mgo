package utils

import "fmt"
import "github.com/zheng-ji/goSnowFlake"
import "errors"

/*
	uid生成类
 */
type Uid struct {
	iw *goSnowFlake.IdWorker
}

func Newuid() *Uid{
	u := new(Uid)
	err := errors.New("")			//错误信息
	u.iw,err = goSnowFlake.NewIdWorker(1)	//唯一标志操作对象初始化
	if err != nil {
		fmt.Println(err)
	}
	return u
}

//获取新的一个uuid
func (u *Uid)Getid()  int64{
	id ,_:= u.iw.NextId()
	return id
}
