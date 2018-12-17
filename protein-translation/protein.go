package protein

type stopError struct {
	s string
}

func (stop *stopError) Error() string {
	return stop.s
}

var STOP = &stopError{"STOP"}

type tuple struct {
	s string
	e error
}

var (
	methionine = tuple{"Methionine", nil}
	phenylalanine = tuple{"Phenylalanine", nil}
	leucine = tuple{"Leucine", nil}
	serine = tuple{"Serine", nil}
	tyrosine = tuple{"Tyrosine", nil}
	cysteine = tuple{"Cysteine", nil}
	tryptophan = tuple{"Tryptophan", nil}
	stop = tuple{"", STOP}
)


var codons = map[string]tuple{
	"AUG": methionine,
	"UUU": phenylalanine,
	"UUC": phenylalanine,
	"UUA": leucine,
	"UUG": leucine,
	"UCU": serine,
	"UCC": serine,
	"UCA": serine,
	"UCG": serine,
	"UAU": tyrosine,
	"UAC": tyrosine,
	"UGU": cysteine,
	"UGC": cysteine,
	"UGG": tryptophan,
	"UAA": stop,
	"UAG": stop,
	"UGA": stop,
}

func FromCodon(input string) (string, error) {
	t := codons[input]
	return t.s, t.e
}

func FromRNA(input string) ([]string, error) {
	slice := []string{}
	for index := 0; index < len(input); index += 3 {
		codon, err := FromCodon(input[index : index+3])
		if err == STOP {
			break
		}
		slice = append(slice, codon)
	}
	return slice, nil
}
