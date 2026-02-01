package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kiamars-mirzaee/GoMongoquent/orm"
	_ "github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	orm.BaseModel `bson:",inline"`
	Name          string `bson:"name"`
	Email         string `bson:"email"`
	// Relationship Field
	Posts []Post `bson:"posts,omitempty"`
}

func (u *User) CollectionName() string { return "users" }

// Define the Relationship
func (u *User) MyPosts() orm.Relationship {
	return orm.HasMany(&Post{}, "_id", "user_id", "posts")
}

type Post struct {
	orm.BaseModel `bson:",inline"`
	Title         string             `bson:"title"`
	UserID        primitive.ObjectID `bson:"user_id"`
}

func (p *Post) CollectionName() string { return "posts" }

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := orm.Connect(ctx, "mongodb://localhost:27017", "test_db")
	if err != nil {
		panic(err)
	}

	repo := orm.NewRepository(db, &User{})

	// Fetch users named 'Alice' and include their Posts in one query
	users, err := repo.Query(ctx).
		Where("name", "Alice").
		With("MyPosts"). // Triggers the Lookup
		Get()

	for _, user := range users {
		fmt.Printf("User: %s has %d posts\n", user.Name, len(user.Posts))
	}
}
