import gleam/list

pub fn new_list() -> List(String) {
  []
}

pub fn existing_list() -> List(String) {
  ["Gleam", "Go", "TypeScript"]
}

pub fn add_language(languages: List(String), language: String) -> List(String) {
  [language, ..languages]
}

pub fn count_languages(languages: List(String)) -> Int {
  list.length(languages)
}

pub fn reverse_list(languages: List(String)) -> List(String) {
  list.reverse(languages)
}

pub fn exciting_list(languages: List(String)) -> Bool {
  let exciting_length_list = list.length(languages) < 4
  case languages {
    [first, ..] if first == "Gleam" -> True
    [_, second, ..] if second == "Gleam" -> case exciting_length_list {
      True -> True
      _ -> False
    }
    _ -> False
  }
}
