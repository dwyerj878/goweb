package vks

import (
	"context"
	"hello/config"
	"log"

	"github.com/valkey-io/valkey-go"
)

type KEY struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func GetValues() []KEY {

	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{config.Get_config().ValkeyHost}})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	err = client.Do(ctx, client.B().Set().Key("key").Value("val").Nx().Build()).Error()
	if err != nil {
		log.Println(err)
	}

	messages, err := client.Do(ctx, client.B().Get().Key("key").Build()).ToString()
	defer client.Close()
	if err != nil {
		log.Println(err)
	}
	log.Println(messages)

	messages, err = client.Do(ctx, client.B().Get().Key("test_key").Build()).ToString()
	defer client.Close()
	if err != nil {
		log.Println(err)
	}
	log.Println(messages)
	return nil
}
