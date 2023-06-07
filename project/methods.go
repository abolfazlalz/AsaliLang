package project

import (
	"fmt"
	"math"
	"strings"
)

type BuildInMethod struct {
}

func NewBuildInMethod() *BuildInMethod {
	return &BuildInMethod{}
}

func (*BuildInMethod) print(args ...interface{}) {
	argsStr := make([]string, len(args))
	for i, arg := range args {
		argsStr[i] = fmt.Sprintf("%v", arg)
	}
	fmt.Println(strings.Join(argsStr, " "))
}

func (*BuildInMethod) sin(args ...interface{}) float64 {
	if len(args) < 1 {
		panic("Few arguments")
	}
	return math.Sin(args[0].(float64))
}
