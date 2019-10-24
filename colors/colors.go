package colors

import (
	"fmt"
	"log"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
)

// logで色付き文字を出力する為の準備
func init() {
	log.SetOutput(colorable.NewColorableStdout())
}

// 色付き文字を出力
func Println(color, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(colorable.NewColorableStdout(), Attach(color, format, a...))
}

// 色付き文字列を返す
func Attach(color, format string, a ...interface{}) string {
	str := fmt.Sprintf(format, a...)
	v := aurora.Reset(str)

	if confirmExist(color, "Bold") {
		v = aurora.Bold(v)
	}
	if confirmExist(color, "Italic") {
		v = aurora.Italic(v)
	}
	if confirmExist(color, "Underline") {
		v = aurora.Underline(v)
	}

	bright := confirmExist(color, "Bright")
	bg := confirmExist(color, "Bg")

	if confirmExist(color, "Black") {
		v = black(v, bright, bg)
	} else if confirmExist(color, "Red") {
		v = red(v, bright, bg)
	} else if confirmExist(color, "Green") {
		v = green(v, bright, bg)
	} else if confirmExist(color, "Yellow") {
		v = yellow(v, bright, bg)
	} else if confirmExist(color, "Blue") {
		v = blue(v, bright, bg)
	} else if confirmExist(color, "Magenta") {
		v = magenta(v, bright, bg)
	} else if confirmExist(color, "Cyan") {
		v = cyan(v, bright, bg)
	} else if confirmExist(color, "White") {
		v = white(v, bright, bg)
	}

	return v.String()
}

// 指定された文字が文字列にあるか確認
func confirmExist(color, format string) bool {
	return strings.Index(color, format) != -1
}
