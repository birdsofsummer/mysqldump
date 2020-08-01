# Mysqldump

**TODO: Add description**

## Installation

```bash

mix deps.get

iex


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


      alias MyXQL.{Client, Protocol}
      import MyXQL.Protocol.{Flags, Records}

      def opts() do
        [
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
      end

      o = [username: "default_auth", password: "secret"] ++ opts()
      Client.connect(o)


    {:ok,:myapp_db} = MyXQL.start_link(username: "root", password: "123456", hostname: "localhost", database: "test")

    "select * from user" |> DB.query(:myapp_db)
    "select * from user" |> DB.paginate |> DB.query(:myapp_db)
    "select * from user" |> DB.paginate(1) |> DB.query(:myapp_db)
    "select * from user where id = ?" |> DB.query(:myapp_db, [1]) |> hd


     {:ok, pid} = MyXQL.start_link(username: "root")
     {:ok, pid} = MyXQL.start_link(username: "root", database: "blog")

     MyXQL.query!(pid, "CREATE DATABASE IF NOT EXISTS blog")

     MyXQL.query!(pid, "INSERT INTO posts (`title`) VALUES ('Post 1')")
     MyXQL.query(pid, "INSERT INTO posts (`title`) VALUES (?), (?)", ["Post 2", "Post 3"])
     MyXQL.query(pid, "SELECT * FROM posts")

     MyXQL.query!(:myxql, "SELECT NOW()").rows


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

