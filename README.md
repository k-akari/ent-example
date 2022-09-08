# ent-example
ORMとしてentを利用した簡易アプリ

## 1. 開発環境
### 1-1. 環境構築
1. イメージをビルドし、コンテナをデーモン起動する。
```bash
docker compose build
docker compose up -d
```
2. コンテナ内でリモート開発
VS Code左下の緑色のアイコンをクリックし、Reopen in Containerをクリックする。

### 1-2. アプリの起動
1. アプリを起動する。
```bash
go run cmd/app/main.go
```

### 1-3. 本番イメージの確認
1. 以下のように`Dockerfile.prod`の`FROM`文から`public.ecr.aws/docker/library/`を除去する。
```Dockerfile
FROM public.ecr.aws/docker/library/golang:1.18-alpine as builder
↓↓↓
FROM golang:1.18-alpine as builder
```
```Dockerfile
FROM public.ecr.aws/docker/library/alpine:latest
↓↓↓
FROM alpine:latest
```

2. イメージをビルドする。
```bash
docker build --no-cache -f build/app/Dockerfile.prod -t ent-example:latest .
```

3. コンテナを起動する。
```bash
docker run -it -p 8080:8080 --rm ent-example:latest
```

## 2. CI/CD
### 2-1. CI
`develop` or `main`ブランチをマージ先としたPRの作成をトリガーとしてGitHub Actionsで静的解析とテストを実行する。

詳しくは`.github/workflows/ci.yml`を参照

### 2-2. CD
`main`ブランチの更新イベントをトリガーとしてCodeBuildが走り、イメージをビルドしてECRへプッシュする。

詳しくは`buildspec.yml`を参照。
