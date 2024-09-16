package iteration

const repeatCounter = 5

func Repeat(character string, reapeatCounter int) string {
	var repeated string

	for i := 0; i < repeatCounter; i++ {
		repeated += character
	}

	return repeated
}
