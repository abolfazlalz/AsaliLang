package project

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"
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

func (*BuildInMethod) time() time.Time {
	return time.Now()
}

func (*BuildInMethod) format(args ...interface{}) string {
	if len(args) < 2 {
		panic("Few arguments")
	}
	if reflect.TypeOf(args[0]).String() == "time.Time" {
		return args[0].(time.Time).Format(args[1].(string))
	}
	return ""
}
