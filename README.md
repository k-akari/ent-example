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

## 2. CI/CD
### 2-1. CI
`develop` or `main`ブランチをマージ先としたPRの作成をトリガーとしてGitHub Actionsで静的解析とテストを実行する。

詳しくは`.github/workflows/ci.yml`を参照

### 2-2. CD
`main`ブランチの更新イベントをトリガーとしてCodeBuildが走り、イメージをビルドしてECRへプッシュする。

詳しくは`buildspec.yml`を参照
