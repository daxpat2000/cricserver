package domain

import (
	"fmt"
	"time"

	"github.com/daxpat2000/cricserver/pkg/types"
)

type Player struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	FirstName    string     `json:"firstName" binding:"required"`
	LastName     string     `json:"lastName" binding:"required"`
	Dob          types.Date `json:"dob" binding:"required"`
	Role         string     `json:"role"`
	BattingStyle string     `json:"battingStyle"`
	BowlingStyle string     `json:"bowlingStyle"`
	Nationality  string     `json:"nationality" binding:"required"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}

func NewPlayer(fname, lname, role, battingStyle, bowlingStyle, nationality string, dob types.Date) *Player {
	return &Player{
		FirstName:    fname,
		LastName:     lname,
		Dob:          dob,
		BattingStyle: battingStyle,
		BowlingStyle: bowlingStyle,
		Nationality:  nationality,
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
