import gleam/int

pub fn egg_count(number: Int) -> Int {
  count_bits(number, 0)
}

fn count_bits(n: Int, acc: Int) -> Int {
  case n {
    0 -> acc
    _ -> count_bits(int.bitwise_shift_right(n, 1), acc + int.bitwise_and(n, 1))
  }
}
