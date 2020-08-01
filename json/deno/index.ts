// curl -fsSL https://deno.land/x/install/install.sh | sh
// deno upgrade
// deno run --allow-net --allow-run --allow-read --allow-write  --allow-all --allow-env --allow-hrtime --allow-plugin --unstable index.ts
// deno run --allow-all --unstable index.ts

console.log("deno run --allow-net --allow-run --allow-read --allow-write --unstable  index.ts ")

//https://deno.land/manual/getting_started/permissions
//https://github.com/denoland/deno_std

//https://github.com/denoland/deno_std/tree/master/fs
import { 
    ensureDir, 
//  ensureDirSync,
//  emptyDir,
//  emptyDirSync ,
//  ensureFile,
//  ensureFileSync,
//  walk, 
//  walkSync,
//  ensureSymlink,
//  ensureSymlinkSync,
//  exists,
//  existsSync,
//  format,
//  detect,
//  EOL,
//  globToRegExp,
//  copy,
//  copySync,
//  move,
//  moveSync,
//  readFileStr, 
//  readFileStrSync,
//  writeFileStr,
//  writeFileStrSync
//  readJson, 
//  readJsonSync,
//  writeJson, 
//  writeJsonSync,
} from "https://deno.land/std/fs/mod.ts";



//https://github.com/eveningkid/denodb
import { DataTypes, Database, Model } from 'https://deno.land/x/denodb/mod.ts';

import { Client } from "https://deno.land/x/mysql/mod.ts";

interface MysqlDbConfig {
    hostname: string;
    username?: string;
    password?: string;
    db?:       string;
}

interface DbConfig{
    host:     string;
    username?: string;
    password?: string;
    database?: string;
}


//"yyyymmdd"
const today=():string=>{
    let t=new Date()
    return t.toJSON().split('T')[0].replace(/-/g,'')
}

const write_json=async (file_name:string="",d: object = {})=>{
/*
    const status = await Deno.permissions.query({ name: "write" });
    if (status.state !== "granted") {
          throw new Error("need write permission");
    }
    await Deno.permissions.revoke({ name: "read" });
    await Deno.permissions.revoke({ name: "write" });
*/
    console.log('save',file_name)
    const f = await Deno.open(file_name, "a+");
    const s=JSON.stringify(d,null,'\t')
    const encoder = new TextEncoder();
    await f.write(encoder.encode(s));
    //await Deno.remove(file_name);
    f.close();
}


const export_json = async (config:MysqlDbConfig) => {
    const s1='show databases;'
    const s2='show tables;'

    const client = await new Client();

    client.connect(config);
    let dbs =  await client.execute(s1)        //?
    console.log("zzzz",dbs)

    for (let db of dbs){
        const file_path=`/tmp/db/${today()}/${config.hostname}/${db}`

        await ensureDir(file_path)

        let s3=`use ${db}`
        await client.execute(s3)
        const tables=await client.execute(s2)  //?
        for (let t of tables) {
            let s4=`select * from ${t}`
            let d=await client.execute(s4)
            const file_name=`file_path/${t}.json`
            console.log('dddd',d)
            console.log("save",db,t,file_name)
            //write_json(file_name,d)
        }
        console.log("export ",db)
    }
    console.log("all done")
}

class User extends Model {
  static table = 'users';
  static timestamps = true;
  static fields = {
    id: {
      primaryKey: true,
      autoIncrement: true,
    },
    name: DataTypes.STRING,
    email: {
      type: DataTypes.STRING,
      unique: true,
      allowNull: false,
      length: 50,
    },
  };
}


class Flight extends Model {
  static table = 'flights';
  static timestamps = true;

  static fields = {
    id: { primaryKey: true, autoIncrement: true },
    departure: DataTypes.STRING,
    destination: DataTypes.STRING,
    flightDuration: DataTypes.FLOAT,
  };

  static defaults = {
    flightDuration: 2.5,
  };
}


// db.query(q: QueryDescription)
// https://github.com/eveningkid/denodb/blob/master/lib/query-builder.ts

const export_json1=async (config:DbConfig)=>{
    let host=config.host

    //"postgres" | "sqlite3" | "mysql" | "mongo";
    const db = new Database('mysql', config); 
    await db.ping()


    let s1="show databases;"
    let s2="show tables;"

    let dbs=await db.query(s1)         //?

    for (let i of dbs) {
        let file_path=`/tmp/db/${today()}/${host}/${i}`
        await ensureDir(file_path)

        await db.query(`use ${i}`);    //?

        let tables=await db.query(s2); //?
        for (let t of tables){
        let file_name=`${file_path}/${t}.json`

           //"SELECT * FROM `${t}`"
           let s={
               table:t;
               select: "*",
           }
           let d=await db.query(s)
           console.log("write json ",host,i,t,file_name)
           write_json(file_name,d)
        }
    }
    //----------------------------------------crud-------------------------------------------
    //db.link([User,Flight]);
    //await db.sync({ drop: true });

    // const flight = new Flight();
    // flight.departure = 'London';
    // flight.destination = 'San Francisco';
    // await flight.save();
    // await Flight.select('destination').all();
    // await Flight.where('destination', 'Tokyo').delete();
    // const sfFlight = await Flight.select('destination').find(2);
    // await Flight.count();
    // await Flight.select('id', 'destination').orderBy('id').get();
    // await sfFlight.delete();

    //await User.create({ name: 'Amelia' });
    //await User.all();
    //await User.deleteById('1');
    //----------------------------------------crud-------------------------------------------

    console.log("all done,bye")
    await db.close();
}


const test=()=>{
  const config:MysqlDbConfig={
    "hostname": "127.0.0.1",
    "username": "root",
    "password": "123456",
    "db": "",
    }
  export_json(config)
}

const test1=()=>{
    const config:DbConfig={
      "host": "127.0.0.1",
      "username": "root",
      "password": "123456",
      "database": "test",
    }
    export_json1(config)
}


test()
test1()
