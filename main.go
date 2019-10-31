package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/bygui86/go-google-cloud-storage/gcpbucket"
	"github.com/bygui86/go-google-cloud-storage/gcpclient"
	"github.com/bygui86/go-google-cloud-storage/gcpobject"
	"github.com/pkg/profile"
)

const (
	projectId  = "gcp-sample-project"
	bucketName = "golang-tests"
	objectName = "notes-test.txt"

	storageClassStd  = "STANDARD"
	storageClassNear = "NEARLINE"
	storageClassCold = "COLDLINE"

	locationEu        = "europe"
	locationUs        = "us"
	locationAs        = "asia"
	locationUsCentral = "us-central1"

	upFilepath   = "notes_up.txt"
	downFilepath = "notes_down.txt"
)

func main() {

	// PROFILING
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	memProf := profile.Start(profile.MemProfile, profile.ProfilePath("."))
	// defer profile.Start(profile.ThreadcreationProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.NoShutdownHook, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MutexProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()

	// setup context
	fmt.Printf("*** setup context\n")
	ctx := context.Background()

	// create client
	fmt.Printf("*** create client\n")
	client := gcpclient.New(ctx)

	if client != nil {
		// Closing the client safely cleans up background resources.
		defer client.Close()

		bucketActions(ctx, client)
		// objectActions(ctx, client)

		go objectBench(ctx, client, "1")
		go objectBench(ctx, client, "2")
		go objectBench(ctx, client, "3")
	}

	time.Sleep(60 * time.Second)

	memProf.Stop()
}

func bucketActions(ctx context.Context, client *storage.Client) {

	// list all buckets in project
	fmt.Printf("*** list buckects in project %v\n", projectId)
	gcpbucket.List(ctx, client, projectId)

	// check bucket existance
	var bucketErr error
	fmt.Printf("*** check existance of bucket %v\n", bucketName)
	if !gcpbucket.CheckExistance(ctx, client, bucketName) {
		fmt.Printf("*** bucket %v does not exist\n", bucketName)
		// create a bucket
		fmt.Printf("*** create bucket %v\n", bucketName)
		bucketErr = gcpbucket.Create(ctx, client, projectId, bucketName, storageClassStd, locationUs)
	}

	if bucketErr == nil {
		// describe a bucket
		fmt.Printf("*** describe bucket %v\n", bucketName)
		gcpbucket.PrintMetadata(ctx, client, bucketName)
}

func objectBench(ctx context.Context, client *storage.Client, id string) {

	for {
		file := getFile(upFilepath)
		// defer file.Close()
		loopObjectName := objectFolder + id + "/" + objectName + "_" + id + "_" + getFormattedTimeWithMs(time.Now()) + objectNameSuffix
		fmt.Printf("upload object %v to bucket %v\n", loopObjectName, bucketName)
		upObjectErr := gcpobject.Upload(ctx, client, bucketName, loopObjectName, file)
		if upObjectErr != nil {
			fmt.Printf("FAILED upload object %v to bucket %v: %v\n", loopObjectName, bucketName, upObjectErr.Error())
		}
		// time.Sleep(10 * time.Millisecond)
		file.Close()
	}
}

func objectActions(ctx context.Context, client *storage.Client) {

	// list all objects in bucket
	fmt.Printf("*** list objects in bucket %v\n", bucketName)
	gcpobject.List(ctx, client, bucketName)

	// check object existance
	var upObjectErr error
	fmt.Printf("*** check existance of object %v in bucket %v\n", objectName, bucketName)
	if !gcpobject.CheckExistance(ctx, client, bucketName, objectName) {
		// upload an object
		file := getFile(upFilepath)
		defer file.Close()
		fmt.Printf("*** upload object %v to bucket %v\n", objectName, bucketName)
		upObjectErr = gcpobject.Upload(ctx, client, bucketName, objectName, file)
	}
	if upObjectErr == nil {
		// describe a object
		// fmt.Printf("*** describe object %v in bucket %v\n", objectName, bucketName)
		// gcpobject.PrintMetadata(ctx, client, bucketName, objectName)

		// download an object
		fmt.Printf("*** download object %v from bucket %v\n", objectName, bucketName)
		data, downObjectErr := gcpobject.Download(ctx, client, bucketName, objectName)
		if downObjectErr == nil {
			writeToFile(downFilepath, data)
		}
	}
}

func getFile(filepath string) *os.File {

	f, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	return f
}

func writeToFile(filepath string, data []byte) error {

	writeErr := ioutil.WriteFile(filepath, data, 0644)
	if writeErr != nil {
		fmt.Printf("Failed to write data to file %v: %v", filepath, writeErr.Error())
		return writeErr
	}
	return nil
}

func getFormattedTimeWithMs(time time.Time) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			time.Format("15:04:05.999999"), ":", "-",
		), ".", "-",
	)
}
