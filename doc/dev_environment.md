# 環境構築

## 前提

- Dockerがインストールされていること

## GitHubからリポジトリをclone

```sh
git clone git@github.com:jin237/mdet.git
```

## コンテナ立ち上げ

```sh
# フォルダ移動
cd mdet

# Dockerイメージの構築
docker compose build --no-cache

# Dockerコンテナ立ち上げ
docker compose up [-d]

# BE 疎通確認
curl http://localhost:8080/ping
# {"message":"pong"} が出力される

# FE 疎通確認
curl http://localhost:3000
# もしくは 上記URLにブラウザでアクセス
```
