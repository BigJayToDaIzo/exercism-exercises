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
  case languages {
    [a, ..] | [_, a] | [_, a, _] if a == "Gleam" -> True
    _ -> False
  }
}
