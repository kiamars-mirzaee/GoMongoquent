package orm

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Query[T Model] struct {
	ctx       context.Context
	db        *DB
	model     T
	filter    bson.M
	relations []Relationship
	limit     int64
	skip      int64
	sort      bson.D
}

func NewQuery[T Model](ctx context.Context, db *DB, model T) *Query[T] {
	return &Query[T]{
		ctx:    ctx,
		db:     db,
		model:  model,
		filter: bson.M{},
		sort:   bson.D{},
	}
}

func (q *Query[T]) Where(field string, value any) *Query[T] {
	q.filter[field] = value
	return q
}

// With adds a relationship to be eagerly loaded
// The model T must have a method with the name 'relationName' that returns orm.Relationship
func (q *Query[T]) With(relationName string) *Query[T] {
	// 1. Use Reflection to find the method on the Model
	method := reflect.ValueOf(q.model).MethodByName(relationName)
	if !method.IsValid() {
		// In a real app, you might log a warning here
		return q
	}

	// 2. Call the method to get the Relationship struct
	results := method.Call(nil)
	if len(results) > 0 {
		if rel, ok := results[0].Interface().(Relationship); ok {
			q.relations = append(q.relations, rel)
		}
	}
	return q
}

func (q *Query[T]) Get() ([]T, error) {
	col := q.db.Collection(q.model)

	// Build Aggregation Pipeline
	pipeline := mongo.Pipeline{}

	// 1. Match (Filtering)
	if len(q.filter) > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: q.filter}})
	}

	// 2. Lookups (Eager Loading)
	for _, rel := range q.relations {
		pipeline = append(pipeline, rel.LookupStage())
	}

	// 3. Sort/Limit/Skip
	if len(q.sort) > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$sort", Value: q.sort}})
	}
	if q.skip > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$skip", Value: q.skip}})
	}
	if q.limit > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$limit", Value: q.limit}})
	}

	// Execute Aggregate
	cursor, err := col.Aggregate(q.ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(q.ctx)

	var results []T
	if err = cursor.All(q.ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
