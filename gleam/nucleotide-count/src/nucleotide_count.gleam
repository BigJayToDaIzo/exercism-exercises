import gleam/bool
import gleam/dict.{type Dict}
import gleam/list
import gleam/string

pub fn nucleotide_count(dna: String) -> Result(Dict(String, Int), Nil) {
  let graphemes = string.to_graphemes(dna)
  use <- bool.guard(!list.all(graphemes, is_valid_nucleotide), Error(Nil))
  let base_dict = dict.from_list([#("A", 0), #("C", 0), #("G", 0), #("T", 0)])
  graphemes
  |> list.fold(base_dict, fn(acc, grapheme) {
    let assert Ok(val) = dict.get(acc, grapheme)
    dict.insert(acc, grapheme, val + 1)
  })
  |> Ok
}

fn is_valid_nucleotide(nuc: String) -> Bool {
  nuc == "A" || nuc == "C" || nuc == "G" || nuc == "T"
}
