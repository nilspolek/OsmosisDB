package paser

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	input := "SETNAME\r\n@NILS\n"
	paser := NewPaser()
	output, err := paser.Parse([]byte(input))

	fmt.Println(output.Type)
	fmt.Println(output.Keyword)
	fmt.Println(string(output.DataType))
	fmt.Println(string(output.Data))

	if err != nil {
		t.Fatal(err)
	}
	if output.Type != SET {
		t.Fatal("Expected Command: SET")
	}

	input = "SkTNAME\r\n@NILS\n"
	output, err = paser.Parse([]byte(input))

	if err == nil {
		t.Fatal("Expected error")
	}

}
