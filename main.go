// @title            GoWork API
// @version        1.0
// @description    API para gerenciamento de vagas de trabalho
// @host            squid-app-ti4wd.ondigitalocean.app
// @BasePath        /api/v1
package main

import (
    "os"

    "github.com/viitorags/gowork/config"
    _ "github.com/viitorags/gowork/docs"
    "github.com/viitorags/gowork/router"
)

var (
    logger *config.Logger
)

func main() {

    logger = config.GetLogger("main")

    err := config.Init()
    if err != nil {
        logger.Error("config initalization error: ", err)
        os.Exit(1)
    }

    router.Initialize()
}
