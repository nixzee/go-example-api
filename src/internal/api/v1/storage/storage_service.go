package storage

import (
	"context"
	"time"
)

type StorageService interface {
	ListFilesInContainer(ctx context.Context, containerName string) (blobInfos []BlobInfo, err error)
}

var _ StorageService = (*storageService)(nil)

type storageService struct {
	accountName string
	sasToken    string
}

func NewStorageService(accountName, sasToken string) StorageService {
	return &storageService{
		accountName: accountName,
		sasToken:    sasToken,
	}
}

func (s *storageService) ListFilesInContainer(ctx context.Context, containerName string) (blobInfos []BlobInfo, err error) {
	// TODO: Implement if when needed. This is to show how to build a example API
	blobInfos = []BlobInfo{
		{
			Name:         "file1.txt",
			Size:         1024,
			LastModified: time.Now().Add(-24 * time.Hour), // 1 day ago
		},
		{
			Name:         "image1.jpg",
			Size:         2048,
			LastModified: time.Now().Add(-48 * time.Hour), // 2 days ago
		},
		{
			Name:         "document.pdf",
			Size:         5120,
			LastModified: time.Now().Add(-72 * time.Hour), // 3 days ago
		},
	}

	return
}
