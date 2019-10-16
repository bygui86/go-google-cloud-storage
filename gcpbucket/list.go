package gcpbucket

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func List(ctx context.Context, client *storage.Client, projectId string) error {

	fmt.Printf("Buckets in project %s:\n", projectId)
	it := client.Buckets(ctx, projectId)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		fmt.Printf("    %s\n", battrs.Name)
	}
	return nil
}
