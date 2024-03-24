package app

func App(readFile func(name string) ([]byte, error), fileName string) int {
	contentInBytes, _ := readFile(fileName)
	return len(contentInBytes)
}
