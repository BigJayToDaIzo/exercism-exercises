import gleam/int
import gleam/result
import gleam/string

pub fn is_armstrong_number(number: Int) -> Bool {
  let power = int.to_string(number) |> string.length
  let digit_list = string.to_graphemes(int.to_string(number))
  let acc = 0
  is_armstrong_number_acc(digit_list, number, power, acc)
}

fn is_armstrong_number_acc(
  digits: List(String),
  number: Int,
  power: Int,
  acc: Int,
) -> Bool {
  case digits {
    [] -> number == acc
    [base, ..rest] ->
      is_armstrong_number_acc(
        rest,
        number,
        power,
        acc + pow(int.parse(base) |> result.unwrap(0), power),
      )
  }
}

// using the int.power method adds messy result unwrapping
// and rounding floats breaks on last 2 tests (many digits)
// rolled my own integer based power fn
fn pow(base: Int, exp: Int) -> Int {
  case exp {
    0 -> 1
    _ -> base * pow(base, exp - 1)
  }
}
