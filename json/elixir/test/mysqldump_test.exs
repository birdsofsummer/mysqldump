defmodule MysqldumpTest do
  use ExUnit.Case
  doctest Mysqldump

  test "greets the world" do
    assert Mysqldump.hello() == :world
  end
end
