// ---------- Example Model ----------

type User struct {
	BaseModel `bson:",inline"`
	Name      string `bson:"name"`
	Email     string `bson:"email"`
}

func (User) CollectionName() string {
	return "users"
}

// ---------- Usage ----------
//
// db, _ := orm.Connect(ctx, "mongodb://localhost:27017", "app")
// users := orm.NewRepository(db, User{})
//
// users.Query(ctx).
//   Where("email", "test@mail.com").
//   First()


users.Query(ctx).
    Where("email", "test@mail.com").
    Sort("created_at", false).
    Limit(10).
    Get()
