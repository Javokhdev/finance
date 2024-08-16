package main

import (
	"context"
	"fmt"
	"log"

	"api-gateway/api"
	"api-gateway/api/handler"
	pb "api-gateway/genproto"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// UserConn, err := grpc.NewClient(fmt.Sprintf("auth-service%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatal("Error while Newclient: ", err.Error())
	// }
	// defer UserConn.Close()

	Connect, err := grpc.NewClient(fmt.Sprintf("budget%s", ":8088"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while Newclient: ", err.Error())
	}
	defer Connect.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB: 0,
	})
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis successfully!")
	// auth := pb.NewAuthServiceClient(UserConn)
	// user := pb.NewUserServiceClient(UserConn)
	rediss := handler.NewInMemoryStorage(rdb)
	account := pb.NewAccountServiceClient(Connect)
	budget := pb.NewBudgetServiceClient(Connect)
	category := pb.NewCategoryServiceClient(Connect)
	goal := pb.NewGoalServiceClient(Connect)
	transaction := pb.NewTransactionServiceClient(Connect)
	notification := pb.NewNotificationtServiceClient(Connect)

	h := handler.NewHandler(account, budget, category, goal, transaction, notification, rediss)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
