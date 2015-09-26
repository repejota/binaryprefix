// Copyright 2015 Raül Pérez. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryprefix

import (
	"strconv"
)

// Prefixes
const (
	_          = iota             // ignore first value by assigning to blank identifier
	GB float64 = 1 << (10 * iota) // Gigabytes
	TB                            // Terabytes
	PB                            // Petabytes
	EB                            // Exabytes
	ZB                            // Zettabytes
	YB                            // Yottabytes
)

func getValues(size string) (string, float64, error) {
	d := size[len(size)-2:]
	s := size[:len(size)-2]
	f, err := strconv.ParseFloat(s, 64)
	return d, f, err
}

// BinaryPrefix ...
type BinaryPrefix float64

// ParseBinaryPrefix ...
func ParseBinaryPrefix(size string) (b BinaryPrefix, err error) {
	format, quantity, err := getValues(size)
	if err != nil {
		b = 0
		return b, err
	}
	switch format {
	case "B":
		b = BinaryPrefix(quantity / 1024 / 1024)
	case "KB":
		b = BinaryPrefix(quantity / 1024)
	case "MB":
		b = BinaryPrefix(quantity)
	case "GB":
		b = BinaryPrefix(quantity * GB)
	case "TB":
		b = BinaryPrefix(quantity * TB)
	case "PB":
		b = BinaryPrefix(quantity * PB)
	case "EB":
		b = BinaryPrefix(quantity * EB)
	case "ZB":
		b = BinaryPrefix(quantity * ZB)
	case "YB":
		b = BinaryPrefix(quantity * YB)
	}
	return b, nil
}

// Bytes get the number of bytes
func (b BinaryPrefix) Bytes() float64 {
	return float64(b) * 1024 * 1024
}

// KBytes get the number of kylo bytes
func (b BinaryPrefix) KBytes() float64 {
	return float64(b) * 1024
}

// MBytes get the number of mega bytes
func (b BinaryPrefix) MBytes() float64 {
	return float64(b)
}

// GBytes get the number of giga bytes
func (b BinaryPrefix) GBytes() float64 {
	return float64(b) / GB
}

// TBytes get the number of tera bytes
func (b BinaryPrefix) TBytes() float64 {
	return float64(b) / TB
}

// PBytes get the number of peta bytes
func (b BinaryPrefix) PBytes() float64 {
	return float64(b) / PB
}

// EBytes get the number of exa bytes
func (b BinaryPrefix) EBytes() float64 {
	return float64(b) / EB
}

// ZBytes get the number of zetta bytes
func (b BinaryPrefix) ZBytes() float64 {
	return float64(b) / ZB
}

// YBytes get the number of yotta bytes
func (b BinaryPrefix) YBytes() float64 {
	return float64(b) / YB
}
