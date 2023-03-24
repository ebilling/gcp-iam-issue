package iamissue

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func simple() (*storage.Client, error) {
	ctx := context.Background()
	return storage.NewClient(ctx, option.WithCredentialsJSON([]byte{}))
}
