// ---------- Core Interfaces ----------

type Model interface {
	CollectionName() string
	PrimaryKey() string // usually _id
}

// ---------- Base Model ----------

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time           `bson:"created_at,omitempty"`
	UpdatedAt time.Time           `bson:"updated_at,omitempty"`
}

func (b *BaseModel) PrimaryKey() string {
	return "_id"
}