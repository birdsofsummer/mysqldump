"""
pip3 install pymysql

"""


import pymysql
import json
import datetime
import os
import decimal



class MyEncoder(json.JSONEncoder):
    def default(self, obj):
        if isinstance(obj, datetime.datetime):
            return obj.strftime('%Y-%m-%d %H:%M:%S')
        elif isinstance(obj, datetime.date):
            return obj.strftime("%Y-%m-%d")
        elif isinstance(obj, decimal.Decimal):
            return float(obj)
        elif isinstance(obj,bytes):
            return obj.__str__()
            #return obj.decode("utf-8")
        else:
            return json.JSONEncoder.default(self, obj)


def write_json(file_name,data={}):
    s=""
    try:
        s=json.dumps(data, ensure_ascii=False,indent=3,cls=MyEncoder)
    except:
       print(file_name,'bad json?')
    with open(file_name, 'w+') as f:
        f.write(s)
    f.close()


def md(file_path="/tmp/a/b/c"):
    print("md",file_path)
    try:
       os.makedirs(file_path)
    except:
       print(file_path,'exist?')

def today():
    t=datetime.date.today()
    return t.strftime('%Y%m%d')


def export_json(config,exclude=["information_schema","mysql","performance_schema"]):
    s1="show databases;"
    s2="show tables;"
    host=config["host"]

    conn=pymysql.connect(**config)

    cursor=conn.cursor()
    rows=cursor.execute(s1)
    dbs=[x[0] for x in cursor.fetchall()]
    cursor.close()
    print("dbs:",dbs)
    for db in dbs:
        if db in exclude:
            print("pass",db)
            continue
        print("export ",host,db)
        path_name="/tmp/db/{}/{}/{}".format(today(),host,db)
        md(path_name)

        cursor=conn.cursor()
        rows=cursor.execute("use "+db)
        cursor.close()

        cursor=conn.cursor()
        rows=cursor.execute(s2)
        tables=[x[0] for x in cursor.fetchall()]
        cursor.close()

        for table in tables:
            file_name="{}/{}.json".format(path_name,table)
            print("export ",host,db,table,file_name)

            sql = 'select * from {}'.format(table)
            cursor=conn.cursor()
            rows=cursor.execute(sql)
            d=cursor.fetchall()
            k=[x[0] for x in cursor.description]
            cursor.close()

            d1=[dict (zip( k, x)) for x in d]
            write_json(file_name,d1)
    conn.close()
    print("done!")


def test():
    config={
        "host":"127.0.0.1",
        "port":3306,
        "user":"root",
        "password":"123456",
        "db":"test",
        "charset":"utf8"
    }
    #export_json(config,[])  #export all
    export_json(config)


test()














'''
py2

The setup.py script uses mysql_config to find all compiler and linker
options, and should work as is on any POSIX-like platform, so long as
mysql_config is in your path.

mysql_config=/usr/local/bin/mysql_config
ls $mysql_config

apt install libmariadbd18
apt install libmariadbd-dev
apt install mysql-community-client

pip install MySQL-python
https://pypi.org/project/MySQL-python/


https://sourceforge.net/projects/mysql-python/files/mysql-python/1.2.3/MySQL-python-1.2.3.tar.gz/download
https://kumisystems.dl.sourceforge.net/project/mysql-python/mysql-python/1.2.3/MySQL-python-1.2.3.tar.gz
file=MySQL-python-1.2.3.tar.gz
tar -xvf $file
cd $file
python setup.py build
python setup.py install



import MySQLdb

db = MySQLdb.connect("localhost", "root", "123456", "test", charset='utf8' )
cursor = db.cursor()
cursor.execute("SELECT VERSION()")
data = cursor.fetchone()
print(data)
db.close()


'''


