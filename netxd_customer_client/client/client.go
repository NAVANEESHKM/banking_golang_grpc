package main

import (
	"context"
	"fmt"

	"log"

	pb "netxd_grpc_mongo/netxd_customer"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &pb.Customer{
		BankID:566,
		Customer_Name: "sachin",
		Customer_ID : 50,
	    Balance:500,

	})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response.Balance )
	response1, err1 := client.UpdateCustomer(context.Background(), &pb.CustomerDetails{
		From_ID: 50,
		TO_ID:   60,
		Amount:  500,
	})
	if err1 != nil {
		log.Fatalf("Failed to call SayHello: %v", err1)
	}

	fmt.Printf("Response: %s\n", response1.From_ID)
}
