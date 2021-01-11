package gcs

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

func NewClient(ctx context.Context) (*storage.Client, error) {
	return storage.NewClient(ctx)
}

func CloseClient(client *storage.Client) {
	err := client.Close()
	if err != nil {
		fmt.Printf("ERROR - GCP client closing failed: %s \n", err.Error())
	}
}
