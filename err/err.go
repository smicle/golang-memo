package err

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"colors"
	"err/errType"

	"github.com/pkg/errors"
)

// エラーなら出力し終了、でなければ何もしない
var Found = func(e error, errType string) {
	if e == nil {
		return
	}

	err := fmt.Errorf(colors.Attach("Red", "%v", e))
	log.Fatalf("%+v\n", errors.Wrap(err, colors.Attach("Yellow", "Error occured: %v", errType)))
}

// err.Foundを使う関数が、errorを返すようにラップする
var Wrapper = func(f func() interface{}) (interface{}, error) {
	r, w, e := os.Pipe()
	Found(e, errType.OSPipe)

	stdout := os.Stdout
	os.Stdout = w

	v := f()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return Branch(v)
}

// 値がある場合と、ない場合の戻り値を作って返す
var Branch = func(v interface{}) (interface{}, error) {
	if v == "" {
		return nil, errors.New("ErrFound")
	} else {
		return v, nil
	}
}
