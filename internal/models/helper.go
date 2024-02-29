package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	mysqldriver "github.com/go-sql-driver/mysql"
)

func mysqlDSN(user, pwd string, address string, port uint16, database, charset string, parseTime bool, loc string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		user, pwd, address, port, database, charset, strings.ToTitle(strconv.FormatBool(parseTime)), loc)
}

func dsn() string {
	return mysqlDSN(user, password, address, port, database, charset, parseTime, loc)
}

func dsnWithoutDatabase() string {
	return mysqlDSN(user, password, address, port, "", charset, parseTime, loc)
}

func createDatabase() error {
	db, err := sql.Open("mysql", dsnWithoutDatabase())
	if err != nil {
		return err
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", database))
	if err != nil {
		var e *mysqldriver.MySQLError
		if !errors.As(err, &e) {
			return err
		}
		if e.Number != 1007 {
			return err
		}
	}

	return nil
}
