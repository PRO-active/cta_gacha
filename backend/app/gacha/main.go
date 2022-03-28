package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pro-active/cta_gacha/app/gacha/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// dsn := "root:password@tcp(127.0.0.1:16002)/tutorial?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:password@tcp(db:3306)/tutorial?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := connectDB()
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
	if err := server.Start("0.0.0.0:80"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func connectDB() (*gorm.DB, error) {
	dbConf := newDBConfig()
	return gorm.Open(mysql.Open(dbConf.dsn()), &gorm.Config{})
}

type dbConfig struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func newDBConfig() dbConfig {
	return dbConfig{
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Database: os.Getenv("DATABASE"),
	}
}

func (c *dbConfig) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
}
