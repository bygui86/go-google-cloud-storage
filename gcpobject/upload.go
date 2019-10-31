package gcpobject

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
)

func Upload(ctx context.Context, client *storage.Client, bucketName, objectName string, sourceFile *os.File) error {

	writer := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	_, copyErr := io.Copy(writer, sourceFile)
	if copyErr != nil {
		fmt.Printf("Failed to write data into object %v in bucket %v: %v\n", objectName, bucketName, copyErr.Error())
		return copyErr
	}
	// PLEASE NOTE: in the google documentation is written to wait for the ObjectWriter closing instaed of defering it.
	defer writer.Close()
	// closeErr := writer.Close()
	// if closeErr != nil {
	// 	fmt.Printf("Failed to close object %v writer to bucket %v: %v\n", objectName, bucketName, closeErr.Error())
	// 	return closeErr
	// }

	return nil
}
