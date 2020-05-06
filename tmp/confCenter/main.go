package main

import (
	"confCenter/common"
	"confCenter/models"
	"fmt"
)

func main() {
	DB := common.GetDB()
	defer DB.Close()

	// m := models.MapConf{Anchor: 1001, Type: 1, UnLockLevel: 2}
	tmpm := models.MapConf{Anchor: 1001}
	DB.First(&tmpm)
	fmt.Println(tmpm)

	if rst := DB.NewRecord(tmpm); rst {
		DB.Create(&tmpm)
		fmt.Println("插入")
	} else {
		fmt.Println("未插入")
	}
}

/*
import (
	_ "confCenter/common"
	"confCenter/handler"
	"confCenter/subscriber"

	"confCenter/common"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	confCenter "confCenter/proto/confCenter"
)

func main() {
	defer common.GetDB().Close()

	// 新建服务
	service := micro.NewService(
		micro.Name(common.GetMicroConfig().Name),
		micro.Version(common.GetMicroConfig().Version),
		micro.Registry(common.GetEtcdRegistry()),
	)

	service.Init()

	// Initialise service
	service.Init()

	// Register Handler
	confCenter.RegisterConfCenterHandler(service.Server(), new(handler.ConfCenter))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.confCenter", service.Server(), new(subscriber.ConfCenter))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.confCenter", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
*/
