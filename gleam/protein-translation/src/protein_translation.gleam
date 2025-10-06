import gleam/bool
import gleam/list
import gleam/string

pub fn proteins(rna: String) -> Result(List(String), Nil) {
  do_parse_codons(rna, [])
}

fn do_parse_codons(rna: String, acc: List(String)) -> Result(List(String), Nil) {
  case rna {
    "" -> Ok(list.reverse(acc))
    _ -> {
      let codon = string.slice(rna, 0, 3)
      use <- bool.guard(string.length(codon) != 3, Error(Nil))
      let rest = string.drop_start(rna, 3)
      case codon {
        // early return for stop codon
        "UAA" | "UAG" | "UGA" -> Ok(list.reverse(acc))
        "AUG" -> do_parse_codons(rest, ["Methionine", ..acc])
        "UUU" | "UUC" -> do_parse_codons(rest, ["Phenylalanine", ..acc])
        "UUA" | "UUG" -> do_parse_codons(rest, ["Leucine", ..acc])
        "UCU" | "UCC" | "UCA" | "UCG" ->
          do_parse_codons(rest, ["Serine", ..acc])
        "UAU" | "UAC" -> do_parse_codons(rest, ["Tyrosine", ..acc])
        "UGU" | "UGC" -> do_parse_codons(rest, ["Cysteine", ..acc])
        "UGG" -> do_parse_codons(rest, ["Tryptophan", ..acc])
        _ -> Error(Nil)
      }
    }
  }
}
