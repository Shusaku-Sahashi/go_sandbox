package main

/*
 errorsパッケージでErrorハンドリングの方法をまとめた。
*/

import (
	"fmt"

	"github.com/pkg/errors"
)

type LowlayerError struct {
	message string
}

func (err LowlayerError) Error() string {
	return fmt.Sprintf("Message: %s", err.message)
}

type OuterError struct {
	inner   error
	message string
}

func (err OuterError) Error() string {
	return fmt.Sprintf("Message: %s", err.message)
}

func fn() error {
	// TODO: LowlayerErrorのStackTraceが取得出来ないので、出力方法を検討する。
	e1 := LowlayerError{message: "MyError is occur."}
	e2 := OuterError{message: "OuterError", inner: errors.Wrap(e1, "LowlayerError")}
	e3 := errors.Wrap(e2, "middle")
	return errors.Wrap(e3, "outer")
}

func CauseError() {
	err := fn()
	// 大元のErrorをで判断。
	switch errors.Cause(err).(type) {
	case OuterError:
		// errorsは%+vで stack traceを出力可能。
		fmt.Printf("%+v", err)
	default:
		fmt.Printf("%v", err)
	}
}

func main() {
	CauseError()
}
