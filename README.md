# chat_go

chat app

## Docker 関連

Backend 側の log を見る

```
$ docker-compose logs -f
```

Container の中に入る

```
docker exec -it chat_go-go_chat-1 sh
```

## Backend のディレクトリ構成

```
.
├── Dockerfile
├── bin
│   └── wait.sh
├── controllers
│   ├── router.go
│   └── users.go
├── main.go
├── model
│   └── db.go
└── tmp
    ├── build-errors.log
    └── main
```

- router ディレクトリは api リクエストを書きます。
- model ディレクトリは DB 操作し、router によって呼びだれれる。db.go を置く
