package gcs

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// LIST

func ListBuckets(ctx context.Context, client *storage.Client, projectId string) error {
	bucketIterator := client.Buckets(ctx, projectId)
	for {
		attrs, err := bucketIterator.Next()
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

// CREATE

func CreateBucket(ctx context.Context, client *storage.Client, projectId,
	bucketName, storageClass, location string) error {

	bucket := client.Bucket(bucketName)
	err := bucket.Create(ctx, projectId, &storage.BucketAttrs{
		StorageClass: storageClass,
		Location:     location,
	})
	if err != nil {
		return err
	}
	return nil
}

// EXIST

func CheckBucketExistence(ctx context.Context, client *storage.Client, bucketName string) error {
	_, err := client.Bucket(bucketName).Attrs(ctx)
	return err
}

// INFO

func GetBucketMetadata(ctx context.Context, client *storage.Client, bucketName string) (*storage.BucketAttrs, error) {
	return client.Bucket(bucketName).Attrs(ctx)
}

func PrintBucketMetadata(attrs *storage.BucketAttrs) {
	fmt.Printf("Name: %v \n", attrs.Name)
	fmt.Printf("Location: %v \n", attrs.Location)
	fmt.Printf("LocationType: %v \n", attrs.LocationType)
	fmt.Printf("StorageClass: %v \n", attrs.StorageClass)
	fmt.Printf("Created: %v \n", attrs.Created)
	fmt.Printf("MetaGeneration: %v \n", attrs.MetaGeneration)
	fmt.Printf("PredefinedACL: %v \n", attrs.PredefinedACL)
	if attrs.Encryption != nil {
		fmt.Printf("Encryption: \n")
		fmt.Printf("\tDefaultKmsKeyName: %v \n", attrs.Encryption.DefaultKMSKeyName)
	}
	if attrs.Website != nil {
		fmt.Printf("Website: \n")
		fmt.Printf("\tIndexPage: %v \n", attrs.Website.MainPageSuffix)
		fmt.Printf("\tNotFoundPage: %v \n", attrs.Website.NotFoundPage)
	}
	fmt.Printf("DefaultEventBasedHold: %v \n", attrs.DefaultEventBasedHold)
	if attrs.RetentionPolicy != nil {
		fmt.Printf("RetentionPolicy: \n")
		fmt.Printf("\tEffectiveTime: %v \n", attrs.RetentionPolicy.EffectiveTime)
		fmt.Printf("\tRetentionPeriod: %v \n", attrs.RetentionPolicy.RetentionPeriod)
		fmt.Printf("\tIsLocked: %v \n", attrs.RetentionPolicy.IsLocked)
	}
	fmt.Printf("RequesterPays: %v \n", attrs.RequesterPays)
	fmt.Printf("VersioningEnabled: %v \n", attrs.VersioningEnabled)
	if attrs.Logging != nil {
		fmt.Printf("Logging: \n")
		fmt.Printf("\tLogBucket: %v \n", attrs.Logging.LogBucket)
		fmt.Printf("\tLogObjectPrefix: %v \n", attrs.Logging.LogObjectPrefix)
	}
	if attrs.Labels != nil {
		fmt.Printf("Labels: \n")
		for key, value := range attrs.Labels {
			fmt.Printf("\t%v = %v \n", key, value)
		}
	}
}

// TODO DELETE
