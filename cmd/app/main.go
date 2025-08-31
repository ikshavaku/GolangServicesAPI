package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ikshavaku/catalogue/api"
	"github.com/ikshavaku/catalogue/utils"
)

func main() {
	err := utils.InitConfig()
	if err != nil {
		log.Printf("%+v\n", err)
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s := api.NewAPIServer(InjectServicesController())
	runWebserver(ctx, cancel, s)
	<-ctx.Done()
	log.Println("Exiting...")
}

func runWebserver(ctx context.Context, cancel context.CancelFunc, server *gin.Engine) {
	go func() {
		err := server.Run(fmt.Sprintf(":%d", utils.GetConfig().Server.Port))
		if err != nil {
			panic(err)
		}
		log.Println("Canceling context as server exited")
		cancel()
	}()
}
