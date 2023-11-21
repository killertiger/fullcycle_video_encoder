# Video encoder
Video encoder app in Golang for the Micro Video application from Full Cycle

## Set up

### Initial configuration

#### Google Cloud Platform

IAM & Admin
- Service Account
  - Create a Service Account
    - Create a json Key on the Service Account
    - Download the key
- IAM
  - Grant Access
    - Add principals - Set your Service Account name
    - Assign roles - Set Environment and Storage Object Administrator


On your computer:
- Add the json key file the root folder of the project
- Rename it to bucket-credential.json

#### Executing the container

Executing the app
```
docker compose up -d
```

#### Set up RabbitMQ

Access: http://localhost:15672/

Create an exchange:
Name: dlx
Type: fanout

Create two queues:
Type: Classic

1. Name: videos-result
2. Name: videos-failed

After creating the queues, we should do the binding.
1. Click on the queue: videos-result
- Binding
  - From exchange: amq.direct
  - Routing key: jobs

2. Click on the queue: videos-failed
- Binding
  - From exchange: dlx


## Running the app
Connect to the container
```
docker compose exec app bash
```

Execute the server
```
go run framework/cmd/server/server.go
```

## Development

### Commands

Update go.mod references
```
go mod tidy
```

Run tests
```
go clean -testcache
go test ./...
```

Executing with race condition checker:
```
go run -race framework/cmd/server/server.go
```

### References

Reference links:
- https://gobyexample.com/channels - How channels works on Golang


Packages:
github.com/jinzhu/gorm - ORM
cloud.google.com/go/storage - Google Cloud Storage
github.com/joho/godotenv - Reads a .env file and add to the Environments Variables
github.com/streadway/amqp - RabbitMQ package


### Testing

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