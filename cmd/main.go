package main

import (
	"myapp/pkg/db"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// func main() {
// 	err := Run("myapp")
// 	if err != nil {
// 		fmt.Print("failed to start the server")
// 	}

// }

func main() {
	_, err := db.Test()
	if err != nil {
		println("Error while running migrations")
	}
}
