package gcpbucket

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

func Create(ctx context.Context, client *storage.Client, projectId, bucketName, storageClass, location string) error {

	bucket := client.Bucket(bucketName)
	err := bucket.Create(ctx, projectId, &storage.BucketAttrs{
		StorageClass: storageClass,
		Location:     location,
	})
	if err != nil {
		fmt.Printf("Failed to create bucket %v: %v\n", bucketName, err.Error())
		return err
	}

	fmt.Printf("Bucket %v created.\n", bucketName)
	return nil
}
