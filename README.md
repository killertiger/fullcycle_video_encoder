# fullcycle_video_encoder
Encoder de vídeo da aplicação de micro vídeos da fullcycle

Connect to the container
```
docker compose exec app bash
```

Update go.mod references
```
go mod tidy
```

Run tests
```
go test ./...
```


Packages:
github.com/jinzhu/gorm - ORM
cloud.google.com/go/storage - Google Cloud Storage