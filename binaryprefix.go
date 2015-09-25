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
type BinaryPrefix int

// ParseBinaryPrefix ...
func ParseBinaryPrefix(size string) (b BinaryPrefix, err error) {
	format, quantity, err := getValues(size)
	if err != nil {
		b = 0
		return b, err
	}
	switch format {
	case "B":
		b = BinaryPrefix(quantity)
	case "KB":
		b = BinaryPrefix(quantity * 1024)
	case "MB":
		b = BinaryPrefix(quantity * 1024 * 1024)
	case "GB":
		b = BinaryPrefix(quantity * GB * 1024 * 1024)
	case "TB":
		b = BinaryPrefix(quantity * TB * 1024 * 1024)
	case "PB":
		b = BinaryPrefix(quantity * PB * 1024 * 1024)
	case "EB":
		b = BinaryPrefix(quantity * EB * 1024 * 1024)
	case "ZB":
		b = BinaryPrefix(quantity * ZB * 1024 * 1024)
	case "YB":
		b = BinaryPrefix(quantity * YB * 1024 * 1024)
	}
	return b, nil
}
