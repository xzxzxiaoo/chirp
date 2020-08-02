package main

import (
	"chirp/src/conf"
	"chirp/src/router"
	"fmt"
)

func main() {

	// config init
	c := conf.Configured()

	// db init
	fateapi := &conf.ApiServerConfiguration{C_dbConfig: c.MysqlConnector, ListenPort: c.ListenPort}
	fateapi.InitDatabaseConnecter()

	// route init
	r := router.Init(fateapi)
	if err := r.Run(fmt.Sprintf(":%v", fateapi.ListenPort)); err != nil {
		panic(err)
	}
}
