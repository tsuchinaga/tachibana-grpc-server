# tachibana-grpc-server

立花証券APIを叩くためのgRPCサーバー

複数ツールから立花証券APIを利用するためにリクエストの受け口を一つにするために使う。

## build

`$ protoc --go_out=./tachibanapb --go-grpc_out=./tachibanapb tachibanapb/tachibana.proto`

* `--go_out`: protobufのmessageとかenumとかが吐かれる先
* `--go-grpc_out`: protobufのserviceから作られるserverとかclientが吐かれる先

## run

`$ go run cmd/tachibana_grpc_server.go -port=8900`

* `port`: ポート。デフォルト8900

## 定義

[protobufファイル](./tachibanapb/tachibana.proto)

[protobufドキュメント](https://tsuchinaga.gitlab.io/tachibana-grpc-server/#tachibanapb.TachibanaService)

## 注意

まだ開発中で全機能の実装はできていません。

[//]: # ([github.com/tsuchinaga/tachibana-grpc-server]&#40;https://github.com/tsuchinaga/tachibana-grpc-server&#41; にミラーリングしていますが、オリジナルは [gitlab.com/tsuchinaga/tachibana-grpc-server]&#40;https://gitlab.com/tsuchinaga/tachibana-grpc-server&#41; にあります。)
