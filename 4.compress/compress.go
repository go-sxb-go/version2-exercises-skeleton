package compress

import "io"

// CompressTo takes a file and a compression mode as input
// And return the path to a temp file containing the compressed
// version of the input
// mode can be 'flate', 'gzip', or 'none' which will just copy the input in the file
func CompressTo(input io.Reader, mode string) (string, error) {
}

func lzwCompress(input io.Reader, out io.Writer) error {
}

func gzipCompress(input io.Reader, out io.Writer) error {
}

func noneCompress(input io.Reader, out io.Writer) error {
}
