import gleam/bool
import gleam/int
import gleam/list
import gleam/result
import gleam/string

pub fn row(index: Int, s: String) -> Result(List(Int), Nil) {
  use matrix <- result.try(string_to_int_matrix(s))
  use <- bool.guard(list.length(matrix) < index, Error(Nil))
  list.drop(matrix, index - 1) |> list.first
}

pub fn column(index: Int, s: String) -> Result(List(Int), Nil) {
  use matrix <- result.try(string_to_int_matrix(s))
  let transposed = list.transpose(matrix)
  use <- bool.guard(list.length(transposed) < index, Error(Nil))
  list.drop(transposed, index - 1) |> list.first
}

fn string_to_int_matrix(s: String) -> Result(List(List(Int)), Nil) {
  string.split(s, "\n")
  |> list.map(fn(str) { string.split(str, " ") })
  |> list.map(fn(row_str) {
    list.map(row_str, fn(digit_str) { int.parse(digit_str) })
    |> result.all
  })
  |> result.all
}

