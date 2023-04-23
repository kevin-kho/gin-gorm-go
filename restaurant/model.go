package restaurant

import (
	"gorm.io/gorm"
	"time"
)

type Restaurant struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`

	Likes    uint `json:"likes"`
	Dislikes uint `json:"dislikes"`
}

func (r *Restaurant) AddLike() {
	r.Likes += 1
}

func (r *Restaurant) AddDislike() {
	r.Dislikes += 1
}
