import gleam/list
import gleam/set.{type Set}
import gleam/string

pub fn is_pangram(sentence: String) -> Bool {
  let set_len =
    sentence
    |> string.lowercase
    |> string.to_graphemes
    |> set_builder(set.new())
  set_len == 26
}

fn set_builder(graphemes: List(String), acc: Set(String)) -> Int {
  let not_allowed_chars = [
    "!", "@", "#", "$", "%", "^", "&", "(", ")", "-", "=", "_", "+", " ", "\n",
    "\t", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", ".", ",",
  ]
  case graphemes {
    [] -> set.size(acc)
    [a, ..rest] -> {
      case list.contains(not_allowed_chars, a) {
        False -> set_builder(rest, set.insert(acc, a))
        _ -> set_builder(rest, acc)
      }
    }
  }
}
