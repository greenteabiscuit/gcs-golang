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

	// オブジェクトのReaderを作成
	bucketName := "experimental-bucket-tt"   // e.g. example-bucket
	objectPath := "sample-object/sample.txt" // e.g. foo/var/sample.txt

	writer := client.Bucket(bucketName).Object(objectPath).NewWriter(ctx)
	if _, err := io.Copy(writer, f); err != nil {
		panic(err)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}

	log.Println("done")
}
