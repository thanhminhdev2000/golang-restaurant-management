package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `json:"invoice_id" validate:"required,min=2,max=100"`
	Order_id         string             `json:"order_id" validate:"required"`
	Payment_method   *string            `json:"payment_method" validate:"required"`
	Payment_status   *string            `json:"payment_status" validate:"required"`
	Payment_due_date time.Time          `json:"payment_due_date" validate:"required"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
}
