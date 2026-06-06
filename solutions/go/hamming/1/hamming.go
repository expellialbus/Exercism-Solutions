package hamming

import "fmt"

func Distance(a, b string) (int, error) {
    var distance int
	if len(a) != len(b) {
        return 0, fmt.Errorf("Incompatible lengths.")
    } else {
        for i := 0; i < len(a); i++ {
            if a[i] != b[i] {
                distance++
            }
        }
    }

    return distance, nil
}
