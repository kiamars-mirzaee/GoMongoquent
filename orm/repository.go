
// ---------- Repository (Eloquent-style) ----------

type Repository[T Model] struct {
	db    *DB
	model T
}

func NewRepository[T Model](db *DB, model T) *Repository[T] {
	return &Repository[T]{db: db, model: model}
}

func (r *Repository[T]) Query(ctx context.Context) *Query[T] {
	return NewQuery(ctx, r.db, r.model)
}

func (r *Repository[T]) Create(ctx context.Context, m *T) error {
	col := r.db.Collection(r.model)

	setTimestamps(m, true)
	_, err := col.InsertOne(ctx, m)
	return err
}

func (r *Repository[T]) Update(ctx context.Context, id primitive.ObjectID, data bson.M) error {
	col := r.db.Collection(r.model)

	data["updated_at"] = time.Now()
	res, err := col.UpdateByID(ctx, id, bson.M{"$set": data})
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("document not found")
	}
	return nil
}

func (r *Repository[T]) Delete(ctx context.Context, id primitive.ObjectID) error {
	col := r.db.Collection(r.model)
	_, err := col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

