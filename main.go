package main

import (
	"github.com/abrarnaim015/KP-golang-tgs7/config"
	"github.com/abrarnaim015/KP-golang-tgs7/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	config.InitDB()
	e := routers.New()
	e.Logger.Fatal(e.Start(":8000"))
}
