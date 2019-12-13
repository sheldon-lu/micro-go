package handler

import (
	"context"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
	"golang.org/x/crypto/bcrypt"
	"micro-go/user/model"
	users "micro-go/user/model/user"
	user "micro-go/user/proto/user"
)

type User struct{
	Repo *users.User
}

func (h *User) Register(ctx context.Context, req *user.RegisterRequest, rsp *user.Response) error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     req.User.Name,
		Phone:    req.User.Phone,
		Password: string(hashedPwd),
	}

	userinfo, err := h.Repo.FindToModel(user)

	if userinfo.Phone == req.User.Phone {
		return errors.BadRequest("go.micro.srv.user.register", "手机号已经存在！")
	}

	if err := h.Repo.Create(user); err != nil {
		log.Log("create error")
		return err
	}

	rsp.Code = "200"
	rsp.Msg = "注册成功"

	return nil
}

func (h *User) Login(ctx context.Context, req *user.LoginRequest, rsp *user.Response) error {
	user, err := h.Repo.FindByField("phone", req.Phone, "id , password")
	if err != err {
		return err
	}

	if user == nil {
		return errors.Unauthorized("go.micro.srv.user.login", "该手机号不存在")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.Unauthorized("go.micro.srv.user.login", "密码错误")
	}

	rsp.Code = "200"
	rsp.Msg = "登录成功"

	return nil
}

func (h *User) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest, rsp *user.Response) error {
	user, err := h.Repo.Find(req.Uid)

	if user == nil {
		return errors.Unauthorized("go.micro.srv.user.login", "该用户不存在")
	}

	if err != nil {
		return err
	}
	//验证老密码是否正常
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return errors.Unauthorized("go.micro.srv.user.login", "旧密码认证失败")
	}

	//验证通过后，对新密码hash存下来
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPwd)
	h.Repo.Update(user, 0)

	rsp.Code = "200"
	rsp.Msg = user.Name + "，您的密码更新成功"

	return nil
}

func (h *User) List(ctx context.Context, rep *user.ListRequest, rsp *user.Response) error {
	return nil
}
