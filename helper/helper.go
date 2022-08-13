package helper

import (
	"database/sql"

	"github.com/vincen320/user-service/exception"
	"golang.org/x/crypto/bcrypt"
)

func PanicIfError(err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			panic(exception.NewNotFoundErr("product not found"))
		}
		panic(err)
	}
}

func CommitofRollBack(tx *sql.Tx) {
	msgError := recover()
	if msgError != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback) //500 Internal
		panic(msgError)           // PANIC YANG MENYEBABKAN ERRORNYA
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit) //500 Internal
	}
}

//perlu diingat
//hasil dari hashed maxnya adalh 72 karakter
//jadi set length di database paling tidak >=72
func BcryptPassword(password string) []byte {
	bytesPassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytesPassword, bcrypt.DefaultCost)
	PanicIfError(err) //500 Internal
	return hashedPassword
}

//mungkin nanti lebih baik return string saja langsung disini, jadi lebih enak
