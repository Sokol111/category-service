package main

import (
	"context"
	"github.com/Sokol111/category-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("did not connect: ", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			log.Printf("failed to close connection: %s", e)
		}
	}()
	client := proto.NewCategoryServiceClient(conn)

	category, _ := client.CreateCategory(context.Background(), &proto.CreateCategoryRequest{Name: "Category 1"})
	log.Printf("Category: %v", category)
	category, err = client.UpdateCategory(context.Background(), &proto.UpdateCategoryRequest{Name: "Category 2", Id: category.Id, Version: 0})
	if err != nil {
		log.Printf("%v\n", err)
	}
}
