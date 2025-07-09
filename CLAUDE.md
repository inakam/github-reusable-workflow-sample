# GitHub Actions 複数Composite Actions プロジェクト固有ルール

## プロジェクト概要
このプロジェクトは、モノレポで複数のComposite Actionsを使用して
GitHub Actionsをサブディレクトリごとに管理し、並列実行を可能にする検証サンプルです。

## アーキテクチャパターン
- **複数Composite Actions**: 各サブディレクトリのactions/に機能別のComposite Action定義
- **並列ジョブ呼び出し**: ルートの.github/workflows/で複数ジョブから並列呼び出し
- **機能分割**: lint, test, security, build が独立したComposite Actions

## ディレクトリ構成
```
.
├── frontend/
│   ├── package.json, src/, test/
│   └── actions/
│       ├── lint/action.yml     # リンティング処理
│       ├── test/action.yml     # テスト処理
│       └── build/action.yml    # ビルド処理
├── backend/
│   ├── go.mod, main.go, main_test.go
│   └── actions/
│       ├── lint/action.yml     # 静的解析処理
│       ├── test/action.yml     # テスト処理
│       ├── security/action.yml # セキュリティスキャン
│       └── build/action.yml    # ビルド処理
└── .github/workflows/
    ├── frontend-ci.yml         # 複数ジョブで並列呼び出し
    └── backend-ci.yml          # 複数ジョブで並列呼び出し
```

## 重要なルール

### 1. Composite Action の定義
- 各サブディレクトリのactions/<機能名>/action.ymlに配置
- `runs.using: 'composite'` を指定
- すべてのstepで `shell: bash` を明示的に指定
- 機能ごと（lint, test, build等）に分割

### 2. 並列実行の実現
- メインワークフローで複数のジョブを定義
- 各ジョブが独立したComposite Actionを呼び出し
- `needs` で依存関係を制御

### 3. 作業ディレクトリの管理
- Composite Action内のstepで `working-directory` を明示的に指定
- ルートディレクトリからの相対パスを使用

### 4. pathsフィルター
- pathsフィルターで適切なトリガー制御を実装
- サブディレクトリの変更のみで対応するワークフローを実行

## 並列実行の例
```yaml
jobs:
  lint:           # 並列実行
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./frontend/actions/lint
      
  test:           # 並列実行
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./frontend/actions/test
      
  build:          # lint と test の完了後
    needs: [lint, test]
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./frontend/actions/build
```

## テスト方法
```bash
# Frontend
cd frontend && npm test

# Backend  
cd backend && go test -v ./...
```

## メリット
- 真の並列実行が可能（複数ジョブで独立実行）
- 各チームが独立してactions/内のComposite Actionsを管理できる
- 機能別に分割されており理解しやすい
- サブディレクトリ内でCI処理定義が完結
- CI処理の論理構造と実行構造が一致
- .github/workflowsディレクトリが簡潔に保たれる