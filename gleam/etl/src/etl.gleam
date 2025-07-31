import gleam/dict.{type Dict}
import gleam/list
import gleam/string

pub fn transform(legacy: Dict(Int, List(String))) -> Dict(String, Int) {
  dict.fold(legacy, dict.new(), fn(acc, score, letter_list) {
    list.fold(letter_list, acc, fn(acc, letter) {
      dict.insert(acc, string.lowercase(letter), score)
      })
    })
}
