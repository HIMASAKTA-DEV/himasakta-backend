package entity

type GlobalSetting struct {
	Key   string `gorm:"type:varchar(50);primaryKey" json:"key"`
	Value string `gorm:"type:jsonb;not null" json:"value"`
}

func (GlobalSetting) TableName() string {
	return "global_settings"
}
