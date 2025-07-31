import gleam/result
import gleam/string
import gleam/list

import simplifile

pub fn read_emails(path: String) -> Result(List(String), Nil) {
  simplifile.read(path)
  |> result.unwrap("")
  |> string.trim
  |> string.split("\n")
  |> Ok()
}

pub fn create_log_file(path: String) -> Result(Nil, Nil) {
  simplifile.create_file(at: path)
  |> result.replace_error(Nil)
}

pub fn log_sent_email(path: String, email: String) -> Result(Nil, Nil) {
  simplifile.append(to: path, contents: email <> "\n")
  |> result.replace_error(Nil)
}

pub fn send_newsletter(
  emails_path: String,
  log_path: String,
  send_email: fn(String) -> Result(Nil, Nil),
) -> Result(Nil, Nil) {
  let assert Ok(_) = create_log_file(log_path)
  let assert Ok(emails) = read_emails(emails_path)
  {
    use email <- list.each(emails)
    use _ <- result.try(send_email(email))
    let assert Ok(_) = log_sent_email(log_path, email)
  }
  Ok(Nil)
}
