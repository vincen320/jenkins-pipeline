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
//buat sebuah komen >> //us *UserRepositoryImplPq
// ketik <package.Interface> contoh: repository.UserRepository

type UserRepositoryImplPq struct {
}

func NewUserRepositoryPq() UserRepository {
	return &UserRepositoryImplPq{}
}

func (us *UserRepositoryImplPq) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `INSERT INTO user_data(username, password, created_at, last_online) VALUES ($1,$2,$3,$4)`
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.CreatedAt, user.LastOnline)
	helper.PanicIfError(err) //tidak usah didefinisikan errornya apa, nanti ditangkap oleh errorhandler saja, karena ini merupakan bagian dari error 500(Internal eror)

	// lastInserId, err := result.LastInsertId()
	// helper.PanicIfError(err) // 500 internal error

	// user.Id = int(lastInserId) //POSTGRES TIDAK SUPPORT LASTINSERTID
	return user
}

func (us *UserRepositoryImplPq) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `UPDATE user_data SET username=$1, password=$2, last_online=$3 WHERE ID=$4`
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.LastOnline, user.Id)
	helper.PanicIfError(err) // 500 Internal Server Error

	if rowsaffected, err := result.RowsAffected(); err != nil {
		panic(err)
	} else if rowsaffected == 1 {
		return user
	}
	return domain.User{}
}

func (us *UserRepositoryImplPq) Delete(ctx context.Context, tx *sql.Tx, user domain.User) bool {
	SQL := `DELETE FROM user_data WHERE ID = $1`
	result, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err) // 500 Internal Error

	if rowsaffected, err := result.RowsAffected(); err != nil {
		panic(err)
	} else if rowsaffected == 1 {
		return true
	}
	return false
}

func (us *UserRepositoryImplPq) FindById(ctx context.Context, tx *sql.Tx, UserId int) (domain.User, error) {
	SQL := `SELECT id, username, password, created_at, last_online FROM user_data WHERE id=$1`
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

func (us *UserRepositoryImplPq) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := `SELECT id, username, password, created_at, last_online FROM user_data`
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

func (us *UserRepositoryImplPq) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	SQL := `SELECT id, username, password, created_at, last_online FROM user_data WHERE username=$1`
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
