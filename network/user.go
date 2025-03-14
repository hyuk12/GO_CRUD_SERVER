package network

import (
	"CRUD_SERVER/service"
	"CRUD_SERVER/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	userRouterInit     sync.Once // 반드시 한번만 호출 되어야함 -> 여러 고루틴에서 동시에 접근하더라도 초기화 함수가 단 한 번만 실행되도록 보장해준다. 그래서 싱글톤의 lazy initialization(필요할 때 초기화)에 유용하게 사용됨.
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service
	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

	})

	router.registerGET("/", userRouterInstance.read)
	router.registerPOST("/", userRouterInstance.create)
	router.registerUPDATE("/", userRouterInstance.update)
	router.registerDELETE("/", userRouterInstance.delete)

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create")
	c.JSON(200, "success")
	err := u.userService.Create(nil)

	if err != nil {
		return
	}

	u.router.okResponse(c, &types.CreateUserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "성공입니다.",
		},
	})
}

func (u *userRouter) read(c *gin.Context) {
	fmt.Println("read")

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "성공입니다.",
		},
		Users: u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update")
	err := u.userService.Update(nil, nil)

	if err != nil {
		return
	}

	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "성공입니다.",
		},
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete")
	err := u.userService.Delete(nil)

	if err != nil {
		return
	}

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "성공입니다.",
		},
	})
}
