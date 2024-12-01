
defmodule DistanceCalculator do
  def total_distance(left, right) do
    Enum.zip(Enum.sort(left), Enum.sort(right))
    |> Enum.map(fn {a, b} -> abs(a - b) end)
    |> Enum.sum()
  end


  def read_lists_from_file(filename) do
    File.read!(filename)
    |> String.split("\n")
    |> Enum.map(fn line ->
      case String.split(line, "   ") do
        [left, right] -> {String.to_integer(left), String.to_integer(right)}
        _ -> raise "Invalid line format: #{line}"
      end
    end)
    |> Enum.unzip()
  end
end

defmodule SimilarityCalculator do
  def similarity_score(left, right) do
    Enum.map(left, fn number ->
      number * Enum.count(right, &(&1 == number))
    end)
    |> Enum.sum()
  end
end


filename = "input.txt"

# Part 1
{left_list, right_list} = DistanceCalculator.read_lists_from_file(filename)
total_distance = DistanceCalculator.total_distance(left_list, right_list)
IO.inspect total_distance

# Part 2
similarity_score = SimilarityCalculator.similarity_score(left_list, right_list)
IO.inspect similarity_score
