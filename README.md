# GitHub Actions モノレポ管理検証サンプル（Composite Actions版）

このリポジトリは、**Composite Actions** を使ってモノレポでGitHub Actionsを
サブディレクトリごとに管理する方法を検証するためのサンプルです。

## プロジェクト構成

```
.
├── README.md
├── frontend/          # Node.js プロジェクト
│   ├── package.json
│   ├── src/
│   ├── test/
│   └── action.yml     # フロントエンドのCI処理を定義
├── backend/           # Go プロジェクト
│   ├── go.mod
│   ├── main.go
│   ├── main_test.go
│   └── action.yml     # バックエンドのCI処理を定義
└── .github/
    └── workflows/
        ├── frontend-ci.yml  # 最小限のトリガー設定のみ
        └── backend-ci.yml   # 最小限のトリガー設定のみ
```

## GitHub Actions の仕組み（Composite Actions）

### 1. 各サブディレクトリに action.yml を配置

実際のCI処理ロジック（テスト、ビルド、デプロイなど）は各サブディレクトリの
`action.yml` に記述します。

### 2. .github/workflows には最小限のトリガーのみ

メインの `.github/workflows/*.yml` には、pathsフィルターと
Composite Actionの呼び出しのみを記述します。

### 3. Composite Action の呼び出し

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./frontend  # frontend/action.yml を呼び出し
```

## メリット

- **関心事の分離**: 各チームは自分のディレクトリ内のaction.ymlのみを管理
- **処理の詳細をサブディレクトリに配置**: .github/workflows を汚さない
- **再利用性**: 同じComposite Actionを異なるワークフローから呼び出し可能

## 検証方法

1. frontend/ 内のファイルを変更 → frontend-ci.yml のみ実行
2. backend/ 内のファイルを変更 → backend-ci.yml のみ実行
3. 各ディレクトリのaction.ymlが実際のCI処理を実行