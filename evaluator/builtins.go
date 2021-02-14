package evaluator

import (
	"fmt"
	"strings"

	"github.com/lusingander/monkey/object"
)

var builtins = map[string]*object.Builtin{
	"puts":    {Fn: builtinPuts},
	"print":   {Fn: builtinPrint},
	"println": {Fn: builtinPrintln},
	"len":     {Fn: builtinLen},
	"first":   {Fn: builtinFirst},
	"last":    {Fn: builtinLast},
	"rest":    {Fn: builtinRest},
	"push":    {Fn: builtinPush},
}

func builtinPuts(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}
	return NULL
}

func builtinPrint(args ...object.Object) object.Object {
	strs := []string{}
	for _, arg := range args {
		strs = append(strs, arg.Inspect())
	}
	fmt.Print(strings.Join(strs, " "))
	return NULL
}

func builtinPrintln(args ...object.Object) object.Object {
	args = append(args, &object.String{Value: "\n"})
	return builtinPrint(args...)
}

func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments: want=1, got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to 'len' not supported: got=%s", arg.Type())
	}
}

func builtinFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments: want=1, got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Array:
		if len(arg.Elements) > 0 {
			return arg.Elements[0]
		}
		return NULL
	default:
		return newError("argument to 'first' not supported: got=%s", arg.Type())
	}
}

func builtinLast(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments: want=1, got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Array:
		l := len(arg.Elements)
		if l > 0 {
			return arg.Elements[l-1]
		}
		return NULL
	default:
		return newError("argument to 'last' not supported: got=%s", arg.Type())
	}
}

func builtinRest(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments: want=1, got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Array:
		l := len(arg.Elements)
		if l > 0 {
			newElems := make([]object.Object, l-1, l-1)
			copy(newElems, arg.Elements[1:])
			return &object.Array{Elements: newElems}
		}
		return NULL
	default:
		return newError("argument to 'rest' not supported: got=%s", arg.Type())
	}
}

func builtinPush(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments: want=2, got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Array:
		l := len(arg.Elements)
		newElems := make([]object.Object, l+1, l+1)
		copy(newElems, arg.Elements)
		newElems[l] = args[1]
		return &object.Array{Elements: newElems}
	default:
		return newError("argument to 'push' not supported: got=%s", arg.Type())
	}
}
