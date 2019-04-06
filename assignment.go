package golangexamples


func ConcatSLice(sliceToConcat []byte) string {
	var s string
	for i := 0; i < len(sliceToConcat); i++ {
		//fmt.Printf("%d\n", i)
		if i == len(sliceToConcat)-1 {
			s += string(sliceToConcat[i])
		} else {

			s += string(sliceToConcat[i]) + "-"

		}
	}

	return s
}

func Encrypt(slicetoConcat []byte, ceaserCount int) []byte {
	for i := 0; i < len(slicetoConcat); i++ {
		slicetoConcat[i] = slicetoConcat[i] + 3
	}
	return slicetoConcat
}

func PrintGreetings(a string) string {
	return "Hello world - " + a
}
