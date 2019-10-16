package gcpbucket

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

func CheckExistance(ctx context.Context, client *storage.Client, bucketName string) bool {

	_, err := client.Bucket(bucketName).Attrs(ctx)
	if err != nil {
		fmt.Printf("Failed to retrieve bucket %v: %v\n", bucketName, err.Error())
		return false
	}
	return true
}

func GetMetadata(ctx context.Context, client *storage.Client, bucketName string) (*storage.BucketAttrs, error) {

	attrs, err := client.Bucket(bucketName).Attrs(ctx)
	if err != nil {
		fmt.Printf("Failed to retrieve bucket %v attributes: %v\n", bucketName, err.Error())
		return nil, err
	}
	return attrs, nil
}

func PrintMetadata(ctx context.Context, client *storage.Client, bucketName string) {

	attrs, err := GetMetadata(ctx, client, bucketName)
	if err != nil {
		fmt.Printf("Failed to print bucket %v attributes: %v\n", bucketName, err.Error())
		return
	}

	fmt.Printf("BucketName: %v\n", attrs.Name)
	fmt.Printf("Location: %v\n", attrs.Location)
	fmt.Printf("LocationType: %v\n", attrs.LocationType)
	fmt.Printf("StorageClass: %v\n", attrs.StorageClass)
	fmt.Printf("TimeCreated: %v\n", attrs.Created)
	fmt.Printf("Metageneration: %v\n", attrs.MetaGeneration)
	fmt.Printf("PredefinedACL: %v\n", attrs.PredefinedACL)
	if attrs.Encryption != nil {
		fmt.Printf("DefaultKmsKeyName: %v\n", attrs.Encryption.DefaultKMSKeyName)
	}
	if attrs.Website != nil {
		fmt.Printf("IndexPage: %v\n", attrs.Website.MainPageSuffix)
		fmt.Printf("NotFoundPage: %v\n", attrs.Website.NotFoundPage)
	}
	fmt.Printf("DefaultEventBasedHold: %v\n", attrs.DefaultEventBasedHold)
	if attrs.RetentionPolicy != nil {
		fmt.Printf("RetentionEffectiveTime: %v\n", attrs.RetentionPolicy.EffectiveTime)
		fmt.Printf("RetentionPeriod: %v\n", attrs.RetentionPolicy.RetentionPeriod)
		fmt.Printf("RetentionPolicyIsLocked: %v\n", attrs.RetentionPolicy.IsLocked)
	}
	fmt.Printf("RequesterPays: %v\n", attrs.RequesterPays)
	fmt.Printf("VersioningEnabled: %v\n", attrs.VersioningEnabled)
	if attrs.Logging != nil {
		fmt.Printf("LogBucket: %v\n", attrs.Logging.LogBucket)
		fmt.Printf("LogObjectPrefix: %v\n", attrs.Logging.LogObjectPrefix)
	}
	if attrs.Labels != nil {
		fmt.Printf("\n\n\nLabels:")
		for key, value := range attrs.Labels {
			fmt.Printf("\t%v = %v\n", key, value)
		}
	}
}
