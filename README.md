# gcs-golang

鍵ファイルはGCPからダウンロードする必要があります。
Key files need to be downloaded from GCP.

```
$ docker-compose up

// In different tab
$ docker exec -it containername bash

// Inside container, install necessary packages
/go/src# go mod init example.com/gcs/write
go: creating new go.mod: module example.com/gcs/write
/go/src# go get -u cloud.google.com/go/storage
...

// Run the code
/go/src# go run main.go
2020/12/29 01:34:33 done

```
