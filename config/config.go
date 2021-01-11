package config

import (
	"fmt"

	"github.com/bygui86/go-google-cloud-storage/utils"
)

const (
	enableProfilingEnvVar  = "ENABLE_PROFILING"
	levelEnvVar            = "LEVEL"
	actionEnvVar           = "ACTION"
	gcpProjectIdEnvVar     = "GCP_PROJECT_ID"
	gcsBucketNameEnvVar    = "GCS_BUCKET_NAME"
	gcsObjectNameEnvVar    = "GCS_OBJ_NAME"
	gcsStorageClassEnvVar  = "GCS_STORAGE_CLASS"
	gcsLocationEnvVar      = "GCS_LOCATION"
	uploadFilePathEnvVar   = "UPLOAD_FILE_PATH"
	downloadFilePathEnvVar = "DOWNLOAD_FILE_PATH"

	enableProfilingDefault  = false
	levelDefault            = BucketLevel
	actionDefault           = ListAction
	gcpProjectIdDefault     = ""
	gcsBucketNameDefault    = ""
	gcsObjectNameDefault    = ""
	gcsStorageClassDefault  = StandardStorageClass
	gcsLocationDefault      = ""
	uploadFilePathDefault   = ""
	downloadFilePathDefault = "./"
)

func LoadConfig() *Config {
	fmt.Println("Load configurations")

	return &Config{
		EnableProfiling:  utils.GetBoolEnv(enableProfilingEnvVar, enableProfilingDefault),
		Level:            Level(utils.GetStringEnv(levelEnvVar, string(levelDefault))),
		Action:           Action(utils.GetStringEnv(actionEnvVar, string(actionDefault))),
		GcpProjectId:     utils.GetStringEnv(gcpProjectIdEnvVar, gcpProjectIdDefault),
		GcsBucketName:    utils.GetStringEnv(gcsBucketNameEnvVar, gcsBucketNameDefault),
		GcsObjectName:    utils.GetStringEnv(gcsObjectNameEnvVar, gcsObjectNameDefault),
		GcsStorageClass:  StorageClass(utils.GetStringEnv(gcsStorageClassEnvVar, string(gcsStorageClassDefault))),
		GcsLocation:      utils.GetStringEnv(gcsLocationEnvVar, gcsLocationDefault),
		UploadFilePath:   utils.GetStringEnv(uploadFilePathEnvVar, uploadFilePathDefault),
		DownloadFilePath: utils.GetStringEnv(downloadFilePathEnvVar, downloadFilePathDefault),
	}
}
