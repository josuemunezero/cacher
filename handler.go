package main

import (
	"fmt"
)

var handlers = map[string]func(v []Value) Value {
	"ping": ping,
	"set": set,
	"get": get,
	"del": del,
	"hset": hset,
	"hget": hget,
	"hgetall":hgetall,
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

func hset(args []Value) Value {
	response := fmt.Sprintf("HSET %s", INVALID_ARGUMENTS)
	if len(args) == 3 {
		putData(args[0].bulk, args[1].bulk, args[2].bulk)
		response = "Data added successfully!"
	}
	return Value{typ:"string", str:response}
}

func hget(args []Value) Value {
	response := fmt.Sprintf("HGET %s", INVALID_ARGUMENTS)
	if len(args) == 2 {
		response = fetchData(args[0].bulk, args[1].bulk)
	}
	return Value{typ:"string", str:response}
}

func hgetall(args []Value) Value {
	res := Value{typ:STRING_TEXT}

	if len(args) == 1 {
		data := fetchAll(args[0].bulk)
		mapStr := "{"
		for k, v := range data {
			mapStr = fmt.Sprintf("%v\"%v\":\"%v\",", mapStr, k, v)
		}
		if len(mapStr) > 1 {
			mapStr = mapStr[:len(mapStr)-1]
		}
		mapStr += "}"
		res.str = mapStr
	} else {
		res.str = fmt.Sprintf("HGETALL %s", INVALID_ARGUMENTS)
	}
	return res
}

func del(args []Value) Value {
	response := fmt.Sprintf("DEL %s", INVALID_ARGUMENTS)
	if(len(args) > 0) {
		response = deleteData(args[0].bulk)
	}
	return Value{typ:"string", str:response}
}