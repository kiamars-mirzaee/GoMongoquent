// ---------- Query Builder ----------

type Query[T Model] struct {
	ctx        context.Context
	db         *DB
	model      T
	filter     bson.M
	options    *options.FindOptions
	singleOpts *options.FindOneOptions
}

func NewQuery[T Model](ctx context.Context, db *DB, model T) *Query[T] {
	return &Query[T]{
		ctx:    ctx,
		db:     db,
		model:  model,
		filter: bson.M{},
	}
}

func (q *Query[T]) Where(field string, value any) *Query[T] {
	q.filter[field] = value
	return q
}

func (q *Query[T]) Limit(limit int64) *Query[T] {
	if q.options == nil {
		q.options = options.Find()
	}
	q.options.SetLimit(limit)
	return q
}

func (q *Query[T]) Sort(field string, asc bool) *Query[T] {
	order := 1
	if !asc {
		order = -1
	}

	if q.options == nil {
		q.options = options.Find()
	}
	q.options.SetSort(bson.M{field: order})
	return q
}

func (q *Query[T]) Get() ([]T, error) {
	col := q.db.Collection(q.model)

	cursor, err := col.Find(q.ctx, q.filter, q.options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(q.ctx)

	var results []T
	for cursor.Next(q.ctx) {
		var item T
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	return results, nil
}

func (q *Query[T]) First() (*T, error) {
	col := q.db.Collection(q.model)

	var result T
	err := col.FindOne(q.ctx, q.filter, q.singleOpts).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

