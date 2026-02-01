# GoMongoquent 🐹🍃

> **Laravel Eloquent–inspired MongoDB ORM for Go (without the magic)**

GoMongoquent is a lightweight, strongly-typed MongoDB abstraction for Go developers who love the **developer experience of Laravel Eloquent**, but still want **idiomatic, explicit, and performant Go**.

This project is designed to be:
* 💼 **Portfolio‑ready** (clean architecture, senior-level patterns)
* 🧠 **Easy to reason about** (no hidden magic)
* ⚡ **Production friendly** (context-aware, testable, predictable)
# Go-Eloquent ORM (MongoDB)

A lightweight, type-safe, and developer-friendly Object-Relational Mapper (ORM) for MongoDB in Go, inspired by Laravel Eloquent.

## 🚀 Features

- **Generics-Based**: Fully type-safe queries. No more `interface{}` casting.
- **Eloquent Syntax**: Fluent API (`Where`, `With`, `Create`, `First`).
- **Eager Loading**: Supports relationships (`HasMany`, `BelongsTo`) via MongoDB `$lookup`.
- **Automatic Timestamps**: Manages `created_at` and `updated_at`.
- **Soft Deletes**: Built-in support for non-destructive deletions.
- **Hooks**: Lifecycle events (`BeforeCreate`, `AfterSave`).

## 📦 Installation

```bash
go get [github.com/kiamars-mirzaee/GoMongoquent](https://github.com/kiamars-mirzaee/GoMongoquent)
---

## ✨ Features

* Fluent **Eloquent-style query builder**
* Strong typing via **Go generics**
* Repository pattern (clean separation of concerns)
* Automatic timestamps (`CreatedAt`, `UpdatedAt`)
* Context-aware MongoDB operations
* No global state, no reflection-heavy hacks
* Easy to test & mock

---

## 📦 Installation

```bash
go get github.com/yourusername/gomongoquent
```

---

## 🚀 Quick Start

### 1️⃣ Connect to MongoDB

```go
ctx := context.Background()

db, err := orm.Connect(ctx, "mongodb://localhost:27017", "app")
if err != nil {
    log.Fatal(err)
}
```

---

### 2️⃣ Define a Model (Eloquent-style)

```go
type User struct {
    orm.BaseModel `bson:",inline"`
    Name  string `bson:"name"`
    Email string `bson:"email"`
}

func (User) CollectionName() string {
    return "users"
}
```

---

### 3️⃣ Create a Repository

```go
users := orm.NewRepository(db, User{})
```

---

### 4️⃣ Query Like Eloquent

```go
result, err := users.
    Query(ctx).
    Where("email", "test@mail.com").
    First()
```

```go
list, err := users.
    Query(ctx).
    Where("active", true).
    Sort("created_at", false).
    Limit(10).
    Get()
```

---

### 5️⃣ Create / Update / Delete

```go
user := User{Name: "John", Email: "john@mail.com"}
_ = users.Create(ctx, &user)
```

```go
_ = users.Update(ctx, user.ID, bson.M{
    "name": "John Updated",
})
```

```go
_ = users.Delete(ctx, user.ID)
```

```go
type User struct {
orm.BaseModel `bson:",inline"`
Name          string `bson:"name"`
Email         string `bson:"email"`
// Relationship Field
Posts         []Post `bson:"posts,omitempty"`
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
2. Query with Eager Loading
   Go

ctx := context.Background()
repo := orm.NewRepository(db, &User{})

// Fetch users named 'Alice' and include their Posts in one query
users, err := repo.Query(ctx).
Where("name", "Alice").
With("MyPosts"). // Triggers the Lookup
Get()

for _, user := range users {
fmt.Printf("User: %s has %d posts\n", user.Name, len(user.Posts))
}
```
---

# 🏗 Architecture
This library leverages Go Generics (1.18+) to provide a strictly typed experience. It wraps the official Mongo driver but abstracts away the complexity of bson.M map building and Aggregate pipelines.

Repository Pattern: Decouples data access logic.

Aggregation Pipelines: Used internally for performant joins.

```text
├── orm/
│   ├── db.go           # Mongo connection wrapper
│   ├── model.go        # Model & BaseModel
│   ├── query.go        # Fluent query builder
│   ├── repository.go  # CRUD layer
│   └── helpers.go      # timestamps, utilities
└── README.md
```

This structure follows **clean architecture principles** and avoids the ActiveRecord anti-pattern while preserving Eloquent’s DX.

---

## 🧪 Testing Strategy

* Repositories are easy to mock
* No global DB state
* Mongo calls isolated behind interfaces

Perfect for:

* Unit tests (mock repositories)
* Integration tests (real Mongo container)

---

## 🛣️ Roadmap
* [ ] Aggregation pipeline builder
* [ ] Transaction support
* [ ] CLI code generator

---

## 👨‍💻 Why This Project Exists

Many Go ORMs are either:

* Too magical
* Too low-level
* Or not developer-friendly

GoMongoquent aims to sit **right in the middle**:

> *Eloquent-like ergonomics with Go-level explicitness.*

This repository also serves as a **real-world example of senior Go design decisions**.

---

## 📄 License

MIT

---

## ⭐ If You Like This

* Star the repo ⭐
* Fork it 🍴
* Or use it as inspiration for your own internal tooling

> Built by a Go dev who secretly misses Laravel 😄
