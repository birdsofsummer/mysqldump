package main
 
import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
	"encoding/json"
	"time"
	"os"
)

var (
    userName  string = "root"
    password  string = "123456"
    ipAddrees string = "127.0.0.1"
    port      int    = 3306
    dbName    string = "test"
    charset   string = "utf8"
)
 


type TableInfo struct {
	ColumnName    string `json:"columnName"`
	ColumnType    string `json:"columnType"`
	MaxLength     string `json:"maxLength"`
	ColumnComment string `json:"columnComment"`
	NullAble      string `json:"nullAble"`
}

type TableInfos []TableInfo


func Md(path string){
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println("exist")
	}else{
		err=os.MkdirAll(path,os.ModePerm)
		if err!=nil{
		   fmt.Println(err)
		   return
	    }
		fmt.Println("created")
	}
}

func today()(string){
	now := time.Now()
	n:=now.Format("2006-01-02")
	return n
}

func write_table(file_name string,d TableInfos){
	b, _ := json.MarshalIndent(d, "", "\t")
	file,_:=os.Create(file_name)
	defer file.Close()
	_, err := file.Write(b)
	if err!=nil{
		fmt.Println("[saved] error")
	}
	fmt.Println("[saved]:",file_name)
}


// db="test" 
// select table_name tableName from information_schema.tables where table_schema='%s'"
// table="user"
// SELECT t.column_name ColumnName,t.data_type ColumnType,t.character_maximum_length MaxLength,t.column_comment columnComment,t.IS_NULLABLE NullAble FROM information_schema.COLUMNS t where t.TABLE_SCHEMA="test" and TABLE_NAME = "user";

func show_table(Db *sqlx.DB,db string,table string)(TableInfos ,error){
	s2:=fmt.Sprintf(`SELECT t.column_name ColumnName,t.data_type ColumnType,t.character_maximum_length MaxLength,t.column_comment columnComment,t.IS_NULLABLE NullAble FROM information_schema.COLUMNS t where t.TABLE_SCHEMA="%s" and TABLE_NAME = "%s"`,db,table)
	rows, err := Db.Query(s2)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	//[]string{"属性名称", "类型", "最大长度", "允许为空", "备注"}
	var (
		ColumnName    string
		ColumnType    string
		MaxLength     string
		ColumnComment string
		NullAble      string
	)
	var tableInfos TableInfos 
	for rows.Next() {
		rows.Scan(&ColumnName, &ColumnType, &MaxLength, &ColumnComment, &NullAble)
		tableInfos = append(tableInfos, TableInfo{ColumnName: ColumnName, MaxLength: MaxLength, ColumnType: ColumnType, ColumnComment: ColumnComment, NullAble: NullAble})
	}
	return tableInfos, nil
}

func show_db(Db *sqlx.DB) (error){
	query1:=func (s string) (error,[]string){
		rows, err := Db.Query(s)
		defer rows.Close()
        var d []string
		if err != nil {
			return err,d
		}
		for rows.Next() {
			var d1 string
			rows.Scan(&d1)
			d=append(d,d1)
		}
		return err,d
	}

	s0:=`select distinct table_schema from information_schema.tables where table_type="BASE TABLE";`
    _,dbs:=query1(s0)
	fmt.Println(dbs)
    exclude:=map[string]string{"mysql" : "","performance_schema":""} //忽略mysql,performance_schema
	for _,db:=range(dbs) {
		_,ok :=exclude[db]
		if ok {
			continue
		}
		s1:=fmt.Sprintf(`select table_name tableName from information_schema.tables where table_schema='%s'`,db)
		_,tables:=query1(s1)
		fmt.Println(tables)

		file_path:=fmt.Sprintf(`/tmp/db/%s/%s`,today(),db)
		Md(file_path)

		for _,table:=range(tables) {
			tableInfos, err := show_table(Db,db,table)
			file_name:=fmt.Sprintf(`%s/%s.json`,file_path,table)
			write_table(file_name,tableInfos)
			if err != nil {
				fmt.Print("query tableInfo error,", err)
				continue
			}
			
			//fmt.Println(db,table,tableInfos)
			fmt.Println(db,table,file_name)
		}
	}

	fmt.Println(dbs,"done")
	return nil
}




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


/*
	ping(Db)
	queryData(Db)
	deleteRecord(Db)
	addRecord(Db)
	updateRecord(Db)
	queryData(Db)
	find(Db)
	findMany(Db)
*/
	show_db(Db)
}
