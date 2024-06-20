# good-growth-interview

This web service provides a simple interface for users to store and retrieve plain text messages. Users can send a message to the service, which returns a unique ID for that message. This ID can then be used to retrieve the message. The service is built in Go and uses MongoDB, a high-performance database, for storing messages. The service is containerized using Docker, making it easy to deploy and scale in various environments.

1. build app binary with:

```bash
make build
```

2. to run the binary in docker image:

```bash
make up
```

3. to shutdown and delete the containers

```bash
make down
```
