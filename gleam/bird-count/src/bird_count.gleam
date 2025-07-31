import gleam/list

pub fn today(days: List(Int)) -> Int {
  case days {
    [] -> 0
    [today, ..] -> today
  }

}

pub fn increment_day_count(days: List(Int)) -> List(Int) {
  case days {
    [] -> [1]
    [a, ..rest] -> [a+1, ..rest]
  }
}

pub fn has_day_without_birds(days: List(Int)) -> Bool {
  list.any(days, fn(d) {d == 0})
}

pub fn total(days: List(Int)) -> Int {
  let acc = 0
  list.fold(days, acc, fn(acc, x) { acc + x })
}

pub fn busy_days(days: List(Int)) -> Int {
  list.count(days, fn(d) {d > 4})
}
