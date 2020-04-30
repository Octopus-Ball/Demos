package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)


// 连接Mogo
func link(addr string) (c *mongo.Client, err error) {
		// 设置客户端连接配置
		clientOptions := options.Client().ApplyURI("mongodb://"+ addr +":27017")

		// 连接到MongoDB
		c, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
			return
		}
	
		// 检查连接
		err = c.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Connected to MongoDB!")
		return
}

// 断开Mogo连接
func unLink(c *mongo.Client) (err error){
	err = c.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
	return
}

type Student struct {
	Name string
	Age int
}


func main() {
	client, err:= link("zy.u")
	if err != nil {
		return	
	}
}
