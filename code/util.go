package util

import (
	"fmt"
	"os"
	"time"

	"err"
	"err/errType"
)

// 三項演算子
func TernaryOperator(b bool, x, y interface{}) interface{} {
	if b {
		return x
	} else {
		return y
	}
}

// 環境変数を取得
func TryGetEnv(env string) (v string) {
	v, ok := os.LookupEnv(env)
	if !ok {
		e := fmt.Errorf("Missing environment variable [%s]", env)
		err.ErrorFound(e, errType.GetEnv)
	}
	return
}

// 現在時刻を取得しファイル名に用いるフォーマットに変換
func GetCurrentTimeByString() string {
	ct := time.Now()
	return fmt.Sprintf("%d_%d_%d_%d", ct.Year(), ct.Month(), ct.Day(), ct.Hour())
}
