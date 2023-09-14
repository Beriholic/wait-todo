package main

import (
	"wait-to-do/app/user/repository/db/dao"
	"wait-to-do/config"
)

func main() {
	config.Init()
	dao.InitDB()
}
