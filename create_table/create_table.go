package main

import (
	"fmt"
	"gentest/dal"
	"gentest/dal/model"
)

func main()  {
	err := dal.DB.AutoMigrate(model.User{})
	if err != nil {
		fmt.Println("Create mysql table failed")
	} else {
		fmt.Println("Create mysql table successful")
	}
}