package project

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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

func (*BuildInMethod) input(args ...interface{}) string {
	if len(args) > 0 {
		fmt.Print(args[0])
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func (*BuildInMethod) sin(args ...interface{}) float64 {
	if len(args) < 1 {
		panic("Few arguments")
	}
	return math.Sin(args[0].(float64))
}

func (*BuildInMethod) float(args ...interface{}) float64 {
	if len(args) < 1 {
		panic("Few arguments")
	}

	numStr := fmt.Sprintf("%v", args[0])
	result, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		panic(err)
	}
	return result
}
