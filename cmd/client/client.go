package main

import (
	"context"
	"github.com/Sokol111/category-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("192.168.49.2:30000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("did not connect: ", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			log.Printf("failed to close connection: %s", e)
		}
	}()
	client := proto.NewCategoryServiceClient(conn)

	category, err := client.CreateCategory(context.Background(), &proto.CreateCategoryRequest{Name: "Category 1"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Category: %v", category)
	category, err = client.UpdateCategory(context.Background(), &proto.UpdateCategoryRequest{Name: "Category 2", Id: category.Id, Version: 0})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Category: %v", category)
}
