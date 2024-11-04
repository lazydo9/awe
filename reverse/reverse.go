// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package reverse can reverse things, particularly strings.
package reverse

import "fmt"

// String returns its argument string reversed rune-wise left to right.
func String(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println("私货!!")
	fmt.Println("这里是v0.0.2")
	return string(r)
}

func Hello() {
	fmt.Println("Hello!")
}
