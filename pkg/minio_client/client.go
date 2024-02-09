package minio_client

import (
	"sync"

    "github.com/minio/minio-go/v7"
)

type ClientManager struct {
	client *minio.Client
	once   sync.Once
}

var clientManager *ClientManager

func GetClient(endpoint, accessKey, secretKey string) (*minio.Client, error) {
	if clientManager == nil {
		clientManager = &ClientManager{}
	}

	clientManager.once.Do(func() {
        clientManager.client, _ = NewClient(endpoint, accessKey, secretKey)
    })

	return clientManager.client, nil
}
