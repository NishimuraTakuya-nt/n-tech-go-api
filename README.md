# n-tech-grpc-go

このプロジェクトは、Go言語とgRPCを使用したサンプルアプリケーションです。クライアントとサーバー間で簡単なメッセージのやり取りを行います。

## 概要

このアプリケーションは以下の機能を提供します：

1. クライアントからサーバーへの挨拶メッセージの送信
2. サーバーからクライアントへの応答メッセージの送信

## 必要条件

- Go 1.16以上
- gRPC
- Protocol Buffers

## セットアップ

1. このリポジトリをクローンします：

   ```
   git clone https://github.com/NishimuraTakuya-nt/n-tech-grpc-go.git
   ```

2. 必要な依存関係をインストールします：

   ```
    go mod download
   ```

3. Protocol Buffersファイルからコードを生成します：

   ```
    Makefile:generate
   ```

## 使用方法

1. サーバーを起動します：

   ```
    go run server/main.go
   ```

2. 別のターミナルでクライアントを実行します：

   ```
    go run client/main.go
   ```

3. クライアントの出力を確認します。サーバーから受信した応答メッセージが表示されるはずです。

## ファイル構成

- `proto/`: Protocol Buffersの定義ファイル
- `server/`: gRPCサーバーの実装
- `client/`: gRPCクライアントの実装

## 参考資料

- [gRPC公式ドキュメント](https://grpc.io/docs/)
- [Go言語公式ドキュメント](https://golang.org/doc/)
