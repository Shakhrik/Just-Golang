package main

import (
	"fmt"
	"log"
	"strconv"

	"net/http"

	"github.com/Shakhrik/Just-Golang/grpc_learn/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, er := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if er != nil {
		log.Println("ERrorrrrr")
		panic(er)
	}
	client := proto.NewAddServiceClient(conn)
	router := gin.Default()
	router.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter for B"})
			return
		}
		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error fa": err.Error()})
		}
	})
	router.GET("/multip/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter for B"})
			return
		}
		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := client.Multiply(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server on 8080: %v", err)
	}
}
