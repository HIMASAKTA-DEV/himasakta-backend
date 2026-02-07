package seeders

import (
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/db/seeder/seeds"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
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
