package hamming

import "errors"

// Distance checks the hamming distance between two DNA strands
func Distance(a, b string) (d int, err error) {
	if len(a) != len(b) {
		return 0, errors.New("strands are not the same length")
	}

	d = 0
	for i := range a {
		if a[i] != b[i] {
			d++
		}
	}

	return d, nil
}
