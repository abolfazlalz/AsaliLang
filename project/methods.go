package project

import (
	"fmt"
	"strings"
)

type BuildInMethod struct {
}

func NewBuildInMethod() *BuildInMethod {
	return &BuildInMethod{}
}

func (*BuildInMethod) print(args ...string) {
	fmt.Println(strings.Join(args, " "))
}
