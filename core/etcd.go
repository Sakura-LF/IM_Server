package core

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func InitEtcd(add string) *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{add},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return client
}

/*
docker run -d --name Etcd-server \
--publish 2379:2379 \
--publish 2380:2380 \
--env ALLOW_NONE_AUTHENTICATION=yes \
--env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
bitnami/etcd:latest
*/
