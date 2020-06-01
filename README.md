## nginx examle (loadbalancer)

```bash
$ docker-compose up
```

## バランシングされていることの確認
- リクエストするたびに、メッセージが交互に入れ替わるはず
```bash
$ curl localhost/hello
```

### LBメモ
- `nginx.conf`の記載
```
http {
    upstream backend {
        server app:8080;
        server app2:8080;
    }
}
```

- 上記記載で`backend`という名前で`app:8080`と`app2:8080`の２つにバランシングする
- デフォルトはラウンドロビンでバランシングするらしい
