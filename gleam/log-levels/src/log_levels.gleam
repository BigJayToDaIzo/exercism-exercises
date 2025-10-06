import gleam/string

type LogEntry {
  LogEntry(level: String, message: String)
}

fn parse_log_line(log_line: String) -> LogEntry {
  case log_line {
    "[ERROR]: " <> rest -> LogEntry("error", string.trim(rest))
    "[INFO]: " <> rest -> LogEntry("info", string.trim(rest))
    "[WARNING]: " <> rest -> LogEntry("warning", string.trim(rest))
    _ -> LogEntry("", "")
  }
}

pub fn message(log_line: String) -> String {
  parse_log_line(log_line).message
}

pub fn log_level(log_line: String) -> String {
  parse_log_line(log_line).level
}

pub fn reformat(log_line: String) -> String {
  let entry = parse_log_line(log_line)
  entry.message <> " (" <> entry.level <> ")"
}
