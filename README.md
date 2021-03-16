# anichat
リアタイアニメチャット

## 技術選定

- fiber - golang
- nuxt.js - javascript
- nginx
- docker
- ubuntu
- firebase
- mysql

## セットアップ Dev（開発環境）

### データベース
```text
# /docker/db/.env

MYSQL_ROOT_PASSWORD=root
MYSQL_USER=docker_user
MYSQL_PASSWORD=docker_pass
MYSQL_DATABASE=chat-db
```

### フロントエンド
```text
# /docker/front/config/.env.development

API_URL="http://api:8080"
API_URL_BROWSER="http://localhost:8080"
BASE_URL=
apiKey=
authDomain=
databaseURL=
projectId=
storageBucket=
messagingSenderId=
appId=
measurementId=
```

### バックエンド
> /api_v1/firebase_admin_sdk.json　Google Authに必要

### Docker-Compose実行
```shell
// main
$ docker-compose -f docker-compose.development.yml up -d --build
// system memory max
$ docker system prune
// build image cache
$ docker-compose -f docker-compose.development.yml build --no-cache
```
