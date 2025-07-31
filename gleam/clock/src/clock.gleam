import gleam/int

pub type Clock {
  Clock(hour: Int, minute: Int)
}

pub fn create(hour hour: Int, minute minute: Int) -> Clock {
  case hour, minute {
    h, _ if h < 0 -> create(24 + h, minute) // lt 0s
    _, m if m < 0 -> create(hour - 1, m + 60)
    h, _ if h >= 24 -> create(h % 24, minute) // gt 24/60s
    _, m if m >= 60 -> create(hour + { minute / 60 }, minute % 60)
    _, _ -> Clock(hour, minute)
  }
}

pub fn add(clock: Clock, minutes minutes: Int) -> Clock {
  create(clock.hour, clock.minute + minutes)
}

pub fn subtract(clock: Clock, minutes minutes: Int) -> Clock {
  create(clock.hour, clock.minute - minutes)
}

pub fn display(clock: Clock) -> String {
  case clock.hour, clock.minute {
    h, m if h < 10 && m < 10 ->
      "0" <> int.to_string(h) <> ":" <> "0" <> int.to_string(m)
    h, _ if h < 10 ->
      "0" <> int.to_string(h) <> ":" <> int.to_string(clock.minute)
    _, m if m < 10 ->
      int.to_string(clock.hour) <> ":" <> "0" <> int.to_string(m)
    _, _ -> int.to_string(clock.hour) <> ":" <> int.to_string(clock.minute)
  }
}
