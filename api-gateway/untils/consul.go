package untils

import (
	capi "github.com/hashicorp/consul/api"
)

type ConsulClient struct { // 定义一个结构体，用于存储 Consul 客户端
	Address *capi.Client
}

type Consul struct { // 定义一个结构体，用于存储服务信息
	ID      string
	Name    string
	Tags    []string
	Address string
	Port    int
}

func NewConsulClient(address string) (*ConsulClient, error) {
	config := capi.DefaultConfig()        // 创建一个默认的 Consul 客户端配置
	config.Address = address              // 设置 Consul 服务器的地址
	client, err := capi.NewClient(config) // 创建一个新的 Consul 客户端
	if err != nil {
		return nil, err
	}
	return &ConsulClient{Address: client}, err
}

func (c *ConsulClient) RegisterConsul(consul Consul) error {
	registration := capi.AgentServiceRegistration{
		ID:      consul.ID,
		Name:    consul.Name,
		Tags:    consul.Tags,
		Port:    consul.Port,
		Address: consul.Address,
	}
	err := c.Address.Agent().ServiceRegister(&registration)
	if err != nil {
		return err
	}
	return nil
}
