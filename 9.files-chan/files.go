package files

type CountCharWorker struct {
	dir  string
	char rune
	sum  int

	fileChan chan string
	sumChan  chan int
	errs     chan error
}

// This function will count the amount of a given character of a type in all files of a directory
func CountCharInFiles(dir string, c rune) (int, error) {
}
