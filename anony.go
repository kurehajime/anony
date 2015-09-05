// anony.go
package main

import (
	"github.com/ikawaha/kagome"
	"unicode/utf8"
	"regexp"
)
var (
	t = kagome.NewTokenizer()
)
func Anony(text string)string{
	return text
}
func Word2initial(kana string)string{
	r,_ :=utf8.DecodeRune([]byte(kana))
	ini:=string(r)
	
 	ini = regexp.MustCompile("[アァ]").ReplaceAllString(ini,"A")
 	ini = regexp.MustCompile("[イィ]").ReplaceAllString(ini,"I")
 	ini = regexp.MustCompile("[ウゥ]").ReplaceAllString(ini,"U")
 	ini = regexp.MustCompile("[エェ]").ReplaceAllString(ini,"E")
 	ini = regexp.MustCompile("[オォ]").ReplaceAllString(ini,"O")
 	ini = regexp.MustCompile("[カキクケコ]").ReplaceAllString(ini,"K")
 	ini = regexp.MustCompile("[サシスセソ]").ReplaceAllString(ini,"S")
 	ini = regexp.MustCompile("[タチツテト]").ReplaceAllString(ini,"T")
 	ini = regexp.MustCompile("[ナニヌネノ]").ReplaceAllString(ini,"N")
 	ini = regexp.MustCompile("[マミムメモ]").ReplaceAllString(ini,"M")
 	ini = regexp.MustCompile("[ヤユヨ]").ReplaceAllString(ini,"Y")
 	ini = regexp.MustCompile("[ラリルレロ]").ReplaceAllString(ini,"R")
 	ini = regexp.MustCompile("[ワヲ]").ReplaceAllString(ini,"W")
 	ini = regexp.MustCompile("[ン]").ReplaceAllString(ini,"N")
 	ini = regexp.MustCompile("[ガギグゲゴ]").ReplaceAllString(ini,"G")
 	ini = regexp.MustCompile("[ザズゼゾ]").ReplaceAllString(ini,"Z")
 	ini = regexp.MustCompile("[ジ]").ReplaceAllString(ini,"J")
 	ini = regexp.MustCompile("[パピプペポ]").ReplaceAllString(ini,"P")
 	ini = regexp.MustCompile("[バビブベボ]").ReplaceAllString(ini,"B")
 	
	return ini
}