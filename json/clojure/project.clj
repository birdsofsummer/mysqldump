(defproject mysqldump "0.1.0-SNAPSHOT"
  :description "FIXME: write description"
  :url "http://example.com/FIXME"
  :license {:name "EPL-2.0 OR GPL-2.0-or-later WITH Classpath-exception-2.0"
            :url "https://www.eclipse.org/legal/epl-2.0/"}
  :dependencies [
                 [org.clojure/clojure "1.10.1"]
                 [clj-http "3.10.1"]
                 [http.async.client "1.3.1"]
                 [cheshire "5.9.0"]
                 [congomongo "1.1.0"]
                 [me.raynes/fs "1.4.6"]
                 [org.clojure/data.json "1.0.0"]
                 [org.apache.logging.log4j/log4j-api "2.11.0"]
                 [org.apache.logging.log4j/log4j-core "2.11.0"]
                 [org.apache.logging.log4j/log4j-1.2-api "2.11.0"]
                 [enlive "1.1.6"]
                 [org.clojure/clojurescript "1.10.773"]
                 [clojureql "1.0.5"]
                 ;[mysql/mysql-connector-java "8.0.21"]
                 ;[c3p0/c3p0 "0.9.5.5"]

                 ]
  :repl-options {:init-ns mysqldump.core}
  :main mysqldump.core
  :plugins [[lein-codox "0.10.7"]]
  )
