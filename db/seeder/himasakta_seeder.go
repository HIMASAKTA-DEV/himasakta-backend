package seeders

import (
	"time"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func HimasaktaSeeder(db *gorm.DB) error {
	// Use transactions for safety
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. Seed Gallery
		gallery1 := entity.Gallery{
			Id:       uuid.New(),
			ImageUrl: "https://via.placeholder.com/150",
			Caption:  "Logo HIMASAKTA",
			Category: "logo",
		}
		gallery2 := entity.Gallery{
			Id:       uuid.New(),
			ImageUrl: "https://via.placeholder.com/800x400",
			Caption:  "Thumbnail News",
			Category: "thumbnail",
		}
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&gallery1).Error; err != nil {
			return err
		}
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&gallery2).Error; err != nil {
			return err
		}

		// 2. Seed Cabinet Info
		cabinet := entity.CabinetInfo{
			Id:          uuid.New(),
			Visi:        "Menjadi himpunan yang progresif.",
			Misi:        "Membangun sinergi antar mahasiswa.",
			Description: "Himpunan Mahasiswa Informatika yang bertujuan untuk mengembangkan kompetensi dan karakter mahasiswa.",
			Tagline:     "Sinergi dalam Aksi",
			PeriodStart: "2024-01-01",
			PeriodEnd:   "2024-12-31",
			LogoId:      &gallery1.Id,
		}
		if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&cabinet).Error; err != nil {
			return err
		}

		// 3. Seed Department
		dept := entity.Department{
			Id:          uuid.New(),
			Name:        "Kaderisasi",
			Description: "Departemen yang berfokus pada pengembangan sumber daya mahasiswa.",
			LogoId:      &gallery1.Id,
			BankRefLink: "https://linktr.ee/himasakta",
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			UpdateAll: true,
		}).Create(&dept).Error; err != nil {
			return err
		}

		// 4. Seed Member
		member := entity.Member{
			Id:           uuid.New(),
			Nrp:          "12345678",
			Name:         "John Doe",
			Role:         "Ketua Departemen",
			DepartmentId: &dept.Id,
			PhotoId:      &gallery1.Id,
			Period:       "2024",
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			UpdateAll: true,
		}).Create(&member).Error; err != nil {
			return err
		}

		// 5. Seed Progenda
		// Cleanup existing progendas first as requested
		tx.Unscoped().Where("1 = 1").Delete(&entity.ProgendaTimeline{})
		tx.Unscoped().Where("1 = 1").Delete(&entity.Progenda{})

		progenda := entity.Progenda{
			Id:            uuid.New(),
			Name:          "LDKM",
			ThumbnailId:   &gallery2.Id,
			Goal:          "Melatih jiwa kepemimpinan.",
			Description:   "Latihan Dasar Kepemimpinan Mahasiswa.",
			InstagramLink: "https://instagram.com/himasakta",
			YoutubeLink:   "https://youtube.com/himasakta",
			DepartmentId:  &dept.Id,
			Timelines: []entity.ProgendaTimeline{
				{EventName: "Pendaftaran", Date: "Maret 2024"},
				{EventName: "Pelaksanaan", Date: "April 2024"},
			},
		}
		if err := tx.Create(&progenda).Error; err != nil {
			return err
		}

		// 6. Seed Monthly Event
		event := entity.MonthlyEvent{
			Id:          uuid.New(),
			Title:       "HIMASAKTA Cup",
			ThumbnailId: &gallery2.Id,
			Description: "Turnamen olahraga antar angkatan.",
			Month:       time.Now(),
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "title"}},
			UpdateAll: true,
		}).Create(&event).Error; err != nil {
			return err
		}

		// 7. Seed News
		news := entity.News{
			Id:          uuid.New(),
			Title:       "Penerimaan Anggota Baru",
			Slug:        "penerimaan-anggota-baru",
			Tagline:     "Ayo bergabung!",
			Hashtags:    "OPREC,HIMASAKTA",
			Content:     "# Selamat Datang\nKami membuka pendaftaran anggota baru.",
			ThumbnailId: &gallery2.Id,
			PublishedAt: time.Now(),
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "title"}},
			UpdateAll: true,
		}).Create(&news).Error; err != nil {
			return err
		}

		return nil
	})
}
