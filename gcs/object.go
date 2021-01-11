package gcs

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// LIST

func ListObjects(ctx context.Context, client *storage.Client, bucketName string) error {
	objectIterator := client.Bucket(bucketName).Objects(ctx, nil)
	for {
		attrs, err := objectIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("\t%s \n", attrs.Name)
	}
	return nil
}

// TODO ListObjectsFilteredByQuery

// UPLOAD

func UploadObject(ctx context.Context, client *storage.Client, bucketName, objectName string, sourceFile *os.File) error {
	writer := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	_, copyErr := io.Copy(writer, sourceFile)
	if copyErr != nil {
		return copyErr
	}
	// PLEASE NOTE: in the google documentation is written to wait for the ObjectWriter closing instaed of defering it.
	// defer writer.Close()
	closeErr := writer.Close()
	if closeErr != nil {
		return closeErr
	}
	return nil
}

// DOWNLOAD

func DownloadObject(ctx context.Context, client *storage.Client, bucketName, objectName string) ([]byte, error) {
	reader, readerErr := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if readerErr != nil {
		return nil, readerErr
	}
	defer closeReader(reader)

	data, readErr := ioutil.ReadAll(reader)
	if readErr != nil {
		return nil, readErr
	}

	return data, nil
}

func closeReader(reader *storage.Reader) {
	err := reader.Close()
	if err != nil {
		fmt.Printf("ERROR - Storage reader closing failed: %s \n", err.Error())
	}
}

// INFO

func CheckObjectExistence(ctx context.Context, client *storage.Client, bucketName, objectName string) error {
	_, err := client.Bucket(bucketName).Object(objectName).Attrs(ctx)
	return err
}

func GetObjectMetadata(ctx context.Context, client *storage.Client, bucketName, objectName string) (*storage.ObjectAttrs, error) {
	return client.Bucket(bucketName).Object(objectName).Attrs(ctx)
}

func PrintObjectMetadata(attrs *storage.ObjectAttrs) {
	fmt.Printf("Bucket: %v \n", attrs.Bucket)
	fmt.Printf("CacheControl: %v \n", attrs.CacheControl)
	fmt.Printf("ContentDisposition: %v \n", attrs.ContentDisposition)
	fmt.Printf("ContentEncoding: %v \n", attrs.ContentEncoding)
	fmt.Printf("ContentLanguage: %v \n", attrs.ContentLanguage)
	fmt.Printf("ContentType: %v \n", attrs.ContentType)
	fmt.Printf("Crc32c: %v \n", attrs.CRC32C)
	fmt.Printf("Generation: %v \n", attrs.Generation)
	fmt.Printf("KmsKeyName: %v \n", attrs.KMSKeyName)
	fmt.Printf("Md5Hash: %v \n", attrs.MD5)
	fmt.Printf("MediaLink: %v \n", attrs.MediaLink)
	fmt.Printf("Metageneration: %v \n", attrs.Metageneration)
	fmt.Printf("Name: %v \n", attrs.Name)
	fmt.Printf("Size: %v \n", attrs.Size)
	fmt.Printf("StorageClass: %v \n", attrs.StorageClass)
	fmt.Printf("TimeCreated: %v \n", attrs.Created)
	fmt.Printf("Updated: %v \n", attrs.Updated)
	fmt.Printf("Event-based hold enabled? %t\n", attrs.EventBasedHold)
	fmt.Printf("Temporary hold enabled? %t\n", attrs.TemporaryHold)
	fmt.Printf("Retention expiration time %v \n", attrs.RetentionExpirationTime)
	fmt.Println("Metadata:")
	for key, value := range attrs.Metadata {
		fmt.Printf("\t%v = %v \n", key, value)
	}
}

// DELETE

func DeleteObject(ctx context.Context, client *storage.Client, bucketName, objectName string) error {
	return client.Bucket(bucketName).Object(objectName).Delete(ctx)
}
