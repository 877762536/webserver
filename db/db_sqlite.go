package db

import (
	"errors"
	"fmt"
	"github.com/localhostjason/webserver/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path/filepath"
)

// ConnectWithSqliteConfig 连接，检验配置是否正确
func ConnectWithSqliteConfig(c SqliteDBConfig) error {
	var dbFile string
	if filepath.IsAbs(c.DbFile) {
		dbFile = c.DbFile
	} else {
		exePath, _ := GetExeDir()
		dbFile = filepath.Join(exePath, c.DbFile)
	}

	path := filepath.Dir(dbFile)
	if !util.PathExists(path) {
		_ = os.MkdirAll(path, os.ModePerm)
	}

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		FullSaveAssociations:   true,
		SkipDefaultTransaction: true,
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return err
	}

	if c.Debug {
		db = db.Debug()
	}

	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect databse %s:%v", c.DbFile, err))
	}
	DB = db
	return nil
}
