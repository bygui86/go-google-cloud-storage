package gcpobject

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

func Delete(ctx context.Context, client *storage.Client, bucketName, objectName string) error {

	err := client.Bucket(bucketName).Object(objectName).Delete(ctx)
	if err != nil {
		fmt.Printf("Failed to delete object %v in bucket %v: %v\n", objectName, bucketName, err.Error())
		return err
	}
	return nil
}
