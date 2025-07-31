import gleam/float
import gleam/result

pub fn score(x: Float, y: Float) -> Int {
  let dist_from_center =
    float.square_root({x *. x} +. {y *. y})
    |> result.unwrap(0.0)
  case dist_from_center {
    d if d >. 10.0 -> 0
    d if d >. 5.0 -> 1
    d if d >. 1.0 -> 5
    _ -> 10
  }
}
