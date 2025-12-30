# hello-world

This repository contains a Go HTTP server with configurable response.
The CI builds a Docker image and pushes it to Docker Hub at https://hub.docker.com/repository/docker/sepetrov/hello-world.

## Usage

Pull and run the container from Docker Hub:

```bash
docker run --rm -p 8080:8080 sepetrov/hello-world:latest
2025/12/30 09:34:00 default content type: text/plain
2025/12/30 09:34:00 default status code: 200
2025/12/30 09:34:00 default response body: Hello World!
2025/12/30 09:34:00 start listening on port 8080
```

Make a request to the running container. This will return the default response: `Hello World!`.

```bash
curl -i http://localhost:8080                                                                                        
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Tue, 30 Dec 2025 09:34:46 GMT
Content-Length: 12

Hello World! 
```
