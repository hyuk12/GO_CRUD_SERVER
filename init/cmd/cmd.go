package cmd

import (
	"CRUD_SERVER/config"
	"CRUD_SERVER/network"
	"CRUD_SERVER/repository"
	"CRUD_SERVER/service"
)

type Cmd struct {
	config *config.Config

	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}
	// 위쪽이 하나의 채널로 구성된 것이라 해당 채널이 종료되기 전에는 아래 실행이 되지않는다.

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)

	c.network = network.NewNetwork(c.service)

	err := c.network.ServerStart(c.config.Server.Port)
	if err != nil {
		return nil
	}
	return c
}
