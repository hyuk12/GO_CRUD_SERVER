package network

import (
	"CRUD_SERVER/service"
	"github.com/gin-gonic/gin"
)

// network 라이브러리로 go get github.com/gin-gonic/gin 씀
type Network struct {
	engin *gin.Engine

	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	r := &Network{
		engin: gin.New(), // gin.New() , gin.Default() 뉴는 운영환경, 디폴트는 테스트환경 근데 굳이 상관은 없다.
	}

	// mux, echo, gin 중에 gin 을 사용
	// 선호하는 이유? 리퀘스트를 밸리데이션 하기 쉬움
	newUserRouter(r, service.User)

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}
