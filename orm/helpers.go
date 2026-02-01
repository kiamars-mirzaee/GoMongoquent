package orm

import (
	"reflect"
	"time"
)

// ---------- Helpers ----------

func setTimestamps[T any](model *T, creating bool) {
	v := reflect.ValueOf(model).Elem()

	if field := v.FieldByName("CreatedAt"); field.IsValid() && creating {
		field.Set(reflect.ValueOf(time.Now()))
	}

	if field := v.FieldByName("UpdatedAt"); field.IsValid() {
		field.Set(reflect.ValueOf(time.Now()))
	}
}
