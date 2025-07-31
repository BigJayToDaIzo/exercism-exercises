pub type Error {
  NonPositiveNumber
}

pub fn steps(number: Int) -> Result(Int, Error) {
  let steps = 0
  do_steps(number, steps)
}

fn do_steps(number: Int, steps: Int) -> Result(Int, Error) {
  case number {
    _ if number <= 0 -> Error(NonPositiveNumber)
    1 -> Ok(steps)
    _ if number % 2 == 0 -> do_steps(number / 2, steps + 1)
    _ -> do_steps(number * 3 + 1, steps + 1)
  }
}
