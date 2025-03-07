package models

import (
	"time"
)

type HeadDataDTO struct {
	UserEmail       string     `json:"userEmail"`
	Name            string     `json:"name"`
	ExperienceSince *time.Time `json:"experienceSince"`
	StartAtProdyna  *time.Time `json:"startAtProdyna"`
	Age             int        `json:"age"`
	Abstract        string     `json:"abstract"`
	AgreedOn        bool       `json:"agreedOn"`
}
