package orm

import (
	"context"
	"errors"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)




// ---------- Mongo Client Wrapper ----------

type DB struct {
	client *mongo.Client
	db     *mongo.Database
}

func Connect(ctx context.Context, uri, dbName string) (*DB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &DB{
		client: client,
		db:     client.Database(dbName),
	}, nil
}

func (db *DB) Collection(model Model) *mongo.Collection {
	return db.db.Collection(model.CollectionName())
}

