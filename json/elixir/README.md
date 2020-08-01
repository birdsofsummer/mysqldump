# Mysqldump

**TODO: Add description**

## Installation

```bash

mix deps.get

iex -S mix

config= [
      name: :myapp_db,
      hostname: "127.0.0.1",
      username: "root",
      database: "test",
      password: "123456",
      timeout: 5000,
      # ssl_opts: [versions: [:"tlsv1.2"]],
      backoff_type: :stop,
      max_restarts: 0,
      pool_size: 1,
      show_sensitive_data_on_connection_error: true
 ]

{:ok,pid}  = MyXQL.start_link(config)
{:ok,r}    = MyXQL.query(pid, "show databases;")
#  MyXQL.child_spec/1
#  MyXQL.Client
#  MyXQL.close/2
#  MyXQL.close/3
#  MyXQL.Connection
#  MyXQL.Cursor
#  MyXQL.Error
#  MyXQL.execute!/3
#  MyXQL.execute/3
#  MyXQL.execute!/4
#  MyXQL.execute/4
#  MyXQL.json_library/0
#  MyXQL.prepare!/3
#  MyXQL.prepare/3
#  MyXQL.prepare!/4
#  MyXQL.prepare/4
#  MyXQL.prepare_execute!/3
#  MyXQL.prepare_execute/3
#  MyXQL.prepare_execute!/4
#  MyXQL.prepare_execute/4
#  MyXQL.prepare_execute!/5
#  MyXQL.prepare_execute/5
#  MyXQL.Protocol
#  MyXQL.Query
#  MyXQL.query!/2
#  MyXQL.query/2
#  MyXQL.query!/3
#  MyXQL.query/3
#  MyXQL.query!/4
#  MyXQL.query/4
#  MyXQL.Result
#  MyXQL.rollback/2
#  MyXQL.start_link/1
#  MyXQL.stream/2
#  MyXQL.stream/3
#  MyXQL.stream/4
#  MyXQL.TextQuery
#  MyXQL.transaction/2
#  MyXQL.transaction/3

```


```elixir
    #https://github.com/elixir-ecto/myxql
    # config :myxql, :json_library, Jason

    # if using supervision tree
        defmodule MyApp.Application do
          use Application

          def start(_type, _args) do
            children = [
              {MyXQL, username: "root", password: "bleepbloop", hostname: "localhost", database: "myapp", name: :myapp_db}
            ]

            Supervisor.start_link(children, opts)
          end
        end


    config= [
          hostname: "127.0.0.1",
          username: "root",
          database: "test",
          password: "123456",
          timeout: 5000,
          # ssl_opts: [versions: [:"tlsv1.2"]],
          backoff_type: :stop,
          max_restarts: 0,
          pool_size: 1,
          show_sensitive_data_on_connection_error: true
     ]

    alias MyXQL.{Client, Protocol}
    import MyXQL.Protocol.{Flags, Records}
    {:ok, %MyXQL.Client{connection_id: connection_id, sock: {:gen_tcp, port}}} = Client.connect(config)


    {:ok,pid}  = MyXQL.start_link(config)
    {:ok,pid}  = MyXQL.start_link(username: "root")
    {:ok,pid}  = MyXQL.start_link(username: "root", database: "blog")
    {:ok,pid}  = MyXQL.start_link(username: "root", password: "123456", hostname: "localhost", database: "test")


    sql=fn x ->MyXQL.query(pid,x) end
    sql.("select * from user")
    "select * from user" |> sql.()


    tb="""
        CREATE TABLE IF NOT EXISTS posts
        (
           `id` INT UNSIGNED AUTO_INCREMENT,
           `title` VARCHAR(100) NOT NULL,
           `author` VARCHAR(40) NOT NULL,
           `date` DATE,
           PRIMARY KEY ( `id` )
        )ENGINE=InnoDB DEFAULT CHARSET=utf8;

    """

    r=MyXQL.query!(pid, "show databases;")
    r=MyXQL.query!(pid, "CREATE DATABASE IF NOT EXISTS blog")
    r=MyXQL.query!(pid, "drop database blog")

    r=MyXQL.query!(pid, tb)
    r=MyXQL.query!(pid, "show tables;")
    r=MyXQL.query!(pid, "INSERT INTO posts (`title`) VALUES ('Post 1')")
    r=MyXQL.query!(pid, "INSERT INTO posts (`title`) VALUES (?), (?)", ["Post 2", "Post 3"])
    r=MyXQL.query!(pid, "SELECT * FROM posts")
    r=MyXQL.query!(pid, "drop table posts;")

    r=MyXQL.query!(pid, "SELECT NOW()").rows

     ## MyXQL.query(pid, "use mysql;")
     ##   {:error,
     ##    %MyXQL.Error{
     ##      connection_id: 4,
     ##      message: "(1295) (ER_UNSUPPORTED_PS) This command is not supported in the prepared statement protocol yet",
     ##      mysql: %{code: 1295, name: :ER_UNSUPPORTED_PS},
     ##      statement: "use test;"
     ##    }}


```

## json


```elixir

Jason.encode!(%{"age" => 44, "name" => "Steve Irwin", "nationality" => "Australian"})
Jason.decode!(~s({"age":44,"name":"Steve Irwin","nationality":"Australian"}))

# When called directly:
    plug Absinthe.Plug,
      schema: MyApp.Schema,
      json_codec: Jason

# When used in phoenix router:
forward "/api",
  to: Absinthe.Plug,
  init_opts: [schema: MyApp.Schema, json_codec: Jason]

```


If [available in Hex](https://hex.pm/docs/publish), the package can be installed
by adding `mysqldump` to your list of dependencies in `mix.exs`:

```elixir
def deps do
  [
    {:mysqldump, "~> 0.1.0"}
  ]
end
```

Documentation can be generated with [ExDoc](https://github.com/elixir-lang/ex_doc)
and published on [HexDocs](https://hexdocs.pm). Once published, the docs can
be found at [https://hexdocs.pm/mysqldump](https://hexdocs.pm/mysqldump).

