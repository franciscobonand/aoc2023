defmodule Solution do
  @cards ["A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"]
  @cards_count 13
  @types [:high, :pair, :two_pair, :three, :full_house, :four, :five]
  @cards_in_hand 5

  def get_hand_type(hand) do
    mapped_hand =
      String.split(hand, "", trim: true)
      |> Enum.reduce(%{}, fn card, map ->
        if not is_nil(map[card]) do
          Map.put(map, card, map[card] + 1)
        else
          Map.put(map, card, 1)
        end
      end)

    jokers_count = mapped_hand["J"]

    mapped_hand
    |> Map.drop(["J"])
    |> Map.values()
    |> Enum.sort(:desc)
    |> List.update_at(0, fn value ->
      if is_nil(jokers_count) do
        value
      else
        value + jokers_count
      end
    end)
    |> then(fn sorted_values ->
      case sorted_values do
        [5] -> :five
        # when all 5 are jokers there will be no values
        [] -> :five
        [4, 1] -> :four
        [3, 2] -> :full_house
        [3, 1, 1] -> :three
        [2, 2, 1] -> :two_pair
        [2, 1, 1, 1] -> :pair
        _ -> :high
      end
    end)
  end

  def get_positional_score(hand) do
    String.split(hand, "", trim: true)
    |> Enum.with_index()
    |> Enum.reduce(0, fn {card, index}, acc ->
      card_index = Enum.find_index(@cards, fn c -> c == card end)

      card_positional_score =
        (@cards_count - card_index) * Integer.pow(@cards_count, @cards_in_hand - index)

      card_positional_score + acc
    end)
  end

  def score_hand(hand) do
    hand_type = Solution.get_hand_type(hand[:hand])
    type_score = Enum.find_index(@types, fn type -> type == hand_type end)
    positional_score = Solution.get_positional_score(hand[:hand])

    %{
      hand: hand[:hand],
      bid: hand[:bid],
      type_score: type_score,
      positional_score: positional_score
    }
  end

  def rank_hands(hands) do
    hands_count = Enum.count(hands)

    hands
    |> Enum.map(&Solution.score_hand/1)
    |> Enum.sort(fn a, b ->
      type_a = a[:type_score]
      type_b = b[:type_score]
      positional_a = a[:positional_score]
      positional_b = b[:positional_score]

      case {type_a > type_b, type_a == type_b, positional_a > positional_b,
            positional_a == positional_b} do
        {true, _, _, _} -> true
        {false, false, _, _} -> false
        {false, true, false, false} -> false
        {false, true, true, _} -> true
        {false, true, false, true} -> true
        {_, _, _, _} -> false
      end
    end)
    |> Enum.with_index()
    |> Enum.map(fn {scored_hand, index} ->
      Map.put(scored_hand, :rank, hands_count - index)
    end)
  end
end

hands =
  File.stream!("input.txt")
  |> Stream.map(&String.trim/1)
  |> Stream.map(&String.split/1)
  |> Enum.reduce([], fn values, acc ->
    acc ++
      [
        %{
          hand: Enum.at(values, 0),
          bid:
            Enum.at(values, 1)
            |> String.to_integer()
        }
      ]
  end)

# IO.inspect(hands, charlists: :as_lists)

ranked_hands = Solution.rank_hands(hands)
# IO.inspect(ranked_hands, charlists: :as_lists)

ranked_hands
|> Enum.map(fn hand ->
  hand[:rank] * hand[:bid]
end)
|> Enum.reduce(fn score, acc -> acc + score end)
|> IO.inspect(charlists: :as_lists)
