import gleam/float
import gleam/int

const pence_per_pound = 100.0

pub fn pence_to_pounds(pence: Int) -> Float {
  int.to_float(pence) /. pence_per_pound
}

pub fn pounds_to_string(pounds: Float) -> String {
  "£" <> float.to_string(pounds)
}
