# mysqldump

- mysqldump -> /tmp/xxx.sql -> cos



## env

`/etc/profile`

```bash
export TENCENTCLOUD_SECRET_ID="xxx"
export TENCENTCLOUD_SECRET_KEY="xxx"
```








## bash

```bash

host=127.0.0.1
password=123456
now=`date "+%Y%m%d"`
time=` date +%Y_%m_%d_%H_%M_%S `
mkdir $now

export_sql(){
    s=`mysql -h$host -p$password -e 'show databases' -E |grep -v 'row' |awk '{print $2}'`
    echo $s
    for i in $s
    do
        echo $i
        mysqldump -h$host -uroot -p$password --default-character-set=utf8 -B $i --skip-lock-tables >$now/$i.sql
        #mysqldump $i | gzip > $now/$i_$time.sql.gz
        #mysqldump $i | xz > $now/$i_$time.sql.xz
        #mysqldump -uroot -p123456  --lock-all-tables --flush-logs $i
        #mysql -uroot -pmysql $i < $i.sql
    done
}

export_all(){
    mysqldump -A -d -h$host -uroot -p$password --default-character-set=utf8  >$now/all.sql
}

export_sql
#export_all
echo "$now done!"

```

