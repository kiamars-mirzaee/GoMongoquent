package orm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Model interface {
	CollectionName() string
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	SetCreatedAt(t time.Time)
	SetUpdatedAt(t time.Time)
	GetDeletedAt() *time.Time
}

// BaseModel implements standard fields
type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

func (b *BaseModel) GetID() primitive.ObjectID   { return b.ID }
func (b *BaseModel) SetID(id primitive.ObjectID) { b.ID = id }
func (b *BaseModel) SetCreatedAt(t time.Time)    { b.CreatedAt = t }
func (b *BaseModel) SetUpdatedAt(t time.Time)    { b.UpdatedAt = t }
func (b *BaseModel) GetDeletedAt() *time.Time    { return b.DeletedAt }
