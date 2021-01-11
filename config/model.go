package config

type Config struct {
	EnableProfiling  bool
	Level            Level
	Action           Action
	GcpProjectId     string
	GcsBucketName    string
	GcsObjectName    string
	GcsStorageClass  StorageClass
	GcsLocation      string
	UploadFilePath   string
	DownloadFilePath string
}

type Level string

type Action string

type StorageClass string
