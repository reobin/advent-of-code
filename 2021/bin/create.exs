# Creates all necessary files for a specific day challenge

alias AoC.Helpers.FileHelper
alias AoC.Helpers.SystemHelper

day = SystemHelper.get_day()

if is_nil(day) do
  raise "Please provide a day: mix aoc.create 12"
  exit(1)
end

IO.puts("Creating files for day #{day}")

module_file_path = "lib/day_#{day}.ex"
test_file_path = "test/day_#{day}_test.exs"
input_file_path = "input/day_#{day}.txt"

[module_file_path, test_file_path, input_file_path]
|> Enum.each(fn path ->
  if File.exists?(path), do: exit("#{path} exists")
end)

# Create module file
module_file_content = """
defmodule AoC.Day#{day} do
  def part_1(input) do
    IO.inspect(input)
    "part 1"
  end

  def part_2(input) do
    IO.inspect(input)
    "part 2"
  end
end
"""

{:ok} = FileHelper.write(module_file_path, module_file_content)

# Create test module file

test_file_content = """
defmodule AoC.Day#{day}Test do
  use ExUnit.Case
  doctest AoC.Day#{day}

  alias AoC.Day#{day}

  describe "part 1" do
    test "part 1" do
      input = "test input"
      assert Day#{day}.part_1(input) == "part 1"
    end
  end

  describe "part 2" do
    test "part 2" do
      input = "test input"
      assert Day#{day}.part_2(input) == "part 2"
    end
  end
end
"""

{:ok} = FileHelper.write(test_file_path, test_file_content)

{:ok} = FileHelper.write(input_file_path, "")
