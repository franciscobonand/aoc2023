defmodule Solution do
  # race time = rt
  # time held = th
  # speed = th
  # distance = (rt - th) * speed -> rt*th - th^2
  # a = -1, b = rt
  # vertex = -b/2a -> rt/2

  def parse_values(line) do
    String.split(line, ":")
    |> Enum.at(1)
    |> String.split()
    |> Enum.reduce("", fn v, acc -> acc <> v end)
    |> String.to_integer()
  end

  def get_distance_by_time_held(race_time, time_held) do
    race_time * time_held - time_held * time_held
  end

  def get_max_time_helds(race_time) do
    max = race_time / 2

    if Float.ceil(max) == Float.floor(max) do
      [floor(max)]
    else
      [floor(max), ceil(max)]
    end
  end

  def find_edge_time_held(race_time, time_held, record_distance, op, step) do
    distance = Solution.get_distance_by_time_held(race_time, time_held)

    adjacent_time_held =
      if op == :inc do
        time_held + 1
      else
        time_held - 1
      end

    adjacent_distance = Solution.get_distance_by_time_held(race_time, adjacent_time_held)

    is_over_edge = distance < record_distance
    is_adjacent_over_edge = adjacent_distance < record_distance

    case {op, is_over_edge, is_adjacent_over_edge} do
      {:inc, false, true} ->
        time_held

      {:inc, true, false} ->
        adjacent_time_held

      {:dec, true, false} ->
        adjacent_time_held

      {:dec, false, true} ->
        time_held

      {_, _, _} ->
        next_step = ceil(step / 2)

        next_time_held =
          case {op, is_over_edge} do
            {:inc, false} -> time_held + next_step
            {:inc, true} -> time_held - next_step
            {:dec, false} -> time_held - next_step
            {:dec, true} -> time_held + next_step
          end

        Solution.find_edge_time_held(
          race_time,
          next_time_held,
          record_distance,
          op,
          next_step
        )
    end
  end

  def get_record_breaking_time_held_range(race_time, record_distance, max_time_helds) do
    max_time_held_floor = Enum.at(max_time_helds, 0)
    max_time_held_ceil = Enum.at(max_time_helds, -1)
    initial_step = ceil(race_time / 2)

    record_breaking_time_held_floor =
      Solution.find_edge_time_held(
        race_time,
        max_time_held_floor,
        record_distance,
        :dec,
        initial_step
      )

    record_breaking_time_held_ceil =
      Solution.find_edge_time_held(
        race_time,
        max_time_held_ceil,
        record_distance,
        :inc,
        initial_step
      )

    {record_breaking_time_held_floor, record_breaking_time_held_ceil}
  end
end

inputs =
  File.stream!("input.txt")
  |> Stream.map(&String.trim/1)

race_time =
  Enum.at(inputs, 0)
  |> Solution.parse_values()

record_distance =
  Enum.at(inputs, 1)
  |> Solution.parse_values()

max_time_helds =
  Solution.get_max_time_helds(race_time)
  |> IO.inspect(charlists: :as_lists)

record_breaking_time_held_range =
  Solution.get_record_breaking_time_held_range(race_time, record_distance, max_time_helds)

IO.inspect(record_breaking_time_held_range, charlists: :as_lists)

{r_start, r_end} = record_breaking_time_held_range
IO.inspect(r_end - r_start + 1)
