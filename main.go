package main

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func main() {
	credentialFilePath := "./key.json" // key.jsonはサービスアカウントを作成してゲットする

	// クライアントを作成する
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	// GCSオブジェクトを書き込むファイルの作成
	f, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	// ......................
	// オブジェクトのアップロード
	// ......................
	bucketName := "experimental-bucket-tt"   // e.g. example-bucket
	objectPath := "sample-object/sample.txt" // e.g. foo/var/sample.txt

	uploadWriter := client.Bucket(bucketName).Object(objectPath).NewWriter(ctx)
	if _, err := io.Copy(uploadWriter, f); err != nil {
		panic(err)
	}

	if err := uploadWriter.Close(); err != nil {
		panic(err)
	}
	log.Println("create file: done")

	// ......................
	// オブジェクトの移動
	// ......................
	dstObjectPath := "destination-folder/30/sample.txt"
	src := client.Bucket(bucketName).Object(objectPath)
	dst := client.Bucket(bucketName).Object(dstObjectPath)
	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		panic(err)
	}
	if err := src.Delete(ctx); err != nil {
		panic(err)
	}
	log.Println("move file: done")
}
