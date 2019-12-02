package rabbitmqv1

import (
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/micro/go-config"
	"github.com/micro/go-log"
)

func RabbitMQV1() broker.Broker {
	//加载配置项
	err := config.LoadFile("../common-config/config.json")
	if err != nil {
		log.Fatalf("Could not load config file: %s", err.Error())
	}

	return rabbitmq.NewBroker(
		broker.Addrs(config.Get("rabbitmq_addr").String("")),
	)
}
