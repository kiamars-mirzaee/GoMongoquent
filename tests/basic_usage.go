func TestUserRepository_CRUD(t *testing.T) {
ctx := context.Background()
db, _ := orm.Connect(ctx, "mongodb://localhost:27017", "test_db")


repo := orm.NewRepository(db, User{})


user := User{Name: "Alice", Email: "alice@test.com"}
err := repo.Create(ctx, &user)
assert.NoError(t, err)


found, err := repo.Query(ctx).Where("email", "alice@test.com").First()
assert.NoError(t, err)
assert.Equal(t, "Alice", found.Name)


err = repo.Delete(ctx, user.ID)
assert.NoError(t, err)
}