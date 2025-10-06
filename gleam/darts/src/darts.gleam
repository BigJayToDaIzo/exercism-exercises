import gleam/float
import gleam/result

pub fn score(x: Float, y: Float) -> Int {
  case calculate_hypot(x, y) {
    h if h <=. 1.0 -> 10
    h if h <=. 5.0 -> 5
    h if h <=. 10.0 -> 1
    _ -> 0
  }
}

fn calculate_hypot(x: Float, y: Float) -> Float {
  result.unwrap(float.square_root(x *. x +. y *. y), 0.0)
}
