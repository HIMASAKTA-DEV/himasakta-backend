package seeders

import (
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"gorm.io/gorm"
)

func HimasaktaSeeder(db *gorm.DB) error {
	webSettings := entity.GlobalSetting{
		Key: "web_settings",
		Value: `{
			"ExternalSOPLink": "https://its.id/m/PostEksternalHimasakta",
			"InternalSOPLink": "https://its.id/m/PostInternalHimasakta",
			"DeskripsiHimpunan": "In the 2024 leadership period, HIMASAKTA ITS adopted the name AVANTURIER as the name of the cabinet. AVANTURIER is derived from Dutch, meaning \"adventurer.\" As the 6th cabinet, Avanturier is expected to carry forward and continue the leadership legacy of HIMASAKTA. It is also hoped that HIMASAKTA ITS will continue to serve the needs of ITS Actuarial students.",
			"FotoHimpunan": "/images/ProfilHimpunan.png",
			"SocialMedia": [
				{"name": "instagram", "link": "https://www.instagram.com/himasakta.its"},
				{"name": "tiktok", "link": "https://www.tiktok.com/@himasakta.its"},
				{"name": "youtube", "link": "https://www.youtube.com/@himasaktaits4262"},
				{"name": "linkedin", "link": "https://www.linkedin.com/company/himasaktaits/posts/?feedView=all"},
				{"name": "linktree", "link": "https://himasaktaits.carrd.co"}
			],
			"InMaintenance": false
		}`,
	}
	if err := db.Create(&webSettings).Error; err != nil {
		if err := db.Save(&webSettings).Error; err != nil {
			return err
		}
	}

	fmt.Println("Himasakta Seeder: OK")
	return nil
}
