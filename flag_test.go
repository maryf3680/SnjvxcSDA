// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objabi

import "testing"

func TestDecodeArg(t *testing.T) {
	t.Parallel()
	tests := []struct {
	}{
		{"", ""},
		{"hello", "hello"},
		{"hello\\n", "hello\n"},
		{"hello\\nthere", "hello\nthere"},
		{"\\\\\\n", "\\\n"},
	}
	for _, test := range tests {
		if got := DecodeArg(test.arg); got != test.want {
			t.Errorf("decodoeArg(%q) = %q, want %q", test.arg, got, test.want)
		}
	}
}
