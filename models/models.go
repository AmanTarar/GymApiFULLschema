package models

import (
	"time"
	"gorm.io/gorm"
)


type Price struct {
	Name  string
	Price float64
}

// users struct
type Users struct {
	User_id        string `json:"user_id" gorm:"default:uuid_generate_v4();uniqueIndex"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Gender         string `json:"gender" `
	Contact_Number string `json:"contact_number"`
	DeleteAT      gorm.DeletedAt
}

// subscriptions struct
type Subscription struct {
	Payment_id        string   `db:"payment_id,foreignkey:Payment(Payment_id)"`
	// Payment           Payment     `gorm:"references:Payment_id"` //FK
	User_id           string      `json:"user_id"`
	User              Users       `gorm:"references:User_id"` //FK
	Start_date        string      `json:"start_date"`
	Membership_type   string      `json:"membership_type"`
	End_date          string      `json:"end_date"`
	Duration          float64      `json:"duration"`
	Employee_id       string      `json:"employee_id"`
	// Employee          GymEmployee `gorm:"references:Employee_id"` //FK
	Trainer_name      string      `json:"trainer_name"`
}

// Payment struct
type Payment struct {
	Payment_id string    `json:"payment_id" gorm:"default:uuid_generate_v4();uniqueIndex"`
	User_id    string    `json:"user_id"`
	User       Users     `gorm:"references:User_id"` //FK
	Amount     float64   `json:"amount"`
	Date       time.Time `json:"date"`
}

// Equipment struct
type Equipment struct {
	Model_number string `json:"model_number" gorm:"default:uuid_generate_v4();unique"`
	Equip_name   string `json:"equip_name"`
	Quantity     int64  `json:"quantity"`
}

type GymEmployee struct {
	Employee_id    string `json:"employee_id" gorm:"default:uuid_generate_v4();unique"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Gender         string `json:"gender"`
	Contact_Number string `json:"contact_number"`
	Role           string `json:"role"`
}