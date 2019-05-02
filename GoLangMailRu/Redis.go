package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	c redis.Conn
)


func getRecord (mkey string ) string {
	println("get", mkey)

	item,err:=redis.String(c.Do("GET",mkey))
	if err==redis.ErrNil {
		fmt.Println("Record not found in redis (return value is nil)")
		return ""
	} else if err!=nil{
		PanicOnErr(err)
	}
}

func main(){
	var err error

	c,err=redis.DialURL("redis://user:@localhost:6379/0")
	PanicOnErr(err)
	defer c.Close()

	mkey:="record_21"


}