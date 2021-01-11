package config

const (
	// level
	BucketLevel Level = "bucket"
	ObjectLevel Level = "object"

	// action
	ListAction     Action = "list"
	CreateAction   Action = "create"
	DeleteAction   Action = "delete"
	InfoAction     Action = "info"
	ExistAction    Action = "exist"
	UploadAction   Action = "upload"
	DownloadAction Action = "download"

	// storage-class
	StandardStorageClass StorageClass = "STANDARD"
	NearlineStorageClass StorageClass = "NEARLINE"
	ColdlineStorageClass StorageClass = "COLDLINE"
)
