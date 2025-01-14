package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/kyzyc/biz-demo/app/user/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/app/user/biz/model"
	user "github.com/kyzyc/biz-demo/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	//bcrypt.CompareHashAndPassword([]byte("123456"), []byte("123456"))
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email or password is empty")
	}

	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}

	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}

	fmt.Println("\nid:" + strconv.Itoa(int(newUser.ID)) + "\n")

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
