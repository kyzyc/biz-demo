package main

import (
	"github.com/joho/godotenv"
	"github.com/kyzyc/biz-demo/demo/demo_proto/biz/dal"
	"github.com/kyzyc/biz-demo/demo/demo_proto/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/demo/demo_proto/biz/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dal.Init()
	// CURD
	//mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "123456"})
	//mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("Password", "22222")

	//var row model.User
	//mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)
	//
	//fmt.Printf("row: %+v\n", row)
	//mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})

	mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})
}
