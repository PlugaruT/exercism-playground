package dna

import (
	"errors"
	"fmt"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

const validNucleotides = "ACGT"

func (d DNA) count(nucleotide byte) (int, error) {
	if !strings.Contains(validNucleotides, string(nucleotide)) {
		return 0, errors.New("dna: invalid nucleotide " + string(nucleotide))
	}
	return strings.Count(string(d), string(nucleotide)), nil
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	t := 0
	histogram := Histogram{}
	for _, nucleotide := range validNucleotides {
		histogram[nucleotide], _ = d.count(byte(nucleotide))
		t += histogram[nucleotide]
	}
	if t != len(d) {
		return nil, fmt.Errorf("invalid nucleotide present")
	}
	return histogram, nil
}
