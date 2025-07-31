import gleam/list
import gleam/string

pub type Student {
  Alice
  Bob
  Charlie
  David
  Eve
  Fred
  Ginny
  Harriet
  Ileana
  Joseph
  Kincaid
  Larry
}

pub type Plant {
  Radishes
  Clover
  Violets
  Grass
}

pub fn plants(diagram: String, student: Student) -> List(Plant) {
  parse_student_for_box_corner(student)
  |> clip_box(diagram)
}

fn clip_box(box_corner: Int, diagram: String) -> List(Plant) {
  let plant_rows = string.split(diagram, "\n")
  case plant_rows {
    [back_row, front_row] ->
      build_plant_list(
        string.to_graphemes(slice_rows([back_row, front_row], box_corner, "")),
        [],
      )
    // Error handle ph
    _ -> []
  }
}

fn build_plant_list(
  row_graphemes: List(String),
  plants: List(Plant),
) -> List(Plant) {
  case row_graphemes {
    [] -> list.reverse(plants)
    [p, ..rest] -> build_plant_list(rest, [grapheme_to_plant(p), ..plants])
  }
}

pub fn slice_rows(rows: List(String), box_corner: Int, acc: String) -> String {
  case rows {
    [] -> acc
    [r, ..rest] ->
      slice_rows(
        rest,
        box_corner,
        acc <> string.slice(r, at_index: box_corner, length: 2),
      )
  }
}

fn parse_student_for_box_corner(student: Student) -> Int {
  case student {
    Alice -> 0
    Bob -> 2
    Charlie -> 4
    David -> 6
    Eve -> 8
    Fred -> 10
    Ginny -> 12
    Harriet -> 14
    Ileana -> 16
    Joseph -> 18
    Kincaid -> 20
    Larry -> 22
  }
}

fn grapheme_to_plant(g: String) -> Plant {
  case g {
    "G" -> Grass
    "C" -> Clover
    "R" -> Radishes
    _ -> Violets
  }
}
