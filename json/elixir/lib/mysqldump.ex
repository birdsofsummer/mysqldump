defmodule Mysqldump do
  @moduledoc """
  Documentation for `Mysqldump`.
  """

  @doc """
  dump mysql to json file.

  ## Examples

      iex> Mysqldump.hello()
      :world

  """
    def show_dbs(config) do
         {:ok,pid} = MyXQL.start_link(config)
         {:ok, %{rows: r}}=MyXQL.query(pid, "show databases;")
         dbs=r|>Enum.map(fn x-> hd(x) end)
    end

    def show_table(pid,table) do
         ##table="user"
         {:ok,%{rows: rows ,columns: columns}}=MyXQL.query(pid, "SELECT * FROM #{table}")
         d=rows
         |>Enum.map(fn x->Enum.zip(columns,x) end)
         |>Enum.map(fn x->Map.new(x) end) 
         |>Jason.encode!
         File.write("/tmp/a.json",d)
    end

    def show_tables(pid) do
         {:ok,%{rows: r}}=MyXQL.query(pid, "show tables;")
         r
         |>Enum.map(&(hd(&1)))
         |>Enum.map(&(show_table(pid,&1)))
    end

    def export_json(config) do



    end
    def export_jsons(config) do
        dbs=show_dbs(config)
        configs=dbs|>Enum.map(&(Dict.put(config,:database,&1)))
        configs|>Enum.map(&(export_json(&1)))
    end
end

