package sources

import (
	"context"
	"news/internal/storage"
	"news/internal/storage/custom"
	"news/internal/storage/mediapost"
	"strings"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetMediaPostList(client *storage.Client, context context.Context) (result []*custom.MediaPost) {
	values := client.MediaPost.Query().Where(mediapost.Status(true)).AllX(context)
	for _, value := range values {
		result = append(result, &custom.MediaPost{
			Type:        custom.MediaPost_Type(custom.MediaPost_Type_value[strings.ToUpper(value.Type.String())]),
			Date:        timestamppb.New(value.Date),
			Description: value.Description,
			Content:     value.Content,
			Source:      value.Source,
			Title:       value.Title,
			Image:       value.Image,
			Link:        value.Link,
			Logo:        value.Logo,
			Live:        value.Live,
		})
	}
	return result
}
