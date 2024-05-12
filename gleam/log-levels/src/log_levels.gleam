import gleam/string
pub fn message(log_line: String) -> String {
  case log_line {
    "[ERROR]: " <> e_message -> string.trim(e_message) 
    "[WARNING]: " <> w_message -> string.trim(w_message)
    "[INFO]: " <> i_message -> string.trim(i_message)
    _ -> "No message"
  }
}
pub fn log_level(log_line: String) -> String {
  case log_line {
    "[ERROR]: " <> _ -> "error"
    "[WARNING]: " <> _-> "warning"
    "[INFO]: " <> _ -> "info"
    _ -> "No log level"
  }
}
pub fn reformat(log_line: String) -> String {
  case log_line {
    "[ERROR]: " <> message -> string.trim(message) <> " (error)"
    "[WARNING]: " <> message -> string.trim(message) <> " (warning)"
    "[INFO]: " <> message -> string.trim(message) <> " (info)"
    _ -> "No message"
  }
}
