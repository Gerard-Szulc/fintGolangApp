package main

import (
	"fintGolangApp/api"
	"fintGolangApp/database"
	"fintGolangApp/migrations"
	"os"
)

func main() {
	argsWithProg := os.Args
	switch argsWithProg[1] {
	case "migrate":
		{
			migrations.Migrate()
			return
		}
	case "start":
		{
			database.InitDatabase()
			api.StartApi()
		}
	}

}
