package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/vincen320/user-service/exception"
	"github.com/vincen320/user-service/helper"
	"github.com/vincen320/user-service/model/appservice/authservice"
	"github.com/vincen320/user-service/model/domain"
	"github.com/vincen320/user-service/model/web"
	"github.com/vincen320/user-service/repository"
)

//us *UserServiceImpl
//pada baris diatas tekan CTRL +SHIFT+ P
//Pilih Go Generate Interface Stubs
//ketik package.Interface >> service.UserService

type UserServiceImpl struct {
	Repository repository.UserRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewUserService(repository repository.UserRepository, db *sql.DB, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validator,
	}
}

func (us *UserServiceImpl) Create(ctx context.Context, userCreate web.UserCreateRequest) web.UserResponse {
	err := us.Validate.Struct(userCreate)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(err) // Validation Error
	}

	tx, err := us.DB.Begin()
	helper.PanicIfError(err) //500 Internal
	defer helper.CommitofRollBack(tx)

	timeNow := time.Now().UTC().UnixMilli()

	hashPassword := helper.BcryptPassword(userCreate.Password)
	userCreate.Password = string(hashPassword)

	result := us.Repository.Save(ctx, tx, domain.User{
		Username:   userCreate.Username,
		Password:   userCreate.Password,
		CreatedAt:  timeNow,
		LastOnline: timeNow,
	})

	return web.UserResponse{
		Id:        result.Id,
		Username:  result.Username,
		CreatedAt: timeNow,
	}
}

func (us *UserServiceImpl) Update(ctx context.Context, userUpdate web.UserUpdateRequest) web.UserResponse {
	err := us.Validate.Struct(userUpdate)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(err) // Validation Error
	}

	tx, err := us.DB.Begin()
	helper.PanicIfError(err) //500 Internal
	defer helper.CommitofRollBack(tx)

	_, err = us.Repository.FindById(ctx, tx, userUpdate.Id)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewNotFoundErr(err.Error())) //404 Not Found
	}

	hashPassword := helper.BcryptPassword(userUpdate.Password)
	userUpdate.Password = string(hashPassword)

	timeNow := time.Now().UTC().UnixMilli()
	userUpdate.LastOnline = timeNow
	result := us.Repository.Update(ctx, tx, domain.User{
		Id:         userUpdate.Id,
		Username:   userUpdate.Username,
		Password:   userUpdate.Password,
		LastOnline: userUpdate.LastOnline,
	})

	if result.Id == 0 { //Karena di repository.Update ada pengondisian jika rowsaffectednya != 1, maka return domain.User{}
		panic("cannot update more than 2 accounts at the same time - vincen") //500 Internal Server Error
	}

	return web.UserResponse{
		Id:        result.Id,
		Username:  result.Username,
		CreatedAt: result.CreatedAt,
	}
}

func (us *UserServiceImpl) Delete(ctx context.Context, userId int) bool {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err) //500 Internal
	defer helper.CommitofRollBack(tx)

	user, err := us.Repository.FindById(ctx, tx, userId)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewNotFoundErr(err.Error())) //404 Not Found
	}

	isDeleted := us.Repository.Delete(ctx, tx, user)
	if isDeleted {
		return true
	} else {
		panic("cannot delete more than 2 accounts at the same time - vincen") // JADIKAN INTERNAL 500 SAJA, nanti dikirim ke commitorrollback, dan msgErrornya != nil, lalu panic lagi (baris ke 3), lalu baru ke error handler
	}
}

func (us *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err) //500 Internal
	defer helper.CommitofRollBack(tx)

	result, err := us.Repository.FindById(ctx, tx, userId)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewNotFoundErr(err.Error())) //Not Found
	}

	return web.UserResponse{
		Id:         result.Id,
		Username:   result.Username,
		CreatedAt:  result.CreatedAt,
		LastOnline: result.LastOnline,
	}
}

func (us *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err) //500 Internal
	defer helper.CommitofRollBack(tx)

	results := us.Repository.FindAll(ctx, tx)
	var userResponses []web.UserResponse
	for _, v := range results {
		userResponses = append(userResponses, web.UserResponse{
			Id:         v.Id,
			Username:   v.Username,
			CreatedAt:  v.CreatedAt,
			LastOnline: v.LastOnline,
		})
	}
	return userResponses
}

func (us *UserServiceImpl) UpdatePatch(ctx context.Context, userUpdate web.UserUpdatePatchRequest) web.UserResponse {
	err := us.Validate.Struct(userUpdate)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(err) // Validation Error
	}

	tx, err := us.DB.Begin()
	helper.PanicIfError(err) //500 Internal
	defer helper.CommitofRollBack(tx)
	//Find Dulu, lalu GET Valuenya
	userUpdateData, err := us.Repository.FindById(ctx, tx, userUpdate.Id)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewNotFoundErr(err.Error())) //404 Not Found
	}

	//Copy dari userUpdate ke userUpdateData, jadi Data yang telah diupdate ada di variable userUpdateData
	copier.CopyWithOption(&userUpdateData, &userUpdate, copier.Option{
		IgnoreEmpty: true,
	})

	//Seharusnya Password tidak bisa diedti barengan, karena harus di endpoint /ChangePassword lebih baik, ini contoh sja, kondisi khusus
	if userUpdate.Password != "" {
		//sehingga dsini pakai userUpdateData bukan userUpdate lagi (lanjutan dari comment diatas copier.CopyWithOption)
		hashPassword := helper.BcryptPassword(userUpdateData.Password)
		userUpdateData.Password = string(hashPassword)
	}

	timeNow := time.Now().UTC().UnixMilli()
	userUpdateData.LastOnline = timeNow

	result := us.Repository.Update(ctx, tx, userUpdateData)

	if result.Id == 0 { //Karena di repository.Update ada pengondisian jika rowsaffectednya != 1, maka return domain.User{}, jadi dibuat ini kalau yang direturn adalah domain.User{} karena fieldnya pasti berisi default semua, int=0, stirng=""
		panic("cannot update more than 2 accounts at the same time - vincen") //500 Internal Server Error
	}

	return web.UserResponse{
		Id:        result.Id,
		Username:  result.Username,
		CreatedAt: result.CreatedAt,
	}
}

func (us *UserServiceImpl) FindByUsername(ctx context.Context, username string) authservice.UserServiceLoginResponse {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err) //500 Internal
	defer helper.CommitofRollBack(tx)

	result, err := us.Repository.FindByUsername(ctx, tx, username)

	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewNotFoundErr(err.Error())) //Not Found
	}

	return authservice.UserServiceLoginResponse{
		Id:         result.Id,
		Username:   result.Username,
		Password:   result.Password,
		CreatedAt:  result.CreatedAt,
		LastOnline: result.LastOnline,
	}
}
