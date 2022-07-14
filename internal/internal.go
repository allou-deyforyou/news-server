package internal

import (
	"context"
	"io"
	"log"
	"math/rand"

	"news/internal/store"
	"news/internal/store/migrate"
	"news/internal/store/schema"

	"entgo.io/ent/dialect"
	"google.golang.org/protobuf/proto"
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

func ProtoDecode(r io.Reader, m proto.Message) error {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return proto.Unmarshal(bytes, m)
}

func ProtoEncode(w io.Writer, m proto.Message) error {
	bytes, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func ConvertToNewsTvPost(data []*store.NewsTvSource) []*schema.NewsTvPost {
	result := make([]*schema.NewsTvPost, len(data))
	for i, item := range data {
		result[i] = &schema.NewsTvPost{
			Description: item.Description,
			Source:      item.Source,
			Video:       item.Video,
			Logo:        item.Logo,
			Live:        item.Live,
		}
	}
	return result
}


func ConvertToNewsCategory(data map[string]string) []*schema.NewsCategory {
	result := make([]*schema.NewsCategory,0)
	for key, value := range data {
		result = append(result, &schema.NewsCategory{
			Value: value,
			Name: key,
		})
	}
	return result
}