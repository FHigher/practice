package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv clientv3.KV
		getResp *clientv3.GetResponse
		putResp *clientv3.PutResponse
	)
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	client, err = clientv3.New(config)
	if nil != err {
		fmt.Println(err)
		return
	}

	kv = clientv3.NewKV(client)

	if putResp, err = kv.Put(context.TODO(), "/learn/one", "bye", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Revision:", putResp.Header.Revision)
		if putResp.PrevKv != nil {	// 打印hello
			fmt.Println("PrevValue:", string(putResp.PrevKv.Value))
		}
	}


	if getResp, err = kv.Get(context.TODO(), "/learn/one", clientv3.WithCountOnly()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Revision", getResp.Header.Revision)
		fmt.Println(getResp.Kvs, getResp.Count)
	}
}
