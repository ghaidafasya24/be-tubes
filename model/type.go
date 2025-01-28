package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Role        string             `json:"role,omitempty" bson:"role,omitempty"`
	Username    string             `json:"username,omitempty" bson:"username,omitempty" gorm:"unique;not null"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	PhoneNumber string             `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
}

type Menu struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	MenuName       string             `bson:"menu_name,omitempty" json:"menu_name,omitempty"`
	Price          float64            `bson:"price,omitempty" json:"price,omitempty"`
	Description    string             `bson:"description,omitempty" json:"description,omitempty"`
	Stock          int                `bson:"stock,omitempty" json:"stock,omitempty"`
	MenuCategories Category           `bson:"menu_categories,omitempty" json:"menu_categories,omitempty"`
	// BahanBaku    BahanBaku          `bson:"bahan_baku,omitempty" json:"bahan_baku,omitempty"`
}

type Category struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	MenuCategories string             `bson:"menu_categories,omitempty" json:"menu_categories,omitempty"`
}

// type BahanBaku struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
// 	BahanBaku string             `bson:"bahan_baku,omitempty" json:"bahan_baku,omitempty"`
// 	Jumlah    string             `bson:"jumlah,omitempty" json:"jumlah,omitempty"`
// }
