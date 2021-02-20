# Event Driven Microservices with Golang

This an example of event driven microservices using GCP PubSub & Golang using Gin microframework, Cobra command manager & Gorm (ORM). 

### What are microservices? 
Microservices are a style of architecture and also a form on how to software. With microservices the applications get divided into their smallest and independent pieces.

### Why event driven microservices?
To test trade offs between non event driven and event driven microservices and benchmark other implementations such us SQS+SNS and Apache Kafka.

### How to start this project

#### Dependencies:

- Docker (to run the project with ```docker-compose```)
- PostgreSQL 
- MongoDB
- GCP service account key with PubSub Admin permissions, or fine grained ones for topics and subscriptions. Check out ```./pkg/topics``` & ```./pkg/subscriptions```

#### Migrate DB

Run the following command or copy paste the SQL script inside ```./services/catalog/migrations```

```` go run ./services/catalog db:migrate  ````

#### Starting the project

To run the PubSubs locally we can run the emulator (Check the make file for the commands) or we can create the Topics:
- ```dead-letters``` 
- ```monitoring```

Followed by the Subscriptions:
- ```dead-letters-sub```
- ```monitoring-sub```

Run the command 

````docker-compose build && docker-compose up````

#### Check the API

Try to create a brand and then check the mongo DB to check if the monitoring service is working.

## Next Steps

- Add tests
- Build a SQS + SNS implementation
- Build Kafka implementation
