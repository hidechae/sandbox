package main

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	var upper Upper
	sr := strings.NewReader("Hello, World")
	r := transform.NewReader(sr, &upper)

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
}

// Upper は入力を大文字に変換するTransformer
type Upper struct {
	// Resetを実装しなくて良い
	transform.NopResetter
}

// Transform
// atEOFは、境界処理が必要な場合に考慮が必要になる
// nDstとnSrcの値が変わるケースは、バイト列の長さが異なる変換を行う場合
func (u Upper) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	if len(dst) == 0 {
		return 0, 0, transform.ErrShortDst
	}
	if len(src) == 0 {
		return 0, 0, transform.ErrShortSrc
	}

	n := min(len(src), len(dst))
	for i := range n {
		v := src[i]
		if v >= 'a' && v <= 'z' {
			dst[i] = v - ('a' - 'A')
		} else {
			dst[i] = v
		}
	}

	return n, n, nil
}

func (u Upper) Reset() {

}

// KatakanaToWide はカタカナなら全角にするTransformer
func KatakanaToWide() transform.Transformer {
	return runes.If(runes.In(unicode.Katakana), width.Widen, nil)
}
