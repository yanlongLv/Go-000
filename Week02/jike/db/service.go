package db

import (
	"database/sql"

	"github.com/Go-000/Week02/jike/base"
	"github.com/pkg/errors"
)

// ServiceInterface interface
type ServiceInterface interface {
	GetUsers() (users []*base.User, err error)
	GetUserNameByID(userID int) (name string, err error)
}

// ServerClient init server
type ServerClient struct {
	dbInstall *Connect
}

// NewServerClient server client
func NewServerClient(dbInstall *Connect) ServiceInterface {
	return &ServerClient{dbInstall}
}

// GetUsers service
func (c *ServerClient) GetUsers() (users []*base.User, err error) {
	rows, err := c.dbInstall.db.Query("select * from userBase")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return users, errors.Wrap(err, "query users sql error")
	}
	defer rows.Close()
	for rows.Next() {
		u := new(base.User)
		err = rows.Scan(&u.Name, &u.Age, &u.Phone, &u.Address, &u.ID)
		if err != nil {
			return users, errors.Wrap(err, "rows scan error")
		}
		users = append(users, u)
	}
	return users, err
}

// GetUserNameByID service get user info by userID
func (c *ServerClient) GetUserNameByID(userID int) (name string, err error) {
	stmt, err := c.dbInstall.db.Prepare("select name from userBase where id = ?")
	if err != nil {
		return name, errors.Wrap(err, "get user by userId prepare error")
	}
	defer stmt.Close()
	err = stmt.QueryRow(userID).Scan(&name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return name, errors.Wrap(err, "get user by userId stmt err")
	}
	return name, nil
}
