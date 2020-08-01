(ns mysqldump.core
     (:require 
   ;   [clj-http.client :as client]
   ;   [http.async.client :as http]
   ;   [cheshire.core :refer :all]
   ;   [somnium.congomongo :as m]
   ;   [clojure.string :as string]
   ;   [clojure.data.json :as json]
   ;   [cheshire.core :as json1]
   ;   [clj-http.conn-mgr :as conn-mgr]
   ;   [net.cgrand.enlive-html :as html]
      [mysqldump.dump :refer [export-dbs]]
    ; [cljs.js :as cljs.js]
    ; [clojure.java.io :as io]
    ; [me.raynes.fs :refer :all]
    ; [me.raynes.fs.compression :refer :all]
   )  
  )

(defn test1
  []
  (export-dbs [
     {
       :classname   "com.mysql.jdbc.Driver"
       :subprotocol "mysql"
       :user        "root"
       :password    "123456"
       :auto-commit true
       :fetch-size  1500
       :subname     "//localhost:3306/test"
     }
  ])
  )


(defn -main
  "dddd"
  [& args]
  (test1)
   )
