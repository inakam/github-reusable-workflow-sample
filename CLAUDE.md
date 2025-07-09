# GitHub Actions Composite Actions プロジェクト固有ルール

## プロジェクト概要
このプロジェクトは、モノレポでComposite Actionsを使用してGitHub Actionsをサブディレクトリごとに管理する検証サンプルです。

## アーキテクチャパターン
- **Composite Actions**: 各サブディレクトリ（frontend/, backend/）にaction.ymlを配置
- **最小限のトリガー**: .github/workflows/には、pathsフィルターとComposite Action呼び出しのみ記述
- **関心事の分離**: 各チームは自分のディレクトリ内のaction.ymlのみを管理

## ディレクトリ構成
```
.
├── frontend/
│   ├── action.yml          # フロントエンドのCI処理定義
│   ├── package.json
│   ├── src/
│   └── test/
├── backend/
│   ├── action.yml          # バックエンドのCI処理定義
│   ├── go.mod
│   ├── main.go
│   └── main_test.go
└── .github/workflows/
    ├── frontend-ci.yml     # 最小限のトリガー設定
    └── backend-ci.yml      # 最小限のトリガー設定
```

## 重要なルール

### 1. Composite Action の定義
- 各サブディレクトリにaction.ymlを配置
- `runs.using: 'composite'` を指定
- すべてのstepで `shell: bash` を明示的に指定

### 2. ワークフローからの呼び出し
- `uses: ./frontend` または `uses: ./backend` で呼び出し
- pathsフィルターで適切なトリガー制御を実装

### 3. 作業ディレクトリの管理
- Composite Action内のstepで `working-directory` を明示的に指定
- ルートディレクトリからの相対パスを使用

## テスト方法
```bash
# Frontend
cd frontend && npm test

# Backend  
cd backend && go test -v ./...
```

## メリット
- 各チームが独立してCI処理を管理できる
- .github/workflowsディレクトリが簡潔に保たれる
- 同一のComposite Actionを複数のワークフローから再利用可能
- サブディレクトリ内でCI処理の詳細を完結できる