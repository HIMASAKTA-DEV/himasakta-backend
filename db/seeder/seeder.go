package seeders

import (
	"fmt"

	"github.com/azkaazkun/be-samarta/db/seeder/seeds"
	mylog "github.com/azkaazkun/be-samarta/internal/pkg/logger"
	"gorm.io/gorm"
)

func Seeding(db *gorm.DB) error {
	seeders := []func(*gorm.DB) error{
		seeds.SeederUser,
		seeds.SeederItem,
		seeds.SeederProposal,
	}

	fmt.Println(mylog.ColorizeInfo("\n=========== Start Seeding ==========="))
	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			return err
		}
	}

	return nil
}
