package dna

var opposite = map[rune]string{
	'A': "T",
	'T': "A",
	'C': "G",
	'G': "C",
}

func DNAStrand(dna string) (reverted string) {
	for _, r := range dna {
		reverted += opposite[r]
	}
	return
}
