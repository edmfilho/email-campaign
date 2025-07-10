package campaign

type Repository interface {
	Save(campaign *Campaign) error
	Get() []Campaign
	GetByID(id string) (Campaign, error)
}
