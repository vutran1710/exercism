use Bitwise, only_operators: true
defmodule SecretHandshake do



  @spec commands(code :: integer) :: list(String.t())
  def commands(code) do
    secret = [1, 2, 4, 8]
    hsk = %{
      1 => "wink",
      2 => "double blink",
      4 => "close your eyes",
      8 => "jump"
    }
    match = &(Enum.map(secret, fn x -> x &&& &1 end))
    translate = &(Enum.map(&1, fn x -> hsk[x] end))
    countbit = &(Integer.digits(&1, 2) |> Enum.count())
    only_truthy = &(Enum.filter(&1, fn x -> !!x end))
    parse = &(match.(&1) |> translate.() |> only_truthy.())

    bitlen = countbit.(code)
    cond do
      bitlen < 5 -> parse.(code)
      bitlen == 5 -> parse.(code) |> Enum.reverse()
      true -> []
    end

  end
end
