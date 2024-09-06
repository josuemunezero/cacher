package main


var db map[string]string

func initDb() {
	db = make(map[string]string)
}

func putData(key, val string) {
	db[key] = val
}

func fetchData(key string) string {
	val := db[key]
	if val == "" {
		return "Key not found!"
	}
	return val
}