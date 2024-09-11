package main

import (
	"fmt"
	"sync"
)


var db map[string]string
var hDb map[string]map[string]string

var dbMutex = sync.RWMutex{}
var hDbMutex = sync.RWMutex{}


func initDb() {
	db = make(map[string]string)
	hDb = make(map[string]map[string]string)
}

func putData(args... string) {
	switch len(args) {
		case 2:
			dbMutex.Lock()
			db[args[0]] = args[1]
			dbMutex.Unlock()
		case 3: 
			hDbMutex.Lock()
			if hDb[args[0]] == nil {
				hDb[args[0]] = make(map[string]string)
			}
			hDb[args[0]][args[1]] = args[2]
			hDbMutex.Unlock()
	}
	
}

func fetchData(args... string) string {
	var val string
	switch len(args) {
		case 1: 
			dbMutex.RLock()
			val = db[args[0]]
			dbMutex.RUnlock()
		case 2: 
			hDbMutex.RLock()
			val = hDb[args[0]][args[1]]
			hDbMutex.RUnlock()
		default:
			val = ""
	} 
	if val == "" {
		return "Key not found!"
	}
	return val
}

func fetchAll(args... string) map[string]string {
	val := make(map[string]string)
	hDbMutex.Lock()
	if v, ok := hDb[args[0]]; ok {
		val = v
	}
	hDbMutex.Unlock()
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