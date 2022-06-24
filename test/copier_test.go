package test

import (
	"fmt"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/vincen320/user-service/model/domain"
	"github.com/vincen320/user-service/model/web"
)

func TestCopier(t *testing.T) {

	UserData := domain.User{
		Id:         1,
		Username:   "ASDF",
		Password:   "Paswre",
		CreatedAt:  123123123,
		LastOnline: 54353534,
	}

	Update := web.UserUpdateRequest{
		Username: "KerenAbis",
		Password: "Ganti-Password",
	}
	//copier.Copy(&UserData, &Update)
	copier.CopyWithOption(&UserData, &Update, copier.Option{
		IgnoreEmpty: true,
	})
	fmt.Println(UserData)
}
