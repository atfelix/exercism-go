package protein

var codons = map[string]string {
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

func FromCodon(input string) string {
	return codons[input]
}

func FromRNA(input string) []string {
	slice := []string{}
	for index := 0; index < len(input); index += 3 {
		codon := FromCodon(input[index:index + 3])
		if codon == "STOP" {
			break
		}
		slice = append(slice, codon)
	}
	return slice
}
