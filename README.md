# GitHub Actions モノレポ管理検証サンプル（複数Composite Actions版）

このリポジトリは、**複数のComposite Actions** を使って
モノレポでGitHub Actionsをサブディレクトリごとに管理し、**並列実行**を可能にする
方法を検証するためのサンプルです。

## 🎯 解決した課題

- ✅ frontend/backend の中にワークフローファイル（Composite Actions）を配置
- ✅ 並列実行の実現（複数ジョブでComposite Actionsを並列呼び出し）
- ✅ 各チームが独立してCI処理を管理

## プロジェクト構成

```
.
├── README.md
├── frontend/          # Node.js プロジェクト
│   ├── package.json, src/, test/
│   └── actions/       # ★ フロントエンドのCI処理定義
│       ├── lint/action.yml
│       ├── test/action.yml
│       └── build/action.yml
├── backend/           # Go プロジェクト
│   ├── go.mod, main.go
│   └── actions/       # ★ バックエンドのCI処理定義
│       ├── lint/action.yml
│       ├── test/action.yml
│       ├── security/action.yml
│       └── build/action.yml
└── .github/workflows/
    ├── frontend-ci.yml     # 複数ジョブで並列呼び出し
    └── backend-ci.yml      # 複数ジョブで並列呼び出し
```

## 🚀 GitHub Actions の仕組み（複数Composite Actions）

### 1. サブディレクトリに機能別Composite Actions

**frontend/actions/lint/action.yml**（抜粋）:
```yaml
name: 'Frontend Lint Action'
runs:
  using: 'composite'
  steps:
    - name: Setup Node.js
    - name: Install dependencies
    - name: Run linting
```

### 2. メインワークフローで並列呼び出し

**frontend-ci.yml**（抜粋）:
```yaml
jobs:
  lint:     # 並列実行
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./frontend/actions/lint
  
  test:     # 並列実行
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./frontend/actions/test
  
  build:    # lint と test の完了後
    needs: [lint, test]
    steps:
      - uses: ./frontend/actions/build
```

## 🎉 メリット

- **真の並列実行**: 複数ジョブで各Composite Actionを同時実行
- **完全な分離**: 各チームは自分のディレクトリ内のactionsのみを管理  
- **機能分割**: lint, test, build が独立したComposite Actions
- **効率性**: CI処理の論理構造と実行構造が一致

## 検証方法

1. frontend/ 内のファイルを変更 → lint と test が並列実行 → build実行
2. backend/ 内のファイルを変更 → lint, test, security が並列実行 → build実行
3. 各チームが自分のactions/内のみを管理