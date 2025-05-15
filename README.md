# 📋 Go Todo App

Go言語で作成した、ローカルで動作するシンプルなWebベースのTodoアプリです。
緊切日、タグ、検索機能を備えており、JSONファイルにデータを保存します。

---

## 🚀 主な機能

* ✅ タスクの追加・完了チェック
* 🗑 タスクの削除
* 🗕 緊切日の入力と表示
* 🏷 タグによる分類
* 🔍 検索バー（タイトル・タグ検索に対応）
* 📀 JSONファイルへの保存と読み込み
* 🌐 シンプルなHTMLテンプレート＋CSSデザイン

---

## 🛠 環境構築と実行手順

### 1. Goのインストール（未インストールの場合）

公式サイトからGoをインストール：

🔗 [https://go.dev/dl/](https://go.dev/dl/)

インストール確認：

```bash
go version
```

---

### 2. リポジトリのクローンと移動

```bash
git clone https://github.com/yourname/go-todo-app.git
cd go-todo-app
```

---

### 3. ディレクトリ構成

```
go-todo-app/
├── main.go              # メインロジック
├── todos.json           # タスク保存用ファイル
├── templates/
│   └── index.html       # HTMLテンプレート
├── static/
│   └── style.css        # CSSファイル
```

---

### 4. アプリの起動

```bash
go run main.go
```

その後、ブラウザで下記にアクセス：

```
http://localhost:8080
```

---

## ✏️ アプリの使い方

1. 入力フォームで「タスク名」「緊切日」「タグ」を記入
2. 「追加」ボタンでタスクを登録
3. タスクを完了したらチェックボックスをON
4. 「削除」でタスクを削除
5. 検索バーでタイトルやタグによる篩り込み

---

## 📀 データの保存形式

アプリ内で扱うtodos.jsonファイルの構造は以下の通り：

```json
[
  {
    "Title": "宿題をやる",
    "Done": false,
    "Deadline": "2025-05-18",
    "Tag": "学校"
  },
  {
    "Title": "牛乳を買う",
    "Done": true,
    "Deadline": "2025-05-16",
    "Tag": "買い物"
  }
]
```

---

## 🛆 使用技術

* 言語: Go
* テンプレート: html/template
* 保存形式: JSON
* フロントエンド: HTML / CSS（Tailwind風の自作CSS）

---


## 💡 よくあるエラーと解決法

| エラー                         | 対処法                                     |
| --------------------------- | --------------------------------------- |
| go: command not found       | Goがインストールされていない。go.dev からインストールしてください。  |
| imported and not used       | `import`文で読み込んだパッケージが未使用のとき発生。使う or 消す。 |
| invalid character U+3000    | ソースコード内に全角スペース（U+3000）が含まれている。半角に修正。    |
| undefined: stringsTrimSpace | 正しくは strings.TrimSpace。スペルミスに注意。        |
