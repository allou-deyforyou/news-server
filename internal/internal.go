package internal

import (
	"context"
	"log"
	"math/rand"

	"news/internal/store"
	"news/internal/store/migrate"

	"entgo.io/ent/dialect"
)

func NewEntClient() *store.Client {
	client, err := store.Open(dialect.SQLite, "yola.db?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed connecting to sqlite: %v", err)
	}
	ctx := context.Background()
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func Shuffle[T any](data []T) []T {
	length := len(data)
	result := make([]T, length)
	perm := rand.Perm(length)
	for i, v := range perm {
		result[v] = data[i]
	}
	return result
}

func Remove[T any](data []T, f func(T, T) bool) []T {
	for i := 0; i < len(data); i++ {
		for j := len(data) - 1; j > i; j-- {
			item := data[i]
			last := data[j]
			if f(item, last) {
				data = append(data[:j], data[j+1:]...)
			}
		}
	}
	return data
}
