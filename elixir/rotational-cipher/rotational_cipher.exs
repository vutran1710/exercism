defmodule RotationalCipher do
  @doc """
  Given a plaintext and amount to shift by, return a rotated string.

  Example:
  iex> RotationalCipher.rotate("Attack at dawn", 13)
  "Nggnpx ng qnja"
  """

  @graph "abcdefghijklmnopqrstuvwxyz"
  @spec rotate(text :: String.t(), shift :: integer) :: String.t()
  def rotate(text, shift) do
    text
    |> String.split("")
    |> Enum.map_join("", &(char_rotate &1, shift))
  end

  def char_rotate(c, shift) do
    normalized = String.downcase(c)
    cond do
      c == "" || current_idx(normalized) == :nomatch -> c
      true -> current_idx(normalized)
      |> elem(0)
      |> calc(shift)
      |> find_char
      |> restore_case(String.upcase(c) == c)
    end
  end

  defp current_idx(c), do: :binary.match(@graph, c)
  defp calc(n, shift), do: Kernel.trunc(n + shift < 26 && n + shift || n + shift - 26)
  defp find_char(idx), do: String.at @graph, idx
  defp restore_case(txt, should?), do: should? && String.upcase(txt) || txt
end
