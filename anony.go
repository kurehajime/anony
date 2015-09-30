// anony.go
package anony

import (
	"github.com/ikawaha/kagome/tokenizer"
	"regexp"
	"unicode/utf8"
)

var (
	t = tokenizer.New()
)

func init() {
	tokenizer.SysDic()
}
func Anony(text string, single bool) string {
	tokens := t.Tokenize(text)
	var rText string
	var IniCount int
	for j := 0; j < len(tokens); j++ {
		tk := tokens[j]
		ft := tk.Features()
		if len(ft) > 7 {
			if ft[2] == "人名" && ft[1] == "固有名詞" {
				if IniCount == 0 {
					rText += Word2initial(ft[7])
				} else if IniCount == 1 && single == false {
					rText += "・"
					rText += Word2initial(ft[7])
				}
				IniCount++
			} else {
				rText += tk.Surface
				IniCount = 0
			}
		} else if len(ft) > 0 {
			rText += tk.Surface
			IniCount = 0
		}
	}
	return rText
}
func Word2initial(kana string) string {
	r, _ := utf8.DecodeRune([]byte(kana))
	ini := string(r)

	ini = regexp.MustCompile("[アァ]").ReplaceAllString(ini, "A")
	ini = regexp.MustCompile("[イィ]").ReplaceAllString(ini, "I")
	ini = regexp.MustCompile("[ウゥ]").ReplaceAllString(ini, "U")
	ini = regexp.MustCompile("[エェ]").ReplaceAllString(ini, "E")
	ini = regexp.MustCompile("[オォ]").ReplaceAllString(ini, "O")
	ini = regexp.MustCompile("[カキクケコ]").ReplaceAllString(ini, "K")
	ini = regexp.MustCompile("[サシスセソ]").ReplaceAllString(ini, "S")
	ini = regexp.MustCompile("[タツテト]").ReplaceAllString(ini, "T")
	ini = regexp.MustCompile("[チ]").ReplaceAllString(ini, "T")
	ini = regexp.MustCompile("[ナニヌネノ]").ReplaceAllString(ini, "N")
	ini = regexp.MustCompile("[ハヒヘホ]").ReplaceAllString(ini, "H")
	ini = regexp.MustCompile("[フ]").ReplaceAllString(ini, "F")
	ini = regexp.MustCompile("[マミムメモ]").ReplaceAllString(ini, "M")
	ini = regexp.MustCompile("[ヤユヨ]").ReplaceAllString(ini, "Y")
	ini = regexp.MustCompile("[ラリルレロ]").ReplaceAllString(ini, "R")
	ini = regexp.MustCompile("[ワヲ]").ReplaceAllString(ini, "W")
	ini = regexp.MustCompile("[ン]").ReplaceAllString(ini, "N")
	ini = regexp.MustCompile("[ガギグゲゴ]").ReplaceAllString(ini, "G")
	ini = regexp.MustCompile("[ザズゼゾ]").ReplaceAllString(ini, "Z")
	ini = regexp.MustCompile("[ダヂヅデド]").ReplaceAllString(ini, "D")
	ini = regexp.MustCompile("[ジ]").ReplaceAllString(ini, "J")
	ini = regexp.MustCompile("[パピプペポ]").ReplaceAllString(ini, "P")
	ini = regexp.MustCompile("[バビブベボ]").ReplaceAllString(ini, "B")

	return ini
}
