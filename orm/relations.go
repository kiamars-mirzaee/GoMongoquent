package orm

import "go.mongodb.org/mongo-driver/bson"

// Relationship defines how two collections connect
type Relationship struct {
	FromCollection string
	LocalField     string // Field in the parent (e.g., _id)
	ForeignField   string // Field in the child (e.g., user_id)
	As             string // Field name in the struct to populate
}

// HasMany relationship helper
func HasMany(childModel Model, localField, foreignField, as string) Relationship {
	return Relationship{
		FromCollection: childModel.CollectionName(),
		LocalField:     localField,
		ForeignField:   foreignField,
		As:             as,
	}
}

// BelongsTo relationship helper
func BelongsTo(parentModel Model, localField, foreignField, as string) Relationship {
	return Relationship{
		FromCollection: parentModel.CollectionName(),
		LocalField:     localField,
		ForeignField:   foreignField,
		As:             as,
	}
}

// LookupStage AggregationStage converts the relationship to a Mongo $lookup
func (r Relationship) LookupStage() bson.D {
	return bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: r.FromCollection},
		{Key: "localField", Value: r.LocalField},
		{Key: "foreignField", Value: r.ForeignField},
		{Key: "as", Value: r.As},
	}}}
}
