package service_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vincen320/user-service/mock"
	usermock "github.com/vincen320/user-service/mock"
	"github.com/vincen320/user-service/model/domain"
	"github.com/vincen320/user-service/model/web"
	"github.com/vincen320/user-service/service"
	"golang.org/x/crypto/bcrypt"
)

type testCreateUserCase struct {
	name          string
	createRequest web.UserCreateRequest
	stubs         func(*mock.MockUserRepository)
	checkPanic    func()
	checkResponse func(web.UserResponse)
}

func TestCreate(t *testing.T) {

	//===DATA
	domainUser := newDomainUser()
	domainUserReturn := domainUser
	domainUserReturn.Id = 1

	var testCreateUserCases = []testCreateUserCase{
		{
			name: "OK",
			createRequest: web.UserCreateRequest{
				Username: domainUser.Username,
				Password: domainUser.Password,
			},
			stubs: func(userRepositoryMock *usermock.MockUserRepository) {
				userRepositoryMock.EXPECT().
					Save(
						gomock.Any(),              // ctx context.Context,
						gomock.Any(),              // tx *sql.Tx,
						eqUserCreate(domainUser)). // domain.User
					Times(1).
					Return(domainUserReturn) //user domain.User
			},
			checkPanic: func() {
				err := recover()
				assert.Nil(t, err)
			},
			checkResponse: func(response web.UserResponse) {
				assert.Equal(t, domainUserReturn.Id, response.Id)
				assert.Equal(t, domainUserReturn.Username, response.Username)
				//assert.Equal(t, domainUserReturn.LastOnline, response.LastOnline) waktu create ternyata tidak diberi response LastOnlinenya
				assert.LessOrEqual(t, domainUserReturn.CreatedAt, response.CreatedAt)
			},
		}, {
			name: "Validation Error", //PANIC
			createRequest: web.UserCreateRequest{
				Username: "",
				Password: domainUser.Password,
			},
			stubs: func(userRepositoryMock *usermock.MockUserRepository) {
			},
			checkPanic: func() {
				err := recover()
				assert.NotNil(t, err)
			},
			checkResponse: func(response web.UserResponse) {
				assert.Nil(t, response)
			},
		},
	}

	//=====RUNNING TEST HERE
	db, sqlmock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	validator := validator.New()

	mockController := gomock.NewController(t)
	userRepositoryMock := usermock.NewMockUserRepository(mockController)

	for _, tc := range testCreateUserCases {
		t.Run(tc.name, func(t *testing.T) {
			sqlmock.ExpectBegin() //KARENA SELALU BEGIN DAN COMMIT SAJA, UNTUK QUERYNYA DI MOCK PAKAI MOCKGEN JADI, SQL.DB NYA CUMA DIPAKAI UNTUK BEGIN AND COMMIT

			tc.stubs(userRepositoryMock)

			sqlmock.ExpectCommit() //KARENA SELALU BEGIN DAN COMMIT SAJA, UNTUK QUERYNYA DI MOCK PAKAI MOCKGEN JADI, SQL.DB NYA CUMA DIPAKAI UNTUK BEGIN AND COMMIT/LEBI LANJUT MUNGKIN BAKAL ROLLBACK SUATU SAAT JDI DITARUH DI FUNC TERTENTU

			defer tc.checkPanic()
			//userService di init setelah melakukan expect/stubs userRepositoryMock
			userService := service.NewUserService(userRepositoryMock, db, validator)
			response := userService.Create(context.Background(), tc.createRequest)

			tc.checkResponse(response)

			//di define seharusnya setelah fungsi yang akan di test. Dalam hal ini adalah service.Create
			if err := sqlmock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}

}

func newDomainUser() domain.User {
	timeNow := time.Now().UTC().UnixMilli()
	return domain.User{
		Username:   "testUsername",
		Password:   "testPassword",
		CreatedAt:  timeNow,
		LastOnline: timeNow,
	}
}

//1. Buat fungsi
func eqUserCreate(user domain.User) gomock.Matcher {
	//3. Return struct yang telah dibuat
	return eqUserCreateMatcher{
		x: user,
	}
}

//2. Buat Struct yang akan mengimplementasi gomock.Matcher
type eqUserCreateMatcher struct {
	x domain.User
}

func (e eqUserCreateMatcher) Matches(x interface{}) bool {
	// In case, some value is nil
	if x == nil {
		return reflect.DeepEqual(e.x, x)
	}

	// Check if types assignable and convert them to common type
	val, ok := x.(domain.User)
	if !ok {
		return false
	}

	//karena password yang di mock belum dalam bentuk hash sedangkan saat di service dipanggil dan masuk ke repository password telah di hash
	//jadi yang telah di hash ditaruh diparameter 1, yang belum di parameter ke 2
	err := bcrypt.CompareHashAndPassword([]byte(val.Password), []byte(e.x.Password))
	if err != nil {
		return false
	}

	e.x.Password = val.Password //yang biasa diubah ke hash atau bisa juga seperti dibawah
	//val.Password = e.x.password //passowrd yang di hash diubah ke yang bentuk biasa, kalau dari definisi yang String() dibawah, sebaiknya yang hash didiubah seperti biasa(yaitu baris ini)

	//PENGECEKKAN BEDA WAKTU SETIDAKNYA KALAU BEDA 200 miliDETIK TIDAK MASALAH KARENA DI SERVICE JUGA GENERATE BAGIAN CREATED_AT, SEHINGGA AKAN BERBEDA JIKA USER.DOMAIN->CREATED_AT JUGA DIDEFINE DI TEST DAN MENGANGGAP AKAN SANGAT BERBAHAYA JIKA PROGRAM EKSPEK DENGAN CREATED_AT YANG SAMA DENGAN WAKTU YANG DIDEFINE DI TEST
	//DISINI CEKNYA 2 JENIS YAITU CREATED_AT DAN LAST_ONLINE
	//INGAT YANG FUNGSI INI PAKAI UNIXMILI, YG SEBELAH PAKAI UNIX SAJA, JADI INI PERBEDAANNYA LEBI BESAR
	if val.CreatedAt-e.x.CreatedAt <= 200 || val.LastOnline-e.x.LastOnline <= 200 {
		e.x.CreatedAt = val.CreatedAt
		e.x.LastOnline = val.LastOnline //JANGAN LUPA SET 2 2NYA
	}
	return reflect.DeepEqual(e.x, val)
}

func (e eqUserCreateMatcher) String() string {
	return fmt.Sprintf("is equal to %v (%T)", e.x, e.x)
}
