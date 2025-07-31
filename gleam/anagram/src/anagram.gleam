import gleam/list
import gleam/string

pub fn find_anagrams(word: String, candidates: List(String)) -> List(String) {
  let anagrams = []
  let word_graphemes =
    word
    |> string.lowercase
    |> string.to_graphemes
  find_anagrams_acc(word_graphemes, candidates, anagrams)
}

fn find_anagrams_acc(
  word: List(String),
  candidates: List(String),
  acc: List(String),
) -> List(String) {
  case candidates {
    [] -> list.reverse(acc)
    [a, ..rest] -> {
      let candidate_graphemes = string.to_graphemes(string.lowercase(a))
      case
        list.sort(candidate_graphemes, string.compare)
        == list.sort(word, string.compare)
        && word != candidate_graphemes
      {
        True -> find_anagrams_acc(word, rest, [a, ..acc])
        _ -> find_anagrams_acc(word, rest, acc)
      }
    }
  }
}
