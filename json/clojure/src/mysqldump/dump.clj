(ns mysqldump.dump
    (:require 
      [clj-http.client :as client]
      [http.async.client :as http]
      [cheshire.core :refer :all]
      [somnium.congomongo :as m]
      [clojure.string :as string]
      [clojure.data.json :as json]
      [cheshire.core :as json1]
      [clj-http.conn-mgr :as conn-mgr]
      [net.cgrand.enlive-html :as html]
    ; [cljs.js :as cljs.js]
    ; [clojure.java.io :as io]
    ; [me.raynes.fs :refer :all]
    ; [me.raynes.fs.compression :refer :all]
   )  
  (:refer-clojure
   :exclude [compile take drop sort distinct conj! disj! resultset-seq case])

  (:use [clojure.java.io])

  (:use [clojureql.internal :only [update-or-insert-vals update-vals]]
        [clojure.java.jdbc ] ;:only [with-connection find-connection]
        clojureql.core
        [clojure.java.io :only [delete-file]]
        ;clojure.contrib.mock
        )
 ;  (:import (javax.sql DataSource))
   (:import java.io.File)
   (:import [java.net URL])
   (:import java.util.Date)
   (:import java.io.Closeable
           java.sql.Connection)

   ;(import (org.apache.http.entity.mime HttpMultipartMode))
 
  
  
  )


;(md "/tmp/a/b/c")
(defn md [p] 
    (println "mkdir " p)
    (def a (string/split p #"/" ))
    (def n (count a))
    (doseq [x (range 1 (+ 1 n))] 
      (def f (string/join "/" (take x a)))
      (println f)
            (try
                (.mkdir (java.io.File. f)) 
                (catch Exception e (str "fail to mkdir " (.getMessage e))))

      )
  )


(defn now [] 
   (.getTime (java.util.Date.)))

(defn today []
  (.format (java.text.SimpleDateFormat. "yyyyMMdd") (System/currentTimeMillis))
)



;;;mysql;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(defn mysql? []
  (isa? (class (connection)) com.mysql.jdbc.JDBC4Connection))

; https://github.com/LauJensen/clojureql/blob/master/test/test.clj
(def mysql
  {
   :classname   "com.mysql.jdbc.Driver"
   :subprotocol "mysql"
   :user        "root"
   :password    "123456"
   :auto-commit true
   :fetch-size  1500
   :subname     "//localhost:3306/test"
 })

(def c (open-global mysql))

(defn drop-if [table]
  (try (drop-table table) (catch Exception _)))


(defn drop-schema []
  ;(drop-if :salary)
  (drop-if :user)
  )



;;;;;;;;;;;;;;;;;;;;;;;;crud;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;rs (table :user)
;rs (select (table :user) (where (= :id 5)))
;rs (project (table :user) [:name])
;rs (sort (table :user) [:id :name])
;rs (table :user)] (println rs) rs))
;rs (project (table :user) [:name])
;rs (clojureql.core/sort (table :user) [:id#desc :name#asc])

;select table_name from information_schema.tables where table_schema='test' and table_type='base table';

;SELECT GROUP_CONCAT(COLUMN_NAME SEPARATOR ",") 
;FROM information_schema.COLUMNS 
;WHERE 
;TABLE_SCHEMA = 'test' 
;AND 
;TABLE_NAME = 'user'

;  :update_time 
;  :table_comment 
;  :avg_row_length 
;  :table_collation 
;  :table_type 
;  :table_schema 
;  :row_format 
;  :auto_increment 
;  :data_length 
;  :table_rows 
;  :check_time 
;  :checksum 
;  :index_length 
;  :create_time 
;  :create_options 
;  :engine 
;  :table_catalog 
;  :max_data_length 
;  :version 
;  :table_name 
;  :data_free

; (query-tables "test")
(defn query-tables [db]
  (with-connection mysql 
  (with-results 
    [
     rs (-> (table :information_schema.tables) 
            (select 
               (where (
                       and 
                       (= :table_schema db) 
                       (= :table_type "base table")
                       ) 
                ) 
             )
            ) 
     ] 
    (println rs) 
    rs
    )
   )
)


(defn query-tables1 [[db-name config]]
  (with-connection config
  (with-results 
    [
     rs (-> (table :information_schema.tables) 
            (select 
               (where (
                       and 
                       (= :table_schema db-name) 
                       (= :table_type "base table")
                       ) 
                ) 
             )
            ) 
     ] 
    (println rs) 
    rs
    )
   )
)


(defn query1 [db-name config {table-name :table_name}]
  (with-connection config 
    (with-results [rs (-> 
                        (table table-name) 
                        ;(clojureql.core/sort [:id#desc :name#asc]) 
                        ;(clojureql.core/take 50)
                        ;(clojureql.core/drop 2)
                        )] 
      ;(doseq [r rs] (println r))
      (println rs)
      rs
      )
   )
)

(defn query [table-name]
  (with-connection mysql 
    (with-results [rs (-> 
                        (table table-name) 
                        ;(clojureql.core/sort [:id#desc :name#asc]) 
                        ;(clojureql.core/take 50)
                        ;(clojureql.core/drop 2)
                        )] 
      ;(doseq [r rs] (println r))
      (println rs)
      rs
      )
   )
)

(defn export-db [db] 
    ( 
     ->>
      (query-tables db)
      (map :table_name )
      (map (fn [x] {x (query x)}))
      (into {})
))

(defn  write-json1 [[file_name data]] 
  (
     let [ 
           p "/tmp/db/"
           n (str p file_name ".json") 
     ]
    (spit  n (json1/encode data {:pretty true}) :append false)      
))

(defn find-db-host [config]
  (-> 
    (:subname config)
    (string/split #":" ) 
    first 
    (string/replace "//" "")
    )
  )

(defn find-db-name [config] 
  (-> 
    (:subname config) 
    (string/split #"/") 
    last
    (string/replace  #"\?.*" "" )
    ))


; (export-dbs [mysql])
(defn export-dbs [dbs] 
      (def configs  (->> (map (fn [x] [ (find-db-name x) x]) dbs) (into {})))
      (def all-tables 
           (map (fn [[db-name config]]
                 (def tables (query-tables1 [db-name config ]))
                 [db-name config tables]   
            )
            configs
           ))
      (map  (fn [[db-name config tables]] 
                 (map (fn [x]  
                        (
                         let [ 
                               host-name (find-db-host config)
                               table-name (:table_name x)
                               p (str "/tmp/db/" host-name "/" (today) "/" db-name  "/") 
                               file-name (str p table-name ".json") 
                         ]
                        (println db-name table-name p file-name)
                        (md p)
                        (def table-data (query1 db-name config x))
                            (println "save to " file-name)
                            (spit  file-name (json1/encode table-data {:pretty true}) :append false)      
                         )
                        ) tables)
                 ) all-tables)
)


(defn export-test []
    (map write-json1 (export-db "test"))
)

(defn query2 [table-name]
  (with-connection mysql 
    (with-results [
          rs (-> (table table-name) (select (where (> :id 5))))
       ] 
      ;(doseq [r rs] (println r))
      (println rs)
      rs
      )
   )
)

(defn update-user [name]
  (with-connection mysql 
    (with-results [
                   rs (update! (table :user) (where (> :id 1)) {:name name})
                  ;rs (update-in! (table :user) (where (> :id 1)) {:name "Bob"})
                   ] 
      (println rs)
      ;(doseq [r rs] (println r))
      rs
      )
   )
)

(defn add-user [name,n]
  (
   let [d (for [x (range n)] {:id x :name name})]
  (with-connection mysql 
    (with-results [rs (clojureql.core/conj! (table :user) d)] 
      (println rs)
      ;(doseq [r rs] (println r))
      rs
   ))
))


(defn del-user []
  (with-connection mysql 
    (with-results [rs (clojureql.core/disj! (table :user) (where (> :id -10)))] 
      ;(doseq [r rs] (println r))
      (println rs)
      rs
    )
   )
)




;;;;;;;;;;;;;;;;;;;;;;;;crud;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;
;(def databases [mysql postgresql sqlite3 derby])
;(def postgresql
;  {:classname   "org.postgresql.Driver"
;   :subprotocol "postgresql"
;   :user        "cql"
;   :password    "cql"
;   :auto-commit true
;   :fetch-size  500
;   :subname     "//localhost:5432/cql"})

;(def sqlite3
;  {:classname   "org.sqlite.JDBC"
;   :subprotocol "sqlite"
;   :subname     "/tmp/cql.sqlite3"
;   :create      true})


;;;mysql;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;




