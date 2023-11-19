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
go clean -testcache
go test ./...
```

Reference links:
https://gobyexample.com/channels 


Packages:
github.com/jinzhu/gorm - ORM
cloud.google.com/go/storage - Google Cloud Storage
github.com/joho/godotenv - Reads a .env file and add to the Environments Variables
github.com/streadway/amqp - RabbitMQ package

Requirement:
Google Cloud Platform
-> IAM & Admin
  -> Service Account
    -> Create a Service Account
      -> Create a json Key on the Service Account
      -> Download the key
  -> IAM
    -> Grant Access
      -> Add principals - Set your Service Account name
      -> Assign roles - Set Environment and Storage Object Administrator


On your computer:
-> Add the json key file the root folder of the project
-> Rename it to bucket-credential.json
