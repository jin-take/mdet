# 環境構築

## 前提

- Dockerがインストールされていること
- ngrokがインストールされていること

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

# ngrok起動
ngrok http 8080
# 「Forwarding」に記載されているURLを client/app/.env の API_URL に記述

# BE,FE 確認（URLにブラウザでアクセス）
http://localhost:3000
```
