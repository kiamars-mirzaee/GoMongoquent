# GoMongoquent 🐹🍃

> **Laravel Eloquent–inspired MongoDB ORM for Go (without the magic)**

GoMongoquent is a lightweight, strongly-typed MongoDB abstraction for Go developers who love the **developer experience of Laravel Eloquent**, but still want **idiomatic, explicit, and performant Go**.

This project is designed to be:

* 💼 **Portfolio‑ready** (clean architecture, senior-level patterns)
* 🧠 **Easy to reason about** (no hidden magic)
* ⚡ **Production friendly** (context-aware, testable, predictable)

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

---

## 🧱 Architecture

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

* [ ] Query scopes (`Active()`, `Recent()`)
* [ ] Soft deletes
* [ ] Relations (`HasMany`, `BelongsTo`)
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
# GoMongoquent
# GoMongoquent
# GoMongoquent
