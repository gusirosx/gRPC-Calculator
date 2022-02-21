package main

import (
	"fmt"
	pb "gRPC-Calculator/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Client pb.AddServiceClient

func main() {
	// Set up a connection to the server.
	conn, err := Connection()
	if err != nil {
		log.Printf("failed to dial server %s: %v", *serverAddr, err)
	}
	defer conn.Close()
	Client = pb.NewAddServiceClient(conn)
	// Set up a http server.
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		fmt.Fprintln(ctx.Writer, "Up and running...")
	})
	router.GET("/add/:a/:b", Add)
	router.GET("/multiply/:a/:b", Multiply)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// Add Function
func Add(ctx *gin.Context) {
	a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
		return
	}

	b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
		return
	}

	req := &pb.Request{A: int64(a), B: int64(b)}
	if response, err := Client.Add(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// Multiply Function
func Multiply(ctx *gin.Context) {
	a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
		return
	}
	b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
		return
	}
	req := &pb.Request{A: int64(a), B: int64(b)}

	if response, err := Client.Multiply(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
