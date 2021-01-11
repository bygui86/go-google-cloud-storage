package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/bygui86/multi-profile/v2"

	"github.com/bygui86/go-google-cloud-storage/config"
	"github.com/bygui86/go-google-cloud-storage/gcs"
	"github.com/bygui86/go-google-cloud-storage/utils"
)

var gcsClient *storage.Client

func main() {
	cfg := config.LoadConfig()

	if cfg.EnableProfiling {
		defer profile.CPUProfile(&profile.Config{}).Start().Stop()
		defer profile.MemProfile(&profile.Config{}).Start().Stop()
		defer profile.GoroutineProfile(&profile.Config{}).Start().Stop()
	}

	fmt.Println("Setup context")
	ctx := context.Background()

	fmt.Println("New GCP client")
	var clientErr error
	gcsClient, clientErr = gcs.NewClient(ctx)
	if clientErr != nil {
		fmt.Printf("ERROR - GCP new client creation failed: %s \n", clientErr.Error())
		os.Exit(501)
	}
	defer gcs.CloseClient(gcsClient)

	actionErr := performAction(ctx, cfg)
	if actionErr != nil {
		fmt.Printf("ERROR - Action failed: %s \n", actionErr.Error())
		os.Exit(501)
	}

	time.Sleep(3 * time.Second)
}

func performAction(ctx context.Context, cfg *config.Config) error {
	switch cfg.Level {
	case config.BucketLevel:
		return performBucketAction(ctx, cfg)
	case config.ObjectLevel:
		return performObjectAction(ctx, cfg)
	default:
		return fmt.Errorf("unknown level %s", cfg.Level)
	}
}

func performBucketAction(ctx context.Context, cfg *config.Config) error {
	fmt.Println("Perform BUCKET action")
	switch cfg.Action {
	case config.ListAction:
		fmt.Printf("List buckets in project %s \n", cfg.GcpProjectId)
		return gcs.ListBuckets(ctx, gcsClient, cfg.GcpProjectId)
	case config.CreateAction:
		fmt.Printf("Create bucket %s in project %s \n", cfg.GcsBucketName, cfg.GcpProjectId)
		return gcs.CreateBucket(ctx, gcsClient, cfg.GcpProjectId,
			cfg.GcsBucketName, string(cfg.GcsStorageClass), cfg.GcsLocation)
	case config.ExistAction:
		fmt.Printf("Check if bucket %s exists \n", cfg.GcsBucketName)
		err := gcs.CheckBucketExistence(ctx, gcsClient, cfg.GcsBucketName)
		if err == nil {
			fmt.Printf("Bucket %s found \n", cfg.GcsBucketName)
		} else {
			fmt.Printf("Bucket %s not found \n", cfg.GcsBucketName)
		}
		return nil
	case config.InfoAction:
		fmt.Printf("Get info about bucket %s \n", cfg.GcsBucketName)
		attrs, err := gcs.GetBucketMetadata(ctx, gcsClient, cfg.GcsBucketName)
		if err != nil {
			return err
		}
		gcs.PrintBucketMetadata(attrs)
		return nil
	// case config.DeleteAction:
	// 	return nil
	default:
		return fmt.Errorf("unknown bucket action %s", cfg.Level)
	}
}

func performObjectAction(ctx context.Context, cfg *config.Config) error {
	fmt.Println("Perform OBJECT action")
	switch cfg.Action {
	case config.ListAction:
		fmt.Printf("List objects in bucket %s \n", cfg.GcsBucketName)
		return gcs.ListObjects(ctx, gcsClient, cfg.GcsBucketName)
	case config.UploadAction:
		fmt.Printf("Upload file %s as object %s to bucket %s \n",
			cfg.UploadFilePath, cfg.GcsObjectName, cfg.GcsBucketName)
		file := utils.GetFile(cfg.UploadFilePath)
		upErr := gcs.UploadObject(ctx, gcsClient, cfg.GcsBucketName, cfg.GcsObjectName, file)
		readErr := file.Close()
		if readErr != nil {
			fmt.Printf("ERROR - File reader closing failed: %s \n", readErr.Error())
		}
		return upErr
	case config.DownloadAction:
		fmt.Printf("Download object %s from bucket %s to file %s \n",
			cfg.GcsObjectName, cfg.GcsBucketName, cfg.DownloadFilePath)
		data, downErr := gcs.DownloadObject(ctx, gcsClient, cfg.GcsBucketName, cfg.GcsObjectName)
		if downErr != nil {
			return downErr
		}
		return utils.WriteToFile(cfg.DownloadFilePath, data)
	case config.ExistAction:
		fmt.Printf("Check if object %s exists in bucket %s \n", cfg.GcsObjectName, cfg.GcsBucketName)
		err := gcs.CheckObjectExistence(ctx, gcsClient, cfg.GcsBucketName, cfg.GcsObjectName)
		if err == nil {
			fmt.Printf("Object %s found in bucket %s \n", cfg.GcsObjectName, cfg.GcsBucketName)
		} else {
			fmt.Printf("Object %s not found in bucket %s \n", cfg.GcsObjectName, cfg.GcsBucketName)
		}
		return nil
	case config.InfoAction:
		fmt.Printf("Get info about object %s in bucket %s \n", cfg.GcsObjectName, cfg.GcsBucketName)
		attrs, err := gcs.GetObjectMetadata(ctx, gcsClient, cfg.GcsBucketName, cfg.GcsObjectName)
		if err != nil {
			return err
		}
		gcs.PrintObjectMetadata(attrs)
		return nil
	case config.DeleteAction:
		fmt.Printf("Delete object %s from bucket %s \n", cfg.GcsObjectName, cfg.GcsBucketName)
		return gcs.DeleteObject(ctx, gcsClient, cfg.GcsBucketName, cfg.GcsObjectName)
	default:
		return fmt.Errorf("unknown object action %s", cfg.Level)
	}
}
