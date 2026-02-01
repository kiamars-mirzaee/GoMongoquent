package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kiamars-mirzaee/GoMongoquent/orm"
	"github.com/stretchr/testify/assert"
)

type User struct {
	orm.BaseModel `bson:",inline"` // <- must prefix with package
	Name          string           `bson:"name"`
	Email         string           `bson:"email"`
}

func (u *User) CollectionName() string {
	return "users"
}

func (u *User) PrimaryKey() string {
	return "_id"
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := orm.Connect(ctx, "mongodb://localhost:27017", "test_db")
	if err != nil {
		panic(err)
	}

	repo := orm.NewRepository(db, &User{})

	user := &User{Name: "Alice", Email: "alice@test.com"}
	err = repo.Create(ctx, user)
	assert.NoError(nil, err) // in tests, replace nil with *testing.T

	found, err := repo.Query(ctx).Where("email", "alice@test.com").First()
	assert.NoError(nil, err)
	fmt.Println(found.Name) // Output: Alice
}
