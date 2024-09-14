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

func deleteData(keys []string) string {
	res := "keys:"
	for _, key := range keys {
		_, exists := db[key]
		if(exists) {
			delete(db, key)
			res = fmt.Sprintf("%s '%s', ", res, key)
		}
	}

	if len(res) > 5 {
		res = res[:len(res)-2]
	}
	res = fmt.Sprintf("%v have been deleted.", res)
	return res
}