package users

import (
	"fmt"

	clientdbs "github.com/kadekchrisna/openbook/datasources/mysql/users_db"
	cryptoutils "github.com/kadekchrisna/openbook/utils/crypto"
	dateutils "github.com/kadekchrisna/openbook/utils/date"
	"github.com/kadekchrisna/openbook/utils/errors"
	mysqlutils "github.com/kadekchrisna/openbook/utils/mysql_utils"
)

var (
	userDN = make(map[int64]*User)
)

// Get asd
func (user *User) Get() *errors.ResErr {
	if errPing := clientdbs.Client.Ping(); errPing != nil {
		panic(errPing)
	}

	stmn, errPrepare := clientdbs.Client.Prepare("SELECT u.first_name, u.last_name, u.email, u.status, u.date_created FROM users u WHERE u.id = ?;")
	if errPrepare != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Preparing error. %s", errPrepare.Error()))
	}

	result := stmn.QueryRow(user.Id)
	if errExec := result.Scan(&user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); errExec != nil {
		// return errors.NewNotFoundError(fmt.Sprintf("user not found with id %d", user.Id))
		return mysqlutils.ParseError(errExec)
	}
	// result := userDN[user.Id]
	// if result == nil {
	// 	return errors.NewNotFoundError(fmt.Sprintf("user not found with id %d", user.Id))
	// }
	// user.Email = result.Email
	// user.FirstName = result.FirstName
	// user.LastName = result.LastName
	// user.DateCreated = result.DateCreated
	return nil
}

// Save asd
func (user *User) Save() *errors.ResErr {
	stmn, errPrepare := clientdbs.Client.Prepare("INSERT INTO users (first_name, last_name, email, status, date_created, password) VALUES(?, ?, ?, ?, ?, ?);")
	if errPrepare != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Preparing error. %s", errPrepare.Error()))
	}
	defer stmn.Close()
	user.DateCreated = dateutils.GetDBFormatedNow()
	user.Password = cryptoutils.GetMD5(user.Password)

	insertResult, errInsert := stmn.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.DateCreated, user.Password)
	if errInsert != nil {
		// sqlErr, ok := errInsert.(*mysql.MySQLError)
		// if !ok {
		// 	return errors.NewInternalServerError(fmt.Sprintf("Inserting new user error. %s", errInsert.Error()))
		// }
		// fmt.Println(sqlErr)
		// fmt.Println(sqlErr)

		// if strings.Contains(errInsert.Error(), "users_email_uindex") {
		// 	return errors.NewBadRequestError(fmt.Sprintf("Duplicate email. %s", user.Email))
		// }
		// return errors.NewInternalServerError(fmt.Sprintf("Inserting new user error. %s", errInsert.Error()))
		return mysqlutils.ParseError(errInsert)

	}

	// Line 43 is the same functionality as line 33-41 but without checking the query first like 33-41
	// result, errInsert := clientdbs.Client.Exec("INSET INTO users (first_name, last_name, email, status, date_created VALUES(?, ?, ?, ?, ?));", user.FirstName, user.LastName, user.Email, user.Status, user.DateCreated)

	userId, errLastInsert := insertResult.LastInsertId()
	if errLastInsert != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Getting new user id error. %s", errLastInsert.Error()))
	}

	// result := userDN[user.Id]
	// if result != nil {
	// 	if result.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("email already taken"))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("user exist with id %d", user.Id))
	// }
	// userDN[user.Id] = user

	user.Id = userId
	return nil
}

// Update asd
func (user *User) Update() *errors.ResErr {
	stmn, errPrepare := clientdbs.Client.Prepare("UPDATE users set first_name=?, last_name=?, email=? WHERE id=?;")
	if errPrepare != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Preparing error. %s", errPrepare.Error()))
	}
	defer stmn.Close()

	_, errUpdate := stmn.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if errUpdate != nil {
		return mysqlutils.ParseError(errUpdate)
	}
	return nil
}

// Delete asd
func (user *User) Delete() *errors.ResErr {
	stmn, errPrepare := clientdbs.Client.Prepare("DELETE u FROM users u WHERE u.id=?")
	if errPrepare != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Preparing error. %s", errPrepare.Error()))
	}
	defer stmn.Close()

	_, errDelete := stmn.Exec(user.Id)
	if errDelete != nil {
		return mysqlutils.ParseError(errDelete)
	}
	return nil
}

// FindUserByStatus asd
func (user *User) FindUserByStatus(status string) (Users, *errors.ResErr) {
	stmn, errPrepare := clientdbs.Client.Prepare(`SELECT first_name, last_name, email, status, date_created FROM users WHERE status=?;`)
	if errPrepare != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Preparing error. %s", errPrepare.Error()))
	}
	defer stmn.Close()
	result, errGet := stmn.Query(status)
	if errGet != nil {
		return nil, mysqlutils.ParseError(errGet)
	}
	defer result.Close()
	users := make([]User, 0)
	for result.Next() {
		var user User
		if err := result.Scan(&user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
			return nil, mysqlutils.ParseError(err)
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No record found with status %s.", status))
	}
	return users, nil
}

// FindUserByEmailAndPassword FindUserByEmailAndPassword
func (user *User) FindUserByEmailAndPassword() *errors.ResErr {
	stmn, errPrepare := clientdbs.Client.Prepare("SELECT id, first_name, last_name, email, status, date_created from users WHERE email=? AND password=? AND status=?")
	if errPrepare != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Preparing error. %s", errPrepare.Error()))
	}
	defer stmn.Close()
	user.Password = cryptoutils.GetMD5(user.Password)
	result := stmn.QueryRow(user.Email, user.Password, StatusActive)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}
