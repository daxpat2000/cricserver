package domain

import "time"

type Team struct {
	Name      string    `json:"name" gorm:"primaryKey"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func NewTeam(name string, teamType string) *Team {
	return &Team{
		Name: name,
		Type: teamType,
	}
}

func (t *Team) String() string {
	return t.Name
}
