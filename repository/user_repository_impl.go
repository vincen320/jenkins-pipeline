package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vincen320/user-service/helper"
	"github.com/vincen320/user-service/model/domain"
)

//cara auto generate
//Ctrl Shift P
//Go Generate Interface stubs
//buat sebuah komen >> //us *UserRepositoryImpl
// ketik <package.Interface> contoh: repository.UserRepository

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (us *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `INSERT INTO user(username, password, created_at, last_online) VALUES (?,?,?,?)`
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.CreatedAt, user.LastOnline)
	helper.PanicIfError(err) //tidak usah didefinisikan errornya apa, nanti ditangkap oleh errorhandler saja, karena ini merupakan bagian dari error 500(Internal eror)

	lastInserId, err := result.LastInsertId()
	helper.PanicIfError(err) // 500 internal error

	user.Id = int(lastInserId)
	return user
}

func (us *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `UPDATE USER SET username=?, password=?, last_online=? WHERE ID=?`
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.LastOnline, user.Id)
	helper.PanicIfError(err) // 500 Internal Server Error

	if rowsaffected, err := result.RowsAffected(); err != nil {
		panic(err)
	} else if rowsaffected == 1 {
		return user
	}
	return domain.User{}
}

func (us *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) bool {
	SQL := `DELETE FROM USER WHERE ID = ?`
	result, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err) // 500 Internal Error

	if rowsaffected, err := result.RowsAffected(); err != nil {
		panic(err)
	} else if rowsaffected == 1 {
		return true
	}
	return false
}

func (us *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, UserId int) (domain.User, error) {
	SQL := `SELECT id, username, password, created_at, last_online FROM user WHERE id=?`
	rows, err := tx.QueryContext(ctx, SQL, UserId)
	helper.PanicIfError(err) //500 internal server error
	defer rows.Close()       //JANGAN LUPA CLOSE

	var user domain.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.LastOnline)
		helper.PanicIfError(err) //500 internal server error
		return user, nil
	} else {
		return user, errors.New("user not found")
	}

}

func (us *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := `SELECT id, username, password, created_at, last_online FROM user`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err) // 500 internal server error
	defer rows.Close()       //JANGAN LUPA CLOSE

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.LastOnline)
		helper.PanicIfError(err) //500 internal serveer error
		users = append(users, user)
	}
	return users
}

func (us *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	SQL := `SELECT id, username, password, created_at, last_online FROM user WHERE username=?`
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err) //500 internal server error
	defer rows.Close()       //JANGAN LUPA CLOSE

	var user domain.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.LastOnline)
		helper.PanicIfError(err) //500 internal server error
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}
