import gleam/dict.{type Dict}
import gleam/list
import gleam/string

pub fn transform(legacy: Dict(Int, List(String))) -> Dict(String, Int) {
  dict.fold(legacy, dict.new(), fn(dict_acc, k, vals) {
    list.fold(vals, dict_acc, fn(acc, v) {
      dict.insert(acc, string.lowercase(v), k)
    })
  })
}
