import gleam/bool
import gleam/float
import gleam/int
import gleam/list
import gleam/order

pub type Classification {
  Perfect
  Abundant
  Deficient
}

pub type Error {
  NonPositiveInt
}

pub fn classify(number: Int) -> Result(Classification, Error) {
  use <- bool.guard(number < 1, Error(NonPositiveInt))
  Ok(is_aliquot_sum(number))
}

fn is_aliquot_sum(number: Int) -> Classification {
  let sum_factors = int.sum(factors(number))
  case int.compare(number, sum_factors) {
    order.Gt -> Deficient
    order.Lt -> Abundant
    order.Eq -> Perfect
  }
}

pub fn factors(n: Int) -> List(Int) {
  case n {
    n if n == 1 -> []
    _ -> {
      let assert Ok(sqrt_n) = int.square_root(n)
      let sqrt_n_int = float.truncate(sqrt_n)
      list.range(1, sqrt_n_int)
      |> list.fold([], fn(acc, i) {
        case n % i {
          // Perfect square case
          0 if i * i == n -> [i, ..acc]
          // Add both factors
          0 -> [i, n / i, ..acc]
          // Not a factor
          _ -> acc
        }
      })
      |> list.filter(fn(x) { x != n })
      |> list.sort(int.compare)
    }
  }
}
