import gleam/list

pub fn recite(inputs: List(String)) -> String {
  case list.length(inputs) {
    0 -> ""
    1 -> {
      let assert Ok(first) = list.first(inputs)
      "And all for the want of a " <> first <> "."
    }
    _ -> {
      let assert Ok(first) = list.first(inputs)
      do_recite(
        list.window_by_2(list.reverse(inputs)),
        "And all for the want of a " <> first <> ".",
      )
    }
  }
}

fn do_recite(inputs: List(#(String, String)), acc: String) -> String {
  case inputs {
    [] -> acc
    [#(first, second), ..rest] ->
      do_recite(
        rest,
        "For want of a " <> second <> " the " <> first <> " was lost.\n" <> acc,
      )
  }
}
