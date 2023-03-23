package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/s3blob" // 対象の driver を blank importする
)

func main() {
	ctx := context.TODO()

	bucket, err := blob.OpenBucket(ctx, "s3://mybucket")
	if err != nil {
		log.Fatal(err)
	}
	defer bucket.Close()

	var token = blob.FirstPageToken
	for {
		opts := &blob.ListOptions{
			Prefix: "oreilly/2022/01/31",
		}
		objs, nextToken, err := bucket.ListPage(ctx, token, 10, opts)
		if err != nil {
			log.Fatalf("list objects, %v", err)
		}
		for _, obj := range objs {
			fmt.Printf("Name:%s LastModified:%s\n", obj.Key, obj.ModTime.Format(time.RFC3339))
		}

		if nextToken == nil {
			break
		}
		token = nextToken
	}

}
