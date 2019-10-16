package gcpobject

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

func CheckExistance(ctx context.Context, client *storage.Client, bucketName, objectName string) bool {

	_, err := client.Bucket(bucketName).Object(objectName).Attrs(ctx)
	if err != nil {
		fmt.Printf("Failed to retrieve object %v: %v\n", bucketName, err.Error())
		return false
	}
	return true
}

func GetMetadata(ctx context.Context, client *storage.Client, bucketName, objectName string) (*storage.ObjectAttrs, error) {

	attrs, err := client.Bucket(bucketName).Object(objectName).Attrs(ctx)
	if err != nil {
		fmt.Printf("Failed to retrieve object %v attributes: %v\n", bucketName, err.Error())
		return nil, err
	}
	return attrs, nil
}

func PrintMetadata(ctx context.Context, client *storage.Client, bucketName, objectName string) {

	attrs, err := GetMetadata(ctx, client, bucketName, objectName)
	if err != nil {
		fmt.Printf("Failed to print object %v attributes: %v\n", bucketName, err.Error())
		return
	}
	fmt.Printf("Bucket: %v\n", attrs.Bucket)
	fmt.Printf("CacheControl: %v\n", attrs.CacheControl)
	fmt.Printf("ContentDisposition: %v\n", attrs.ContentDisposition)
	fmt.Printf("ContentEncoding: %v\n", attrs.ContentEncoding)
	fmt.Printf("ContentLanguage: %v\n", attrs.ContentLanguage)
	fmt.Printf("ContentType: %v\n", attrs.ContentType)
	fmt.Printf("Crc32c: %v\n", attrs.CRC32C)
	fmt.Printf("Generation: %v\n", attrs.Generation)
	fmt.Printf("KmsKeyName: %v\n", attrs.KMSKeyName)
	fmt.Printf("Md5Hash: %v\n", attrs.MD5)
	fmt.Printf("MediaLink: %v\n", attrs.MediaLink)
	fmt.Printf("Metageneration: %v\n", attrs.Metageneration)
	fmt.Printf("Name: %v\n", attrs.Name)
	fmt.Printf("Size: %v\n", attrs.Size)
	fmt.Printf("StorageClass: %v\n", attrs.StorageClass)
	fmt.Printf("TimeCreated: %v\n", attrs.Created)
	fmt.Printf("Updated: %v\n", attrs.Updated)
	fmt.Printf("Event-based hold enabled? %t\n", attrs.EventBasedHold)
	fmt.Printf("Temporary hold enabled? %t\n", attrs.TemporaryHold)
	fmt.Printf("Retention expiration time %v\n", attrs.RetentionExpirationTime)
	fmt.Print("\n\nMetadata\n")
	for key, value := range attrs.Metadata {
		fmt.Printf("\t%v = %v\n", key, value)
	}
}
