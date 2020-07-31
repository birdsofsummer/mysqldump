package cmd

import (
	"os/exec"
	. "fmt"
	"log"
	"io/ioutil"
)


func Exec1( app string, s []string, name string)(error,string){
    //Print(app,s)
    c:= exec.Command(app,s...)
	stdout, err := c.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return err,""
	}

	if err := c.Start(); err != nil {
		log.Fatal(err)
		return err,""
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
		return err,""
	}
	err = ioutil.WriteFile(name, bytes, 0644)
	if err != nil {
		panic(err)
		return err,""
	}
	//Print("1111111",string(bytes))
    Print("save to ",name,"\n")
	return nil,name
}


func Dump_sql(
	host string,
	pwd string,
	name string,
)(error,string){
    app := "mysqldump"
    s :=[] string {
		"-A",
		"--opt", 
		"-h"+host,
		"-p"+pwd,
		"-uroot", 
		"--default-character-set=utf8",
	}
	return Exec1(app,s,name)
}





