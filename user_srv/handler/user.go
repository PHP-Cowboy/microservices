package handler

import (
	"context"
	"gorm.io/gorm"
	"user_srv/model"

	"user_srv/protos/proto"
	"user_srv/global"

	"google.golang.org/protobuf/types/known/emptypb"

)

type UserServer struct {

}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {

		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func ModelToResponse(user model.User) *proto.UserInfoResponse {
	response := &proto.UserInfoResponse{
		Id:       uint32(user.Id),
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Gender:   uint32(user.Gender),
		Role: int32(user.Role),
	}
	if user.Birthday != nil {
		response.BirthDay = uint64(user.Birthday.Unix())
	}
	return response
}


func (u *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error){
	users := []model.User{}
	var count int64
	global.DB.Find(&users).Count(&count);

	result := global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	var data = make([]*proto.UserInfoResponse,0)

	for _, user := range users {
		response := ModelToResponse(user)
		data = append(data,response)
	}

	var rsp proto.UserListResponse
	rsp.Total = uint32(count)
	rsp.Data = data

	return &rsp, nil
}

func (u *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error){
	return nil, nil
}

func (u *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error){
	return nil, nil
}

func (u *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error){
	return nil, nil
}

func (u *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*emptypb.Empty, error){
	return nil, nil
}

func (u *UserServer) CheckPassWord(ctx context.Context, req *proto.PassWordCheckInfo) (*proto.CheckResponse, error){
	return nil, nil
}
