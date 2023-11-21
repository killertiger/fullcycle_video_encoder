# fullcycle_video_encoder
Encoder de vídeo da aplicação de micro vídeos da fullcycle

Executing the app
```
docker compose up -d
```

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


Setting up:
RabbitMQ:
http://localhost:15672/

On RabbitMQ

Create an exchange:
Name: dlx
Type: fanout

Create two queues:
Type: Classic

1. Name: videos-result
2. Name: videos-failed

After creating the queue, we should do the binding.
1. Click on the queue: videos-result
- Binding
  - From exchange: amq.direct
  - Routing key: jobs

2. Click on the queue: videos-failed
- Binding
  - From exchange: dlx

Executing the server
```
go run framework/cmd/server/server.go
```

Testing:
Access RabbitMQ, queue videos: http://localhost:15672/#/queues/%2F/videos

Publish a message, example:
```
{
"resource_id": "id-cliente1",
"file_path": "video_example.mp4"
}
```

Validating:
1. On the GCP - Cloud storage, you should be able to see a new folder with the video
2. On the RabbitMQ - videos-result, it should exist a new message on this queue