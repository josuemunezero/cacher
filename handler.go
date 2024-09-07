package main

import "fmt"

var handlers = map[string]func(v []Value) Value {
	"ping": ping,
	"set": set,
	"get": get,
	"del": del,
}

const (
	INVALID_ARGUMENTS="Command does not have required argument(s)!"
)

func ping(args []Value) Value {
	if (len(args) > 0) {
		return Value{typ:STRING_TEXT, str:args[0].bulk}
	}
	return Value{typ:STRING_TEXT, str:"PONG"}
}

func set(args []Value) Value {
	response := fmt.Sprintf("SET %s", INVALID_ARGUMENTS)
	if len(args) == 2 {
		key := args[0].bulk
		val := args[1].bulk
		putData(key, val)
		response = "Data added successfully!"
	}
	return Value{typ:"string", str:response}
}

func get(args []Value) Value {
	response := fmt.Sprintf("GET %s", INVALID_ARGUMENTS)
	if len(args) == 1 {
		key := args[0].bulk
		val := fetchData(key)
		response = val
	}
	return Value{typ:"string", str:response}
}

func del(args []Value) Value {
	response := fmt.Sprintf("DEL %s", INVALID_ARGUMENTS)
	if(len(args) > 0) {
		response = deleteData(args[0].bulk)
	}
	return Value{typ:"string", str:response}
}
