import gleam/list
import gleam/string

pub fn distance(strand1: String, strand2: String) -> Result(Int, Nil) {
  case string.length(strand1) != string.length(strand2) {
    True -> Error(Nil)
    _ -> {
      let str1_graphemes = string.to_graphemes(strand1)
      let str2_graphemes = string.to_graphemes(strand2)

      let hamming_distance =
        list.zip(str1_graphemes, str2_graphemes)
        |> list.filter(fn(t) { t.0 == t.1 })
        |> list.length

      case hamming_distance, list.length(str1_graphemes) {
        ham_dist, str_len if ham_dist < str_len -> Ok(str_len - ham_dist)
        ham_dist, str_len if ham_dist == str_len -> Ok(0)
        _, _ -> Error(Nil)
      }
    }
  }
}
