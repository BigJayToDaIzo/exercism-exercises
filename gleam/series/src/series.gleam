import gleam/list
import gleam/string

pub opaque type GTZeroInt {
  GTZeroInt(inner: Int)
}

pub fn from_int(i: Int) -> Result(GTZeroInt, Error) {
  case i {
    _ if i == 0 -> Error(SliceLengthZero)
    _ if i < 0 -> Error(SliceLengthNegative)
    _ -> Ok(GTZeroInt(i))
  }
}

pub fn to_int(i: GTZeroInt) -> Int {
  i.inner
}

pub fn slices(input: String, size: Int) -> Result(List(String), Error) {
  let acc = []
  let series_len = string.length(input)
  case series_len == 0, size > series_len, from_int(size) {
    True, _, _ -> Error(EmptySeries)
    _, True, _ -> Error(SliceLengthTooLarge)
    _, _, Error(e) -> Error(e)
    _, _, Ok(size) -> slices_acc(input, to_int(size), acc)
  }
}

fn slices_acc(
  input: String,
  size: Int,
  acc: List(String),
) -> Result(List(String), Error) {
  case string.length(input) >= size {
    True ->
      slices_acc(string.drop_start(input, 1), size, [
        string.slice(from: input, at_index: 0, length: size),
        ..acc
      ])
    _ -> Ok(list.reverse(acc))
  }
}

pub type Error {
  SliceLengthTooLarge
  SliceLengthZero
  SliceLengthNegative
  EmptySeries
}
