package gcpobject

import (
	"context"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/storage"
)

func Download(ctx context.Context, client *storage.Client, bucketName, objectName string) ([]byte, error) {

	reader, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		fmt.Printf("Failed to open reader to object %v in bucket %v: %v\n", objectName, bucketName, err.Error())
		return nil, err
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("Failed to read data from object %v in bucket %v: %v\n", objectName, bucketName, err.Error())
		return nil, err
	}

	return data, nil
}
