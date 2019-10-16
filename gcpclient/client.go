package gcpclient

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

func New(ctx context.Context) *storage.Client {

	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err.Error())
		return nil
	}
	return client
}
