package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {}

func init() {
	fmt.Println(string(ModifierRegisterer), "loaded!!!")
}

var ModifierRegisterer = registerer("krakend-debugger")

type registerer string

func (r registerer) RegisterModifiers(f func(
	name string,
	factoryFunc func(map[string]interface{}) func(interface{}) (interface{}, error),
	appliesToRequest bool,
	appliesToResponse bool,
)) {
	f(string(r)+"-response", r.responseDump, false, true)
	fmt.Println(string(r), "registered!!!")
}

type ResponseWrapper interface {
	Data() map[string]interface{}
	Io() io.Reader
	IsComplete() bool
	StatusCode() int
	Headers() map[string][]string
}

var unknownTypeErr = errors.New("unknown request type")

func (r registerer) responseDump(
	cfg map[string]interface{},
) func(interface{}) (interface{}, error) {
	fmt.Println("response dumper injected!!!")
	return func(input interface{}) (interface{}, error) {
		resp, ok := input.(ResponseWrapper)
		if !ok {
			return nil, unknownTypeErr
		}

		fmt.Println("data:", resp.Data())
		fmt.Println("is complete:", resp.IsComplete())
		fmt.Println("headers:", resp.Headers())
		fmt.Println("status code:", resp.StatusCode())

		return input, nil
	}
}
