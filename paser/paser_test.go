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

func TestGET(t *testing.T) {
	input := "GETNAME\n"
	paser := NewPaser()
	output, err := paser.Parse([]byte(input))

	if err != nil {
		t.Fatal(err)
	}
	if output.Type != GET {
		t.Fatal("Expected Command: GET")
	}
	if output.Keyword != "NAME" {
		t.Fatalf("Expected Keyword: NAME, got %s", output.Keyword)
	}
	if output.DataType != 0 {
		t.Fatalf("Expected DataType: 0, got %v", output.DataType)
	}
	if output.Data != nil {
		t.Fatal("Expected no Data")
	}
}

func TestDEL(t *testing.T) {
	input := "DELNAME\n"
	paser := NewPaser()
	output, err := paser.Parse([]byte(input))

	if err != nil {
		t.Fatal(err)
	}
	if output.Type != DEL {
		t.Fatal("Expected Command: DEL")
	}
	if output.Keyword != "NAME" {
		t.Fatalf("Expected Keyword: NAME, got %s", output.Keyword)
	}
	if output.DataType != 0 {
		t.Fatalf("Expected DataType: 0, got %v", output.DataType)
	}
	if output.Data != nil {
		t.Fatal("Expected no Data")
	}
}

func TestUPT(t *testing.T) {
	input := "UPTNAME\r\n@NILS\n"
	paser := NewPaser()
	output, err := paser.Parse([]byte(input))

	if err != nil {
		t.Fatal(err)
	}
	if output.Type != UPT {
		t.Fatal("Expected Command: UPT")
	}
	if output.Keyword != "NAME" {
		t.Fatal("Expected Keyword: NAME")
	}
	if output.DataType != TYPESTRING {
		t.Fatalf("Expected DataType: %v", TYPESTRING)
	}
	if output.Data == nil {
		t.Fatal("Expected Data")
	}
	if string(output.Data) != "NILS" {
		t.Fatalf("Expected Data: NILS, got %s", string(output.Data))
	}
}

func TestERR(t *testing.T) {
	input := "ERR\n"
	paser := NewPaser()
	output, err := paser.Parse([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if output.Type != ERR {
		t.Fatal("Expected Command: ERR")
	}
}

func TestOK(t *testing.T) {
	input := "OK \n"
	paser := NewPaser()
	output, err := paser.Parse([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if output.Type != OK {
		t.Fatal("Expected Command: OK ")
	}
}

func TestSET(t *testing.T) {
	input := "SETNAME\r\n@NILS\n"
	paser := NewPaser()
	output, err := paser.Parse([]byte(input))

	if err != nil {
		t.Fatal(err)
	}
	if output.Type != SET {
		t.Fatal("Expected Command: SET")
	}
	if output.Keyword != "NAME" {
		t.Fatal("Expected Keyword: NAME")
	}
	if output.DataType != TYPESTRING {
		t.Fatalf("Expected DataType: %v", TYPESTRING)
	}
	if output.Data == nil {
		t.Fatal("Expected Data")
	}
	if string(output.Data) != "NILS" {
		t.Fatalf("Expected Data: NILS, got %s", string(output.Data))
	}
}
