package campaign

import (
	internalerrors "campaign-project/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending string = "Pending"
	Started string = "Started"
	Done    string = "Done"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:50"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Status    string    `validate:"required" gorm:"size:20"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for idx, email := range emails {
		contacts[idx].Email = email
		contacts[idx].ID = xid.New().String()
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Status:    Pending,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}

	err := internalerrors.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err
}
