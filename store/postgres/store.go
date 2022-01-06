package postgres

import (
	"fmt"
	"log"

	"github.com/odpf/guardian/store"
	"github.com/odpf/guardian/store/model"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(c *store.Config) (*Store, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=%s",
		c.Host,
		c.User,
		c.Name,
		c.Password,
		c.Port,
		c.SslMode,
	)

	gormDB, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	return &Store{gormDB}, nil
}

func (s *Store) DB() *gorm.DB {
	return s.db
}

func (s *Store) Migrate() error {
	return s.db.AutoMigrate(
		&model.Provider{},
		&model.Policy{},
		&model.Resource{},
		&model.Appeal{},
		&model.Approval{},
		&model.Approver{},
	)
}