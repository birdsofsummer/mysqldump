const u=require('utility')
const R=require('ramda')
const mysql= require('mysql');

// tails([1,2,3,4]) =>  [ [ 1 ], [ 1, 2 ], [ 1, 2, 3 ], [ 1, 2, 3, 4 ] ]
const tails=(p=[])=>p.map((x,i)=>p.slice(0,i+1))
const md=(p=[],prefix="/tmp/")=>tails(p)
    .map(x=>prefix+x.join('/'))
    .map(x=>{
        try{
            fs.mkdirSync(x)
        }catch(e){
           // console.log(e)
        }
    })

const export_json=(config={})=>{
   var connection = mysql.createConnection(config);
   const host=connection.config.host
   const today=u.YYYYMMDD()
   connection.connect();
   connection.query('show databases;' , function (error, results, fields) {
       let dbs = results
       for (let {Database} of dbs) {

           let file_path=["db",today,host,Database,]
           md(file_path)
           console.log('wait to export ',Database,file_path )

           connection.query(`use ${Database};`, function(a,b,c) {
                connection.query('show tables;',function (a1,b1,c1){
                   let t=b1.map(Object.values).flat()
                   for (let i of t) {
                       connection.query(`select * from ${i};`,(a2,b2,c2)=>{
                          file_name=R.prepend('/tmp',file_path).join('/')+"/"+i+".json"
                          u.writeJSON(file_name,b2)
                       })
                   }
                })
           })
       }
    })
}


test1=()=>{
    config={
          host     : '127.0.0.1',
          user     : 'root',
          password : '123456',
          //database : 'test' , //
    }
    export_json(config)
}

test2=()=>{
    config={
          host     : '127.0.0.1',
          user     : 'root',
          password : '123456',
          database : 'test' , //
    }

    var connection = mysql.createConnection(config);
    connection.query('SELECT 1 + 1 AS solution', function (error, results, fields) {
      if (error) throw error;
      console.log(results);
    });

    //connection.query('INSERT INTO user(Id,name) VALUES(0,?)',['root',], function (error, results, fields) {})
    //connection.query('UPDATE user SET name = ? WHERE Id = ?',['root11',16], function (error, results, fields) {})
    //connection.query("DELETE FROM user where id=16",function (err, result) {})
    connection.query('show tables;', function (error, results, fields) {
      if (error) throw error;
      console.log(results);
        for (let i of results) {
                let t=i.Tables_in_test
                console.log("sss",t)

                let tt=`SELECT
                GROUP_CONCAT(COLUMN_NAME SEPARATOR ",") as a
                FROM information_schema.COLUMNS
                WHERE TABLE_SCHEMA = 'test'
                AND
                TABLE_NAME = '${t}'
                `

                let t1=`SELECT
                COLUMN_NAME as name
                FROM information_schema.COLUMNS
                WHERE TABLE_SCHEMA = 'test'
                AND
                TABLE_NAME = '${t}'
                `
                connection.query(t1, function (error, results, fields) {
                  if (error) throw error;
                  console.log('zzz',results.map(x=>x.name));
                });
                connection.query('SELECT * from ' + t, function (error, results, fields) {
                      if (error) throw error;
                      console.log(results);
                });
        }
    });
}


module.exports={
    export_json,
}
