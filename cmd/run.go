package main

import (
	"context"
	"fmt"
	"myapp/pkg/db"
	"myapp/pkg/server"
	"net/http"
)

type App struct {
	app    *Appcontext
	server *server.Server
}

type Appcontext struct {
	DBClient *db.Client
}

func Run(configPath string) error {

	fmt.Println("Hi Naresh, Welcome to Golang Run World")
	ctx := context.Background()

	app := App{}
	initErr := app.Init(ctx, configPath)
	if initErr != nil {
		panic(initErr)
	}

	func() {
		if err := app.server.Start(); err != nil && err != http.ErrServerClosed {
			fmt.Print("Error while running the server: ", err.Error())
			panic(err.Error())
		}
	}()
	return nil
}

func (a *App) Init(ctx context.Context, configpath string) error {
	Config, err := db.LoadDbconfig()
	if err != nil {
		fmt.Print("Error while declaring Config in Init: ", err.Error())
		panic(err.Error())
	}
	DbClient, err := db.InitDB(Config)
	if err != nil {
		fmt.Print("Error while declaring DbClient in Init: ", err.Error())
		panic(err.Error())
	}
	AppContext := &Appcontext{DBClient: DbClient}
	server, err := server.NewServer("127.0.0.1", 8080)
	if err != nil {
		fmt.Print("Error while declaring NewServer in Init: ", err.Error())
		panic(err.Error())
	}
	a.app = AppContext
	a.server = server
	return nil
}
