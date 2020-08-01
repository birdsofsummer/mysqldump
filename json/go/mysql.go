package main
 
import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var (
    userName  string = "root"
    password  string = "123456"
    ipAddrees string = "127.0.0.1"
    port      int    = 3306
    dbName    string = "test"
    charset   string = "utf8"
)
 
/*
go get "github.com/go-sql-driver/mysql"
go get "github.com/jmoiron/sqlx"

create database test;
CREATE TABLE IF NOT EXISTS `user`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `name` VARCHAR(100) NOT NULL,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO user ( name ) VALUES ( "tt" );
SELECT * FROM user;

use mysql; 
grant all privileges on `test`.* to 'root'@'localhost';
ALTER USER 'root'@'localhost' IDENTIFIED BY '123456';
update mysql.user set authentication_string = password("123456") where user="root";
flush privileges;

SELECT user, host, plugin FROM mysql.user;
UPDATE mysql.user SET plugin = '' WHERE plugin = 'unix_socket';
FLUSH PRIVILEGES;
SELECT user, host, plugin FROM mysql.user;







*/

type User struct {
	Id int `db:"id"`
	Name string `db:"name"`
}


func connectMysql() (*sqlx.DB) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
    Db, err := sqlx.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
    }
    return Db
}

func ping(Db *sqlx.DB) {
    err := Db.Ping()
    if err != nil {
        fmt.Println("ping failed")
    } else {
        fmt.Println("ping success")
    }
}
 
func find(Db *sqlx.DB) {
	var u *User = new(User)
    err := Db.Get(u,"select * from user where id = 10")
    if err != nil {
        fmt.Printf("query faied, error:[%v]", err.Error())
        return
    }
    fmt.Println("ddddd",u)
}

func findMany(Db *sqlx.DB) {
    var u []User
    err := Db.Select(&u,"select * from user")
    if err != nil {
        fmt.Printf("query faied, error:[%v]", err.Error())
        return
    }
    for _, uu := range u {
        fmt.Println("mmmmmmmmm",uu)
    }
}

func queryData(Db *sqlx.DB) {
    rows, err := Db.Query("select * from user")
    if err != nil {
        fmt.Printf("query faied, error:[%v]", err.Error())
        return
    }
    for rows.Next() {
        var id int
        var name string
        err := rows.Scan(&id, &name)
        if err != nil {
            fmt.Println("get data failed, error:[%v]", err.Error())
        }
        fmt.Println(id,name)
    }
    rows.Close()
}




func deleteRecord(Db *sqlx.DB){
    result, err := Db.Exec("delete from user") // where id = 2
    if err != nil {
        fmt.Printf("delete faied, error:[%v]", err.Error())
        return
    }
    num, _ := result.RowsAffected()
    fmt.Printf("delete success, affected rows:[%d]\n", num)
}
 

func addRecord(Db *sqlx.DB) {
    for i:=0; i<20; i++ {
        //result, err := Db.Exec("insert into user values(?,?)",i, "root")
        result, err := Db.Exec("insert into user ( name )  values(?)", "root")
        if err != nil {
            fmt.Printf("data insert faied, error:[%v]", err.Error())
            return
        }
        id, _ := result.LastInsertId()
        fmt.Printf("insert success, last id:[%d]\n", id)
    }
}


func updateRecord(Db *sqlx.DB){
    result, err := Db.Exec("update user set name = 'anson' where id > 15") // 
    if err != nil {
        fmt.Printf("update faied, error:[%v]", err.Error())
        return
    }
    num, _ := result.RowsAffected()
    fmt.Printf("update success, affected rows:[%d]\n", num)
}

func main(){
    var Db *sqlx.DB = connectMysql()
	Db.SetMaxOpenConns(100)
	//DB.SetConnMaxLifetime(d time.Duration)
    defer Db.Close()



	ping(Db)
	queryData(Db)
	deleteRecord(Db)
	addRecord(Db)
	updateRecord(Db)
	queryData(Db)
	find(Db)
	findMany(Db)

}
