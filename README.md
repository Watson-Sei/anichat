# anichat
リアタイアニメチャット

![5c9544db-c750-42fc-a4c1-d1d74cef769f](https://user-images.githubusercontent.com/55475145/121808068-e76df780-cc91-11eb-8d2e-07c79dbfa222.png)

![6abb244f-e56c-48bb-9d9d-d8a4006f4615](https://user-images.githubusercontent.com/55475145/121808090-f81e6d80-cc91-11eb-85fc-4992c3510527.png)


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

## セットアップ Production(本番環境)

### フロントエンド
```shell
// build
$ yarn build
// system up -d 
$ docker-compose -f docker-compose.production.yml up -d --build
```
