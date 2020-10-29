package iteration

//Repeat concatenates a letter 5 times and return it
func Repeat(letter string) string {

	var res string

	for i := 0; i < 5; i++ {
		res += letter
	}

	return res
}
