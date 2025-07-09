# GitHub Actions Reusable Workflows プロジェクト固有ルール

## プロジェクト概要
このプロジェクトは、モノレポでReusable Workflows + サブディレクトリ構造を使用して
GitHub Actionsをサブディレクトリごとに管理し、並列実行を可能にする検証サンプルです。

## アーキテクチャパターン
- **Reusable Workflows**: 各サブディレクトリの.github/workflows/に完全なワークフロー定義
- **相対パス呼び出し**: ルートの.github/workflows/から相対パスで呼び出し
- **並列実行対応**: Composite Actionsの制限を克服し、真の並列実行を実現

## ディレクトリ構成
```
.
├── frontend/
│   ├── package.json, src/, test/
│   └── .github/workflows/
│       └── ci.yml          # フロントエンドの完全なワークフロー定義
├── backend/
│   ├── go.mod, main.go, main_test.go
│   └── .github/workflows/
│       └── ci.yml          # バックエンドの完全なワークフロー定義
└── .github/workflows/
    ├── frontend-ci.yml     # 最小限のトリガー + 相対パス呼び出し
    └── backend-ci.yml      # 最小限のトリガー + 相対パス呼び出し
```

## 重要なルール

### 1. Reusable Workflow の定義
- 各サブディレクトリの.github/workflows/ci.ymlに配置
- `on: workflow_call:` を指定して外部呼び出しを可能にする
- 複数ジョブで並列実行を設計（lint, test, security など）

### 2. 相対パス呼び出し
- `uses: ./frontend/.github/workflows/ci.yml` で呼び出し
- pathsフィルターで適切なトリガー制御を実装

### 3. 並列実行の設計
- 独立したジョブ（lint, test, security）を並列実行
- `needs: [lint, test]` で依存関係を制御

### 4. 作業ディレクトリの管理
- 各ジョブで `defaults.run.working-directory` を指定
- ルートディレクトリからの相対パスを使用

## 並列実行の例
```yaml
jobs:
  lint:    # 並列実行
    runs-on: ubuntu-latest
  test:    # 並列実行  
    runs-on: ubuntu-latest
  build:   # lint と test の完了後
    needs: [lint, test]
    runs-on: ubuntu-latest
```

## テスト方法
```bash
# Frontend
cd frontend && npm test

# Backend  
cd backend && go test -v ./...
```

## メリット
- 真の並列実行が可能（Composite Actionsの制限なし）
- 各チームが独立してワークフローを管理できる
- 複雑な依存関係やジョブ構成を自由に設計可能
- サブディレクトリ内でワークフロー定義が完結
- .github/workflowsディレクトリが簡潔に保たれる