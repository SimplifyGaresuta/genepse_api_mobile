# genepse_api_mobile
genpseのモバイル用API

# Installation

```
$ go get github.com/SimplifyGaresuta/genepse_api_mobile
```

# Usage

[APIドキュメント](https://docs.google.com/spreadsheets/d/1-q8nE-WqRuiR_29qE8KX-_7tY1fdgvvmYsHC2DYsNOg/edit#gid=0)を御覧ください。

# How to test request

### ユーザー登録

```
$ curl -H 'Content-Type:application/json' -H 'User-Agent:iPhone' http://localhost:8080/users -d @samples/request/user_create.json
```
