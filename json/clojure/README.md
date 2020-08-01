# mysqldump

## mongodb

```clojure
  [somnium.congomongo :as m]
;    m/add-index!
;    m/aggregate
;    m/calculate-query-options
;    m/close-connection
;    m/collection-exists?
;    m/collections
;    m/command
;    m/connection?
;    m/create-collection!
;    m/databases
;    m/db-ref
;    m/db-ref?
;    m/destroy!
;    m/destroy-file!
;    m/distinct-values
;    m/drop-all-indexes!
;    m/drop-coll!
;    m/drop-database!
;    m/drop-index!
;    m/fetch
;    m/fetch-and-modify
;    m/fetch-by-id
;    m/fetch-by-ids
;    m/fetch-count
;    m/fetch-files
;    m/fetch-one
;    m/fetch-one-file
;    m/get-coll
;    m/get-collection-read-preference
;    m/get-collection-write-concern
;    m/get-db
;    m/get-gridfs
;    m/get-indexes
;    m/get-timestamp
;    m/group
;    m/insert!
;    m/insert-file!
;    m/make-connection
;    m/map-reduce
;    m/mass-insert!
;    m/mongo-options
;    m/named
;    m/object-id
;    m/query-option-map
;    m/read-preference
;    m/server-eval
;    m/set-collection-read-preference!
;    m/set-collection-write-concern!
;    m/set-connection!
;    m/set-database!
;    m/set-read-preference
;    m/set-write-concern
;    m/stream-from
;    m/StringNamed
;    m/update!
;    m/with-db
;    m/with-mongo
;    m/with-ref-fetching
;    m/write-concern-map
;    m/write-file-to

```


## mysql

[clojureql](https://clojureql.sabrecms.com/en/documentation)


```clojure
(:use clojureql.core)
;  clojureql.core/aggregate
;  clojureql.core/apply-on
;  clojureql.core/build-join
;  clojureql.core/case
;  clojureql.core/close-global
;  clojureql.core/compile
;  clojureql.core/conj!
;  clojureql.core/*debug*
;  clojureql.core/declare-tables
;  clojureql.core/difference
;  clojureql.core/disj!
;  clojureql.core/distinct
;  clojureql.core/drop
;  clojureql.core/global-connections
;  clojureql.core/grouped
;  clojureql.core/interpolate-sql
;  clojureql.core/intersection
;  clojureql.core/join
;  clojureql.core/limit
;  clojureql.core/map->RTable
;  clojureql.core/modify
;  clojureql.core/offset
;  clojureql.core/open-global
;  clojureql.core/order-by
;  clojureql.core/outer-join
;  clojureql.core/pick
;  clojureql.core/predicate-symbols
;  clojureql.core/project
;  clojureql.core/Relation
;  clojureql.core/rename
;  clojureql.core/->RTable
;  clojureql.core/select
;  clojureql.core/select-if
;  clojureql.core/sort
;  clojureql.core/table
;  clojureql.core/table?
;  clojureql.core/take
;  clojureql.core/transform
;  clojureql.core/union
;  clojureql.core/update!
;  clojureql.core/update-in!
;  clojureql.core/where
;  clojureql.core/with-cnx
;  clojureql.core/with-cnx*
;  clojureql.core/with-results


```

## jdbc

```clojure
;jdbc


[clojure.java.jdbc ] ;:only [with-connection find-connection]


;  clojure.java.jdbc/as-identifier
;  clojure.java.jdbc/as-key
;  clojure.java.jdbc/as-keyword
;  clojure.java.jdbc/as-named-identifier
;  clojure.java.jdbc/as-named-keyword
;  clojure.java.jdbc/as-quoted-identifier
;  clojure.java.jdbc/as-quoted-str
;  clojure.java.jdbc/as-str
;  clojure.java.jdbc/connection
;  clojure.java.jdbc/create-table
;  clojure.java.jdbc/create-table-ddl
;  clojure.java.jdbc/delete-rows
;  clojure.java.jdbc/do-commands
;  clojure.java.jdbc/do-prepared
;  clojure.java.jdbc/do-prepared-return-keys
;  clojure.java.jdbc/drop-table
;  clojure.java.jdbc/find-connection
;  clojure.java.jdbc/insert-record
;  clojure.java.jdbc/insert-records
;  clojure.java.jdbc/insert-rows
;  clojure.java.jdbc/insert-values
;  clojure.java.jdbc/is-rollback-only
;  clojure.java.jdbc/prepare-statement
;  clojure.java.jdbc/print-sql-exception
;  clojure.java.jdbc/print-sql-exception-chain
;  clojure.java.jdbc/print-update-counts
;  clojure.java.jdbc/resultset-seq
;  clojure.java.jdbc/set-rollback-only
;  clojure.java.jdbc/transaction
;  clojure.java.jdbc/transaction*
;  clojure.java.jdbc/update-or-insert-values
;  clojure.java.jdbc/update-values
;  clojure.java.jdbc/with-connection
;  clojure.java.jdbc/with-connection*
;  clojure.java.jdbc/with-naming-strategy
;  clojure.java.jdbc/with-query-results
;  clojure.java.jdbc/with-query-results*
;  clojure.java.jdbc/with-quoted-identifiers

```


## json

```clojure
;  clojure.data.json/json-str
;  clojure.data.json/JSONWriter
;  clojure.data.json/pprint
;  clojure.data.json/pprint-json
;  clojure.data.json/print-json
;  clojure.data.json/read
;  clojure.data.json/read-json
;  clojure.data.json/read-str
;  clojure.data.json/write
;  clojure.data.json/-write
;  clojure.data.json/write-json
;  clojure.data.json/write-str
;  
;  
;  
;  cheshire.core/copy-arglists
;  cheshire.core/create-generator
;  cheshire.core/create-pretty-printer
;  cheshire.core/decode
;  cheshire.core/decode-smile
;  cheshire.core/decode-stream
;  cheshire.core/decode-strict
;  cheshire.core/default-pretty-print-options
;  cheshire.core/encode
;  cheshire.core/encode-smile
;  cheshire.core/encode-stream
;  cheshire.core/eof
;  cheshire.core/generate-cbor
;  cheshire.core/generate-smile
;  cheshire.core/generate-stream
;  cheshire.core/generate-string
;  cheshire.core/*generator*
;  cheshire.core/*opt-map*
;  cheshire.core/parse-cbor
;  cheshire.core/parsed-seq
;  cheshire.core/parsed-smile-seq
;  cheshire.core/parse-smile
;  cheshire.core/parse-stream
;  cheshire.core/parse-stream-strict
;  cheshire.core/parse-string
;  cheshire.core/parse-string-strict
;  cheshire.core/with-writer
;  cheshire.core/write
```

