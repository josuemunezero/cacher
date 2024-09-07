package main

import "fmt"


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

func deleteData(key string) string {
	val, exists := db[key]
	if(exists) {
		delete(db, key)
		return fmt.Sprintf("{\"%s\": \"%s\"} has been deleted.", key, val)
	}
	return fmt.Sprintf("No Record with Key \"%s\" was found!", key)
}