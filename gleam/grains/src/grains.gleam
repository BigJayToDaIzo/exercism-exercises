import gleam/result
import gleam/list

pub type Error {
  InvalidSquare
}

pub fn square(sq: Int) -> Result(Int, Error) {
  case sq {
    n if n < 1 || n > 64 -> Error(InvalidSquare)
    1 -> Ok(1)
    n -> result.map(square(n - 1), fn(x) { 2 * x })
  }
}


pub fn total() -> Int {
  list.fold(list.range(1, 64), 0, fn(acc, sq) {
    acc + result.unwrap(square(sq), 0)
    })
}
