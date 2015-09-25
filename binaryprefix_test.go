package binaryprefix

import (
	"fmt"
	"testing"
)

func TestParseBinaryPrefixMB(t *testing.T) {
	b, err := ParseBinaryPrefix("1MB")
	if err != nil {
		t.Error("Error parsing 1MB")
	}
	if int(b) != 1048576 {
		t.Error("1MB are 1048576 Bytes but got ", int(b))
	}
	if b.Bytes() != 1048576 {
		t.Error("1MB are 1048576 Bytes but got ", b.Bytes())
	}
	if b.KBytes() != 1024 {
		t.Error("1MB are 1024 KBytes but got ", b.KBytes())
	}
	if b.MBytes() != 1 {
		t.Error("1MB are 1 MBytes but got ", b.MBytes())
	}
	if b.GBytes() != 0.0009765625 {
		t.Error("1MB are 0.0009765625 GBytes but got ", b.GBytes())
	}
	fmt.Println(b.YBytes())
}
