package ini

import (
	"os"
	"testing"
)

func TestIni(t *testing.T) {
	testFil, err := os.Open("test.ini")
	if err != nil {
		t.Fatal(err)
	}
	f, err := Parse(testFil)
	if err != nil {
		t.Fatal(err)
	}
	if f.Section("Hello").Value("quoteCommentTest").String() != "Hello\\\"#YELLOW" {
		t.Fatal("quoteCommentTest not correct:", f.Section("Hello").Value("quoteCommentTest").String())
	}
	if f.PreSection().Value("notATest").String() != "Hello my name is george" {
		t.Fatal("notATest not correct:", f.PreSection().Value("notATest").String())
	}
	if f.Sections()[0] != "Hello" || f.Sections()[1] != "part 2" {
		t.Fatal("sections not correct:", f.Sections())
	}
}
