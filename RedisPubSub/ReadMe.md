# Pub Sub sử dụng Redis

[Bài viết gốc ở đây](https://itnext.io/redis-as-a-pub-sub-engine-in-go-10eb5e6699cc)

Trong bài này chúng ta thí nghiệm với Redis Server làm trung gian, có một ứng dụng publish và một ứng dụng subscribe

Để cài đặt redis hãy dùng docker tạo một container lắng nghe ở cổng `localhost:6379`
```
docker run --name=redis -p 6379:6379 -d -e REDIS_PASSWORD=123 -v $PWD/data:/data redis:alpine3.14 /bin/sh -c 'redis-server --appendonly yes --requirepass ${REDIS_PASSWORD}'
```

Mở một terminal
```
$ cd subscribe
$ go run main.go
```

Mở thêm terminal nữa

```
$ cd publish
$ go run main.go
```

Bạn sẽ thấy terminal ở subscribe
```
User: Adam Coder registered with Email: AdamCoder@Gmail.com
User: Edward Redisson registered with Email: EdwardRedisson@Gmail.com
User: Amber Bolmer registered with Email: AmberBolmer@Redis.com
User: Redis Fish registered with Email: RedisFish@Gmail.com
User: Redis Coder registered with Email: RedisCoder@Redis.com
User: Angelica Edisson registered with Email: AngelicaEdisson@Gmail.com
User: Amber Tesla registered with Email: AmberTesla@Gmail.com
User: Watch Coder registered with Email: WatchCoder@Redis.com
User: Edward Edisson registered with Email: EdwardEdisson@Awesomeness.com
User: Redis Fish registered with Email: RedisFish@Redis.com
User: Jennifer Redisson registered with Email: JenniferRediss
```

