import gleam/list
import gleam/string

pub fn is_isogram(phrase phrase: String) -> Bool {
  string.trim(phrase)
  |> string.lowercase
  |> string.replace("-", "")
  |> string.replace(" ", "")
  |> string.to_graphemes()
  |> list.unique()
  |> list.length()
  == string.length(string.replace(phrase, "-", "") |> string.replace(" ", ""))
}
