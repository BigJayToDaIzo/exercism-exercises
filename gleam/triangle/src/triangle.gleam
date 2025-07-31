import gleam/float

pub type Triangle {
  Triangle(a: Float, b: Float, c: Float)
}

fn is_valid(triangle: Triangle) -> Bool {
  case triangle.a, triangle.b, triangle.c {
    a, b, c
      if a +. b +. c != 0.0 && a +. b >=. c && b +. c >=. a && a +. c >=. b
    -> True
    _, _, _ -> False
  }
}

pub fn equilateral(a: Float, b: Float, c: Float) -> Bool {
  is_valid(Triangle(a, b, c))
  && float.loosely_equals(a, b, 0.01)
  && float.loosely_equals(b, c, 0.01)
  && float.loosely_equals(a, c, 0.01)
}

pub fn isosceles(a: Float, b: Float, c: Float) -> Bool {
  is_valid(Triangle(a, b, c))
  && {
    float.loosely_equals(a, b, 0.01)
    || float.loosely_equals(b, c, 0.01)
    || float.loosely_equals(a, c, 0.01)
  }
}

pub fn scalene(a: Float, b: Float, c: Float) -> Bool {
  is_valid(Triangle(a, b, c))
  && {
    !float.loosely_equals(a, b, 0.01)
    && !float.loosely_equals(b, c, 0.01)
    && !float.loosely_equals(a, c, 0.01)
  }
}
