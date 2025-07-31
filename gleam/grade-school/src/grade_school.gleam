import gleam/dict.{type Dict}
import gleam/int
import gleam/list

pub type School {
  School(roster: Dict(String, Int))
}

pub fn create() -> School {
  School(roster: dict.new())
}

pub fn roster(school: School) -> List(String) {
  dict.to_list(school.roster)
  |> list.sort(fn(a, b) {
    case a, b {
      #(_, g1), #(_, g2) -> int.compare(g1, g2)
    }
  })
  |> list.map(fn(grade) { grade.0 })
}

pub fn add(
  to school: School,
  student student: String,
  grade grade: Int,
) -> Result(School, Nil) {
  // assert student isn't already in the school
  case dict.has_key(school.roster, student) {
    True -> Error(Nil)
    _ -> Ok(School(roster: dict.insert(school.roster, student, grade)))
  }
}

pub fn grade(school: School, desired_grade: Int) -> List(String) {
  dict.filter(school.roster, keeping: fn(_, grade) { grade == desired_grade })
  |> dict.to_list()
  |> list.map(fn(grade) { grade.0 })
}
