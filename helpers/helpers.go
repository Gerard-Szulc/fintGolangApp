package helpers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)
	return string(hashed)
}

func Init() {
	// loads values from .env into the system
	env := os.Getenv("FINT_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	err := godotenv.Load()
	HandleErr(err)
}

func ConnectDB() *gorm.DB {
	Init()

	user, exists := os.LookupEnv("DBUSER")

	if exists {
		fmt.Println(user)
	}
	password, exists := os.LookupEnv("DBPASSWORD")

	if exists {
		fmt.Println(exists)
	}
	dbname, exists := os.LookupEnv("DBNAME")

	if exists {
		fmt.Println(exists)
	}

	dbargs := fmt.Sprintf("host=127.0.0.1 port=5432 user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	fmt.Println(dbargs)
	db, err := gorm.Open("postgres", dbargs)
	HandleErr(err)
	return db
}
