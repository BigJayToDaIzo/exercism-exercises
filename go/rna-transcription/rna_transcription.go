package strand

func ToRNA(dna string) string {
	rna := ""
	for _, char := range dna {
		switch char {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		default:
		}
	}
	return rna
}
