package main

var handlers = map[string]func(v []Value) Value {
	"ping": ping,
	"set": set,
	"get": get,
}

func ping(args []Value) Value {
	if (len(args) > 0) {
		return Value{typ:STRING_TEXT, str:args[0].bulk}
	}
	return Value{typ:STRING_TEXT, str:"PONG"}
}

func set(args []Value) Value {
	response :="Set Command doesn't have right arguments!"
	if len(args) == 2 {
		key := args[0].bulk
		val := args[1].bulk
		putData(key, val)
		response = "Data added successfully!"
	}
	return Value{typ:"string", str:response}
}

func get(args []Value) Value {
	response :="Get Command doesn't have right arguments!"
	if len(args) == 1 {
		key := args[0].bulk
		val := fetchData(key)
		response = val
	}
	return Value{typ:"string", str:response}
}