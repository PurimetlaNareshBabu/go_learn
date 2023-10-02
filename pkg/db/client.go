package db

import (
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	MYSQL "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DSN          string `mapstructure:"dsn"`
	MaxConns     int    `mapstructure:"max_conns" validate:"required"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" validate:"required"`
}

func LoadDbconfig() (*Config, error) {

	dsn := "root:unacademy@123@tcp(127.0.0.1:3306)/go_learn_1"
	maxConns := 10
	maxIdleConns := 5

	// Validate the required fields
	if dsn == "" {
		return nil, errors.New("missing DSN in the configuration")
	}

	if maxConns <= 0 {
		return nil, errors.New("max_conns must be greater than 0")
	}

	if maxIdleConns <= 0 {
		return nil, errors.New("max_idle_conns must be greater than 0")
	}

	config := &Config{
		DSN:          dsn,
		MaxConns:     maxConns,
		MaxIdleConns: maxIdleConns,
	}

	return config, nil
}

type Client struct {
	DB *gorm.DB
}

func InitDB(c *Config) (*Client, error) {
	db, dberr := gorm.Open(
		mysql.New(mysql.Config{
			DSN: c.DSN,
		}),
		&gorm.Config{})
	if dberr != nil {
		fmt.Printf("Failed to intilise db %s", dberr.Error())
	}
	dbConfig, _ := db.DB()
	dbConfig.SetMaxOpenConns(c.MaxConns)
	dbConfig.SetMaxIdleConns(c.MaxIdleConns)
	dbConfig.SetConnMaxLifetime(12 * time.Hour)
	dbConfig.SetConnMaxIdleTime(-1)
	return &Client{db}, nil
}

func (d *Client) Close() error {
	db, err := d.DB.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) MigrateDB(file_path string) error {
	absPath, err := filepath.Abs("../pkg/db/user/migrations") // Update the path based on your project structure
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", "root:unacademy@123@tcp(127.0.0.1:3306)/go_learn_1?multiStatements=true")
	if err != nil {
		return err
	}
	driver, err := MYSQL.WithInstance(db, &MYSQL.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+absPath,
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func Test() (*Client, error) {
	Con, err := LoadDbconfig()
	if err != nil {
		return nil, err
	}
	db, err := InitDB(Con)
	if err != nil {
		return nil, err
	}
	err = db.MigrateDB("file://migartions")
	if err != nil {
		return db, err
	}
	return db, nil
}
