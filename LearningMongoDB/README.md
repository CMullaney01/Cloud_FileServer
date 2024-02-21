# learning mongodb and how to connect to mmy go application

- Should I use ATLAS?

- Note for learning Go and backend development i really like this guy [here](https://www.youtube.com/watch?v=-gW7oSFxT2I&list=PL0xRBLFXXsP7-0IVCmoo2FEWBrQzfH2l8&index=1)

# Tutorial Building a cat fact JSON API

[Tutorial 1](https://www.youtube.com/watch?v=iak56rgR05A)

### Go dependencies
```
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson
```

### GOLANG QUICKSTART
```
client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27018)) 
    if err != nil {
        panic(err)
    }
```
1. Start mongo db (You have already set this up in the docker compose)
