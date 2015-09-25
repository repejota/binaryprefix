package binaryprefix

import (
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
}

func TestParseBinaryPrefixGB(t *testing.T) {
	b, err := ParseBinaryPrefix("4GB")
	if err != nil {
		t.Error("Error parsing 4GB")
	}
	if int(b) != 4294967296 {
		t.Error("4GB are 4294967296 Bytes but got ", int(b))
	}
}

func TestParseBinaryPrefixTB(t *testing.T) {
	b, err := ParseBinaryPrefix("3TB")
	if err != nil {
		t.Error("Error parsing 3TB")
	}
	if int(b) != 3298534883328 {
		t.Error("3TB are 3298534883328 Bytes but got ", int(b))
	}
}
