package core

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestInitEtcd(t *testing.T) {
	client := InitEtcd("154.8.197.123:2379")
	key := "auth_api"
	value := "127.0.0.1:20021"
	_, err := client.Put(context.Background(), key, value)
	if err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println("Put key:", key, "value:", value, "success")
	}
	getResponse, err := client.Get(context.Background(), "auth_api")
	if err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println(string(getResponse.Kvs[0].Value))
	}
}
