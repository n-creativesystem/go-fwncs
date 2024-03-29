// Copyright 2020 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// at https://github.com/gin-gonic/gin/blob/master/LICENSE

package bytesconv

import (
	"unsafe"
)

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
