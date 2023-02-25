# 環境構築手順

1. リポジトリをクローンし、プロジェクトへ移動

```
$ git clone git@github.com:Tomoya185-miyawaki/attend-log-gin.git
$ cd attend-log-gin
```

2. initコマンドを実行

```
$ make init
```

- フロント（Vue）URL
  - http://localhost:8080
- API（Gin）URL
  - http://localhost:3000
- API（godoc）
  - http://localhost:9000/pkg/github.com/Tomoya185-miyawaki/attend-log-gin/

# コミットルール

```:emoji: #Issue番号 変更内容``` でコミットすること

## 絵文字Prefix一覧

|  type  |  emoji  |
| ---- | ---- |
|  初めてのコミット（Initial Commit）  |  🎉  |
|  バージョンタグ（Version Tag）  |  🔖  |
|  新機能（New Feature）  |  ✨  |
|  バグ修正（Bugfix）  |  🐛  |
|  リファクタリング(Refactoring)  |  ♻️  |
|  ドキュメント（Documentation）  |  📚  |
|  デザインUI/UX(Accessibility)  |  🎨  |
|  パフォーマンス（Performance）  |  🐎  |
|  ツール（Tooling）  |  🔧  |
|  テスト（Tests）  |  🚨  |
|  非推奨追加（Deprecation）  |  💩  |
|  削除（Removal）  |  🗑️  |
|  WIP(Work In Progress)  |  🚧  |

```
==================== Emojis ====================
🎉 :tada: 初めてのコミット（Initial Commit）
🔖 :bookmark: バージョンタグ（Version Tag）
✨ :sparkles: 新機能（New Feature）
🐛 :bug: バグ修正（Bugfix）
♻️  :recycle: リファクタリング(Refactoring)
📚 :books: ドキュメント（Documentation）
🎨 :art: デザインUI/UX(Accessibility)
🐎 :horse: パフォーマンス（Performance）
🔧 :wrench: ツール（Tooling）
🚨 :rotating_light: テスト（Tests）
💩 :hankey: 非推奨追加（Deprecation）
🗑️ :wastebasket: 削除（Removal）
🚧 :construction: WIP(Work In Progress)
```