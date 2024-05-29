package protein

import "errors"

var ErrInvalidBase error = errors.New("Invalid Base")
var ErrStop error = errors.New("STOP")

func FromRNA(rna string) ([]string, error) {
	proteins := []string{}
	for range rna {
		if len(rna) < 3 {
			break
		}
		codon := rna[:3]
		protein, err := FromCodon(codon)
		if err == ErrStop {
			return proteins, nil
		} else if err == ErrInvalidBase {
			return proteins, ErrInvalidBase
		}
		proteins = append(proteins, protein)
		rna = rna[3:]
	}
	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	codonToProtein := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	if v, ok := codonToProtein[codon]; ok {
		if v == "STOP" {
			return "", ErrStop
		}
		return v, nil
	} else {
		return "", ErrInvalidBase
	}
}
