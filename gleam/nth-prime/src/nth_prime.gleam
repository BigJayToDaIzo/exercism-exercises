import gleam/bool
import gleam/float
import gleam/int
import gleam/list.{Continue, Stop}

pub fn prime(number: Int) -> Result(Int, Nil) {
  use <- bool.guard(number < 1, Error(Nil))
  list.range(2, 1_000_000)
  |> list.fold_until(#(0, 0), fn(acc, candidate) {
    use <- bool.guard(is_prime(candidate), {
      let count = acc.0 + 1
      use <- bool.guard(count == number, Stop(#(count, candidate)))
      Continue(#(count, candidate))
    })
    Continue(acc)
  })
  |> fn(result) { Ok(result.1) }
}

fn is_prime(n: Int) -> Bool {
  case n {
    n if n < 2 -> False
    n if n == 2 || n == 3 -> True
    n if n % 2 == 0 -> False
    n -> {
      let assert Ok(limit) = int.square_root(n)
      let int_limit = float.round(limit)
      list.range(3, int_limit)
      |> list.filter(fn(x) { x % 2 == 1 })
      |> list.all(fn(divisor) { n % divisor != 0 })
    }
  }
}
