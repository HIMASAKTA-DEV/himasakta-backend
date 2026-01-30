package seeds

import (
	"encoding/json"
	"os"

	"github.com/azkaazkun/be-samarta/internal/entity"
	mylog "github.com/azkaazkun/be-samarta/internal/pkg/logger"
	"gorm.io/gorm"
)

func SeederProposal(db *gorm.DB) error {
	mylog.Infof("[PROCESS] Seeding proposals...")

	jsonFile, err := os.Open("./db/seeder/data/proposal_data.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	var listEntity []entity.Proposal
	if err := json.NewDecoder(jsonFile).Decode(&listEntity); err != nil {
		return err
	}

	for _, proposal := range listEntity {
		if err := db.Create(&proposal).Error; err != nil {
			mylog.Errorf("[ERROR] Failed to seed proposal %s: %v", proposal.No, err)
			return err
		}
	}

	mylog.Infof("[COMPLETE] Seeding proposals completed")
	return nil
}
