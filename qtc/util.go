package main

import (
	"bytes"
	"path/filepath"
	"unicode"
)

// mangleSuffix is used for mangling quicktemplate-specific names
// in the generated code, so they don't clash with user-provided names.
const mangleSuffix = "422016"

func stripLeadingSpace(b []byte) []byte {
	for len(b) > 0 && isSpace(b[0]) {
		b = b[1:]
	}
	return b
}

func stripTrailingSpace(b []byte) []byte {
	for len(b) > 0 && isSpace(b[len(b)-1]) {
		b = b[:len(b)-1]
	}
	return b
}

func collapseSpace(b []byte) []byte {
	return stripSpaceExt(b, true)
}

func stripSpace(b []byte) []byte {
	return stripSpaceExt(b, false)
}

func stripSpaceExt(b []byte, isCollapse bool) []byte {
	if len(b) == 0 {
		return b
	}

	isLastSpace := isSpace(b[len(b)-1])
	var dst []byte
	for len(b) > 0 {
		n := bytes.IndexByte(b, '\n')
		if n < 0 {
			n = len(b)
		}
		z := b[:n]
		if n == len(b) {
			b = b[n:]
		} else {
			b = b[n+1:]
		}
		z = stripLeadingSpace(z)
		z = stripTrailingSpace(z)
		if len(z) == 0 {
			continue
		}
		dst = append(dst, z...)
		if isCollapse {
			dst = append(dst, ' ')
		}
	}
	if isCollapse && !isLastSpace && len(dst) > 0 {
		dst = dst[:len(dst)-1]
	}
	return dst
}

func isSpace(c byte) bool {
	return unicode.IsSpace(rune(c))
}

func isUpper(c byte) bool {
	return unicode.IsUpper(rune(c))
}

func getPackageName(filename string) (string, error) {
	filenameAbs, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}
	dir, _ := filepath.Split(filenameAbs)
	return filepath.Base(dir), nil
}
