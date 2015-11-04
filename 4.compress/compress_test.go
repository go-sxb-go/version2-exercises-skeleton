package compress

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

var (
	testCases = []struct {
		compressionType string
		expectedData    []byte
	}{
		{
			"none",
			[]byte("Hello World!"),
		}, {
			"gzip",
			[]byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xf2, 0x48, 0xcd, 0xc9, 0xc9, 0x57, 0x08, 0xcf, 0x2f, 0xca, 0x49, 0x51, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x1c, 0x29, 0x1c, 0x0c, 0x00, 0x00, 0x00},
		}, {
			"lzw",
			[]byte{0x24, 0x19, 0x4d, 0x86, 0xc3, 0x78, 0x80, 0xae, 0x6f, 0x39, 0x1b, 0x0c, 0x82, 0x18, 0x08},
		},
	}
)

func TestCompressToNone(t *testing.T) {
	for _, testCase := range testCases {
		buffer := bytes.NewBuffer([]byte("Hello World!"))
		file, err := CompressTo(buffer, testCase.compressionType)
		if err != nil {
			t.Errorf("[%v] expected no error, got %v", testCase.compressionType, err)
		}

		out, err := ioutil.ReadFile(file)
		if err != nil {
			t.Errorf("[%v] error when opening created file %v: %v", testCase.compressionType, file, err)
		}

		if !bytes.Equal(out, testCase.expectedData) {
			t.Errorf("[%v] expecting '%x', got '%x'", testCase.compressionType, string(testCase.expectedData), string(out))
		}
		os.Remove(file)
	}
}
