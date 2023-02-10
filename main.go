package main

import (
	"BaigeiCode/yearbook_api/db"
	"BaigeiCode/yearbook_api/routers"
)

func Init() error {
	err := db.Init()
	if err != nil {
		return err
	}
	return routers.SetupRouters()
}
func main() {
	err := Init()
	if err != nil {
		panic(err)
	}
}
