package handler

import (
	"context"
	"strconv"

	"github.com/laixhe/go-micro-grpc/protorpc"
)

type UserService struct{}

// 获取某个 user 数据
func (u *UserService) GetUser(c context.Context, req *protorpc.GetUserRequest, res *protorpc.GetUserResponse) error {

	// 待返回数据结构
	res.Userid = req.Userid
	res.Username = "laixhe..."
	res.Sex = protorpc.UserSex_MEN

	return nil
}

// 获取 user 所有数据
func (u *UserService) GetUserList(c context.Context, ret *protorpc.GetUserListRequest, res *protorpc.UserListResponse) error {

	list := make([]*protorpc.GetUserResponse, 0, 3)

	for i := 1; i <= 3; i++ {
		list = append(list, &protorpc.GetUserResponse{Userid: int64(i), Username: "laiki" + strconv.Itoa(i), Sex: protorpc.UserSex_MEN})
	}

	res.List = list

	return nil
}
