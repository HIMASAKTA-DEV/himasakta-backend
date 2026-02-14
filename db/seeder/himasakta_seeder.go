package seeders

import (
	"fmt"
	"time"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func HimasaktaSeeder(db *gorm.DB) error {
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
	if err := db.Create(&gallery1).Error; err != nil {
		return err
	}
	if err := db.Create(&gallery2).Error; err != nil {
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
	if err := db.Create(&cabinet).Error; err != nil {
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
	if err := db.Create(&dept).Error; err != nil {
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
	if err := db.Create(&member).Error; err != nil {
		return err
	}

	// Time Line
	timeline := entity.Timeline{
		Date: time.Date(1016, 12, 11, 0, 0, 0, 0, time.UTC),
		Info: "Oprec",
		Link: "https://...",
	}
	// 5. Seed Progenda
	progenda := entity.Progenda{
		Id:           uuid.New(),
		Name:         "LDKM",
		ThumbnailId:  &gallery2.Id,
		Goal:         "Melatih jiwa kepemimpinan.",
		Description:  "Latihan Dasar Kepemimpinan Mahasiswa.",
		Timelines:    []entity.Timeline{timeline},
		DepartmentId: &dept.Id,
	}
	if err := db.Create(&progenda).Error; err != nil {
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
	if err := db.Create(&event).Error; err != nil {
		return err
	}

	// 7. Seed News
	news := entity.News{
		Id:          uuid.New(),
		Title:       "Penerimaan Anggota Baru",
		Tagline:     "Ayo bergabung!",
		Hashtags:    "OPREC,HIMASAKTA",
		Content:     "# Selamat Datang\nKami membuka pendaftaran anggota baru.",
		ThumbnailId: &gallery2.Id,
		PublishedAt: time.Now(),
	}
	if err := db.Create(&news).Error; err != nil {
		return err
	}

	fmt.Println("Himasakta Seeder: OK")
	return nil
}
