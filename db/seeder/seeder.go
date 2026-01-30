package seeders

import (
	"fmt"

	"github.com/Flexoo-Academy/Golang-Template/db/seeder/seeds"
	mylog "github.com/Flexoo-Academy/Golang-Template/internal/pkg/logger"
	"gorm.io/gorm"
)

func Seeding(db *gorm.DB) error {
	seeders := []func(*gorm.DB) error{
		seeds.SeederUser,
	}

	fmt.Println(mylog.ColorizeInfo("\n=========== Start Seeding ==========="))
	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			return err
		}
	}

	return nil
}
