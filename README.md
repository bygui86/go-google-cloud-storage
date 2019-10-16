
# Go Google Cloud Storage

## Build & run

```
export GOOGLE_APPLICATION_CREDENTIALS="$PWD/gc-storage-credentials.json"
export GO111MODULE=on
go run main.go
```

---

## Download with gsutil
```
gsutil -m cp -r gs://<BUCKET_NAME>/<FOLDER_NAME>/ <DESTINATION_PATH>
```

---

## Links

* https://cloud.google.com/storage/docs/reference/libraries#client-libraries-install-go
* https://github.com/googleapis/google-cloud-go
* https://cloud.google.com/storage/docs/request-rate
* https://cloud.google.com/storage/quotas

### Buckets
* https://cloud.google.com/storage/docs/creating-buckets#storage-create-bucket-go
* https://cloud.google.com/storage/docs/locations#location-r
* https://cloud.google.com/storage/docs/listing-buckets#storage-list-buckets-go
* https://cloud.google.com/storage/docs/deleting-objects

### Objects
* https://cloud.google.com/storage/docs/uploading-objects#storage-upload-object-go
* https://cloud.google.com/storage/docs/downloading-objects#storage-download-object-go
* https://cloud.google.com/storage/docs/listing-objects#storage-list-objects-go
* https://cloud.google.com/storage/docs/naming
