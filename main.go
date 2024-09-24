package main

import (
	"fmt"
	"log"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/Kyadlapati/go-grpc-gateway/protogen/golang/orders"
	"github.com/Kyadlapati/go-grpc-gateway/protogen/golang/product"
)

func main() {
	OrderItem := orders.Orders{
		OrderId:    10,
		CustomerId: 11,
		IsActive:   true,
		OrderDate:  &date.Date{Year: 2021, Month: 1, Day: 1},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "CocoCola", ProductType: product.ProductType_DRINK},
		},
	}

	bytes, err := protojson.Marshal(&OrderItem)
	if err != nil {
		log.Fatal("deserialization error:", err)
	}
	fmt.Println(string(bytes))
}
