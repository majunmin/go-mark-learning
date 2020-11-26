/**
  @Author: majm@ushareit.com
  @date: 2020/11/26
  @note:
**/
package beegomock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	orm.Debug = true
}

func testWithDb(t *testing.T, f func(t *testing.T, mock sqlmock.Sqlmock)) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("failed to open stub database connection, error: %v", err)
	}

	orm.AddAliasWthDB("default", "mysql", db)
	orm.SetDataBaseTZ("default", time.Now().Location())
	defer db.Close()

	f(t, mock)
}

func TestDBWithMockedSqlDriver(t *testing.T) {
	testWithDb(t, func(t *testing.T, mock sqlmock.Sqlmock) {
		// setup mock
		columns := []string{"id"}
		mock.ExpectQuery("SELECT (.+) FROM `XXX`").
			WillReturnRows(
				sqlmock.NewRows(columns).
					FromCSVString("1").
					FromCSVString("2"))

		// call function to test
		//db.WhateverToTest()

		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
