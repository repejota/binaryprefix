package binaryprefix

import (
	"testing"
)

func TestParseBinaryPrefixMB(t *testing.T) {
	b, err := ParseBinaryPrefix("1MB")
	if err != nil {
		t.Error("Error parsing 1MB")
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
	if b.TBytes() != 9.5367431640625e-07 {
		t.Error("1MB are 0.0009765625 TBytes but got ", b.TBytes())
	}
}

func TestParseBinaryPrefixGB(t *testing.T) {
	b, err := ParseBinaryPrefix("4GB")
	if err != nil {
		t.Error("Error parsing 4GB")
	}
	if int(b) != 4096 {
		t.Error("4GB are 4096 Bytes but got ", int(b))
	}
}

func TestParseBinaryPrefixTB(t *testing.T) {
	b, err := ParseBinaryPrefix("3TB")
	if err != nil {
		t.Error("Error parsing 3TB")
	}
	if int(b) != 3145728 {
		t.Error("3TB are 3145728 Bytes but got ", int(b))
	}
}

func TestParseBinaryPrefixPB(t *testing.T) {
	b, err := ParseBinaryPrefix("1PB")
	if err != nil {
		t.Error("Error parsing 1PB")
	}
	if int(b) != 1073741824 {
		t.Error("1PB are 1073741824 Bytes but got ", int(b))
	}
}

func TestParseBinaryPrefixEB(t *testing.T) {
	b, err := ParseBinaryPrefix("1EB")
	if err != nil {
		t.Error("Error parsing 1EB")
	}
	if int(b) != 1099511627776 {
		t.Error("1EB are 1099511627776 Bytes but got ", int(b))
	}
}

func TestParseBinaryPrefixZB(t *testing.T) {
	b, err := ParseBinaryPrefix("1ZB")
	if err != nil {
		t.Error("Error parsing 1ZB")
	}
	if int(b) != 1125899906842624 {
		t.Error("1ZB are 1125899906842624 Bytes but got ", int(b))
	}
}

func TestParseBinaryPrefixYB(t *testing.T) {
	b, err := ParseBinaryPrefix("0.00001YB")
	if err != nil {
		t.Error("Error parsing 0.1YB")
	}
	if int(b) != 11529215046068 {
		t.Error("0.1YB are 11529215046068 Bytes but got ", int(b))
	}
}
