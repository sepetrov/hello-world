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

## Configuration

The response can be customized using environment variables:
- `SERVER_PORT`: sets the server listening port (default: `8080`)
- `CONTENT_TYPE`: sets the `Content-Type` header (default: `text/plain`)
- `STATUS_CODE`: sets the HTTP status code (default: `200`)
- `RESPONSE_BODY`: sets the response body (default: `Hello World!`)

```bash
docker run --rm -p 8080:8088 -e SERVER_PORT=8088 -e CONTENT_TYPE=application/json -e STATUS_CODE=201 -e RESPONSE_BODY='{"status":"ok"}' sepetrov/hello-world:latest
2025/12/30 09:48:08 default content type: application/json
2025/12/30 09:48:08 default status code: 201
2025/12/30 09:48:08 default response body: {"status":"ok"}
2025/12/30 09:48:08 start listening on port 8088
```

```bash
curl -i http://localhost:8080                                                                                        
HTTP/1.1 201 Created
Content-Type: application/json
Date: Tue, 30 Dec 2025 09:48:32 GMT
Content-Length: 15

{"status":"ok"}
```

Alternatively, the response can be customised by the caller using query parameters:
- `content_type`: sets the `Content-Type` header
- `status_code`: sets the HTTP status code
- `response_body`: sets the response body

```bash
curl -i 'http://localhost:8080?status_code=404&response_body=%7B%22status%22%3A%22not%20found%22%7D'                        
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Tue, 30 Dec 2025 09:49:57 GMT
Content-Length: 22

{"status":"not found"}
```

```bash
curl -i http://localhost:8080 -d 'status_code=404&response_body={"status":"not found"}'             
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Tue, 30 Dec 2025 09:50:18 GMT
Content-Length: 22

{"status":"not found"}
```
