package bloblang

import (
	"errors"
	"fmt"

	"github.com/Jeffail/benthos/v3/public/bloblang"
)

func init() {
	bloblang.RegisterFunction("crazy_object", func(args ...interface{}) (bloblang.Function, error) {
		// Expects a single integer argument
		if len(args) != 1 {
			return nil, errors.New("oi! I expected an integer argument")
		}
		iArg, ok := args[0].(int64)
		if !ok {
			return nil, errors.New("oi! I expected an integer argument")
		}
		return func() (interface{}, error) {
			obj := map[string]interface{}{}
			for i := int64(0); i < iArg; i++ {
				obj[fmt.Sprintf("key%v", i)] = fmt.Sprintf("value%v", i)
			}
			return obj, nil
		}, nil
	})

	bloblang.RegisterMethod("into_object", func(args ...interface{}) (bloblang.Method, error) {
		// Expects a single string argument
		if len(args) != 1 {
			return nil, errors.New("oi! I expected a string argument")
		}
		sArg, ok := args[0].(string)
		if !ok {
			return nil, errors.New("oi! I expected a string argument")
		}
		return func(v interface{}) (interface{}, error) {
			return map[string]interface{}{
				sArg: v,
			}, nil
		}, nil
	})
}
