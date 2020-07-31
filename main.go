package main

import (
	. "fmt"
	. "./cmd"
	. "./cos"
	"time"
)

var (
	HOST="localhost"
	PWD="123456"
)

func now() string{
	t:=time.Now().Unix()
	return Sprintf("%d", t)
}

func Backup(){
	t:=now()
	name:="/tmp/"+t+".sql"
	err,_:=Dump_sql(HOST,PWD,name)
	if err!=nil{
		Printf("dump fail",err)
		return
	}
	Upload1(name,name)
	err,_=Exec1("rm",[] string {name},"/tmp/1.log")
	if err!=nil{
		Printf("del fail",err)
	}
}

func main() {
	Backup()
}
