package main

import (
	"grpc-microservices/internal/proto-files/domain"
	"grpc-microservices/internal/proto-files/service"
	"fmt"
	"log"
	"net/http"
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
)

func main() {
	serverAddress := "localhost:7000"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := service.NewRepositoryServiceClient(conn)

	g:=gin.Default()
	g.GET("/example/", func(ctx *gin.Context){
		
		for i := range [10]int{} {
			repositoryModel := domain.Repository{
				Id:        int64(i),
				IsPrivate: true,
				Name:      string("Grpc-Demo"),
				UserId:    1245,
			}
	
			if responseMessage, e := client.Add(ctx, &repositoryModel); e != nil {
				panic(fmt.Sprintf("Was not able to insert Record %v", e))
				
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"result": fmt.Sprint(responseMessage),
				})
			}
		}
	
	})

	if err:=g.Run(":8080"); err!=nil{
		log.Fatalf("Failed	to run server: %v", err)
	}
	
}
