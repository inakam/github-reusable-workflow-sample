# GitHub Actions モノレポ管理検証サンプル（Reusable Workflows版）

このリポジトリは、**Reusable Workflows + サブディレクトリ構造** を使って
モノレポでGitHub Actionsをサブディレクトリごとに管理し、**並列実行**を可能にする
方法を検証するためのサンプルです。

## 🎯 解決した課題

- ✅ frontend/backend の中にワークフローファイルを配置
- ✅ 並列実行の実現（Composite Actionsの制限を克服）
- ✅ 各チームが独立してCI処理を管理

## プロジェクト構成

```
.
├── README.md
├── frontend/          # Node.js プロジェクト
│   ├── package.json
│   ├── src/ & test/
│   └── .github/workflows/
│       └── ci.yml     # ★ フロントエンドのCI処理（並列実行対応）
├── backend/           # Go プロジェクト
│   ├── go.mod & main.go
│   └── .github/workflows/
│       └── ci.yml     # ★ バックエンドのCI処理（並列実行対応）
└── .github/workflows/
    ├── frontend-ci.yml     # 最小限のトリガー + 相対パス呼び出し
    └── backend-ci.yml      # 最小限のトリガー + 相対パス呼び出し
```

## 🚀 GitHub Actions の仕組み（Reusable Workflows）

### 1. サブディレクトリに完全なワークフロー定義

**frontend/.github/workflows/ci.yml**（抜粋）:
```yaml
on:
  workflow_call:  # 外部から呼び出し可能

jobs:
  lint:     # 並列実行
    runs-on: ubuntu-latest
    steps: [リンティング処理]
  
  test:     # 並列実行
    runs-on: ubuntu-latest  
    steps: [テスト処理]
  
  build:    # lint と test の完了後に実行
    needs: [lint, test]
    steps: [ビルド処理]
```

### 2. メインワークフローからの相対パス呼び出し

**frontend-ci.yml**（抜粋）:
```yaml
jobs:
  call-frontend-ci:
    uses: ./frontend/.github/workflows/ci.yml  # ★相対パス呼び出し
```

## 🎉 メリット

- **真の並列実行**: lint, test, security チェックが同時実行
- **完全な分離**: 各チームは自分のディレクトリ内のワークフローのみを管理
- **柔軟性**: 複雑な依存関係やジョブ構成を自由に設計可能
- **効率性**: .github/workflows は最小限の呼び出しコードのみ

## 検証方法

1. frontend/ 内のファイルを変更 → frontend の並列CI実行
2. backend/ 内のファイルを変更 → backend の並列CI実行
3. 両方変更 → 各々の並列CIが独立実行