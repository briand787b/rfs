package main

import (
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func main() {
	// github.com/pkg/errors TRUE
	if err := func1(true); err != nil {
		fmt.Println("full error before conversion: ", err)
		fmt.Println("error cause by errors.Cause(): ", errors.Cause(err))
		fmt.Printf("full extended errors.Error: %+v\n", err)
		fmt.Println("DONE")
	}

	// github.com/pkg/errors FALSE
	if err := func1(false); err != nil {
		fmt.Println("full error before conversion: ", err)
		fmt.Println("error cause by errors.Cause(): ", errors.Cause(err))
		fmt.Printf("full extended errors.Error: %+v\n", err)
		fmt.Println("DONE")
	}

	// see what happens when normal err is printed with %+v
	fmt.Printf("full extended normal error: %+v\n", fmt.Errorf("this is a normal error"))

	l, _ := zap.NewDevelopment()
	sl := l.Sugar()

	sl.Errorw("the message from the app",
		"stack_trace", fmt.Sprintf("%+v\n", func1(true)),
	)
}

func func1(pkgErr bool) error {
	if err := func2(pkgErr); err != nil {
		return errors.Wrap(err, "error in func1")
	}

	return nil
}

func func2(pkgErr bool) error {
	if err := func3(pkgErr); err != nil {
		return errors.Wrap(err, "error in func2")
	}

	return nil
}

func func3(pkgErr bool) error {
	if pkgErr {
		return errors.New("this error was created using github.com/pkg/errors.New()")
	}

	return fmt.Errorf("this error was caused by fmt.Errorf")
}
