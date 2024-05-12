// import necessary modules
import gleam/int.{to_float}
import gleam/float.{to_string}

// pence_to_pounds turns pence to pounds and gives it back
pub fn pence_to_pounds(pence: Int) -> Float { to_float(pence) /. 100.0 }

// pounds_to_string concats the pounds to the pounds
pub fn pounds_to_string(pounds: Float) -> String {
  "£" <> to_string(pounds)
}
