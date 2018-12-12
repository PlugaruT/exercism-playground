package strand

// ToRNA converts a RND sequence to RNA complement
func ToRNA(dna string) string {
	maps := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	var response = ""
	for _, char := range dna {
		response += maps[string(char)]
	}
	return response
}
