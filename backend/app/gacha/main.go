package main

import (
	"log"
	"net/http"

	"github.com/pro-active/cta_gacha/app/gacha/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// dsn := "root:password@tcp(127.0.0.1:16002)/tutorial?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:password@tcp(db:3306)/tutorial?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	/*
		userRepo := user.NewUserRepository(db)

		user, err := userRepo.CreateUser("001", "Name", "password", "a@gmail.com", "hash")
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("user:%v", user)
	*/

	server := api.InitalizeServer(db)
	// server.Server.ListenAndServe()
	// firewallが出てくるのでローカルでホストを持たせる
	if err := server.Start("0.0.0.0:8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
