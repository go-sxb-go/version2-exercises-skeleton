package files

import "testing"

func TestCountCharInFile(t *testing.T) {
	count, err := CountCharInFiles("./files_test", 'e')
	if err != nil {
		t.Fatal("Got error", err, "expected no error.")
	}

	expected := 6
	if expected != count {
		t.Fatal("Expected", expected, "'e', got", count)
	}

	count, err = CountCharInFiles("./files_test", ' ')
	if err != nil {
		t.Fatal("Got error", err, "expected no error.")
	}

	expected = 9
	if expected != count {
		t.Fatal("Expected", expected, "spaces, got", count)
	}
}
