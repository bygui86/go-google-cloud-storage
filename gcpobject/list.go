package gcpobject

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func List(ctx context.Context, client *storage.Client, bucketName string) error {

	fmt.Printf("Objects in bucket %s:\n", bucketName)
	it := client.Bucket(bucketName).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		fmt.Printf("    %s\n", attrs.Name)
	}
	return nil
}
