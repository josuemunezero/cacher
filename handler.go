package main

var handlers = map[string]func(v []Value) Value {
	"ping": ping,
}

func ping(args []Value) Value {
	if (len(args) > 0) {
		return Value{typ:STRING_TEXT, str:args[0].bulk}
	}
	return Value{typ:STRING_TEXT, str:"PONG"}
}
