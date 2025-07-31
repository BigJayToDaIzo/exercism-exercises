import gleam/list
import gleam/string

pub fn recite(start_verse start_verse: Int, end_verse end_verse: Int) -> String {
  list.range(start_verse, end_verse)
  |> list.map(fn(v) {"This is " <> build_verse(v)})
  |> string.join("\n")
}

fn build_verse(number: Int) -> String {
  case number {
    1 -> "the house that Jack built."
    2 -> "the malt that lay in " <> build_verse(1)
    3 -> "the rat that ate " <> build_verse(2)
    4 -> "the cat that killed " <> build_verse(3)
    5 -> "the dog that worried " <> build_verse(4)
    6 -> "the cow with the crumpled horn that tossed " <> build_verse(5)
    7 -> "the maiden all forlorn that milked " <> build_verse(6)
    8 -> "the man all tattered and torn that kissed "<> build_verse(7)
    9 -> "the priest all shaven and shorn that married " <> build_verse(8)
    10 -> "the rooster that crowed in the morn that woke " <> build_verse(9)
    11 -> "the farmer sowing his corn that kept " <> build_verse(10)
    12 -> "the horse and the hound and the horn that belonged to " <> build_verse(11)
    _ -> ""
  }
}
