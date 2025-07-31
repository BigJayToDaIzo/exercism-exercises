import gleam/int
import gleam/string

pub fn egg_count(number: Int) -> Int {
  int.to_base2(number) 
  |> string.replace("0", "") 
  |> string.length
}
