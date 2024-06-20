# good-growth-interview

1. build app usng dockerfile with:

```
docker build -t good-growth-api:1.0 .
```

2. run in docker image

```
docker run -d -p 5000:5000 -e "mongodb://mongoadmin:secret@db:27017/" good-growth-api:1.0
```

3. to test running image

```
curl -i localhost:5000/messages
```

```
docker run --name messagerepo -p 27017:27017 -d mongo
```
