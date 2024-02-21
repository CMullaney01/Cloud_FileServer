# learning mongodb and how to connect to mmy go application

- Should I use ATLAS?

- Note for learning Go and backend development i really like this guy [here](https://www.youtube.com/watch?v=-gW7oSFxT2I&list=PL0xRBLFXXsP7-0IVCmoo2FEWBrQzfH2l8&index=1)

# Tutorial 1: Building a cat fact JSON API with mongoDB storage

[Tutorial 1: MongoDB](https://www.youtube.com/watch?v=iak56rgR05A)

# Tutorial 2:

[Tutorial 2: Amazon S3 Absolute Basics](https://www.youtube.com/watch?v=FLIp6BLtwjk)

- note not in video: I am generating an IAM user with a access key and secret key. This is to allow my backend to sign urls to provide my frontend access to files. However a cool solution in the future if this ever becomes production would be to use IAM roles which creates temporary access and this can be used if my backend is run on a EC2 instance.

# Usage
- If you would like to test the backend returning required URLS you will need to create your own s3 instance and set the environment variables
- after this, you need to run the docker-compose in the docker compose file to create the mongodb database. There is a .bat file ready for you
- after this you are free to run ```go run .``` and you can make the appropriate requests in the thunderclient