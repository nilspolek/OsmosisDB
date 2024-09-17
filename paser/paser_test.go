package paser

import (
	"testing"
)

func TestParse(t *testing.T) {
	input := "SETNAME\r\nNILS\n"
	paser := NewPaser()
	paser.Parse()
}
