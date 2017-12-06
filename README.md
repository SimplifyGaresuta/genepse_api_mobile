# genepse_api_mobile
genpseのモバイル用API

# Installation

```
$ go get github.com/SimplifyGaresuta/genepse_api_mobile
```

# Usage

[APIドキュメント](https://docs.google.com/spreadsheets/d/1-q8nE-WqRuiR_29qE8KX-_7tY1fdgvvmYsHC2DYsNOg/edit#gid=0)を御覧ください。

# How to deploy to Google App Engine

gcloudアカウントの認証を行います。

```
$ gcloud auth login
```

プロジェクトidをローカルに設定します。

```
$ gcloud config set project [YOUR_PROJECT_ID]
```

app.yamlがあるディレクトリに移動し、デプロイします。

```
$ cd [THIS_PROJECT]/app
$ gcloud app deploy
```

# How to Connect to CloudSQL

Google Cloud Platform Console で、右上隅にある Cloud Shell アイコンをクリックします。

```
gcloud beta sql connect genepse --user=root
```

# How to test request

### ユーザー登録

```
$ curl -H 'Content-Type:application/json' -H 'User-Agent:iPhone' http://localhost:8080/users -d @samples/requests/user_create.json
```

### 複数ユーザー取得

```
$ curl -H 'User-Agent:iPhone' http://localhost:8080/users?per_page=20&page=1
```

### ユーザー更新

```
curl -X PATCH -H "Content-Type: application/json" http://localhost:8080/v1/users/1 -d @samples/requests/user_update.json
```
