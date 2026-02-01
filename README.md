# hello-world

This repository contains a Go HTTP server with configurable response.
The CI builds a Docker image and pushes it to Docker Hub at https://hub.docker.com/repository/docker/sepetrov/hello-world.

## Features

- Configurable response
- Logging of requests
- Support for `robots.txt`

## Usage

Pull and run the container from Docker Hub:

```bash
docker run --rm -p 8080:8080 sepetrov/hello-world:latest
2026/02/01 08:36:02 config: default content type: text/plain
2026/02/01 08:36:02 config: default status code: 200
2026/02/01 08:36:02 config: default response body: Hello World!
2026/02/01 08:36:02 config: with robots.txt: false
2026/02/01 08:36:02 start listening on port 8080
10.0.2.100:44474 - curl/8.5.0 [01/Feb/2026:08:38:40 +0000] "GET / HTTP/1.1" 201 15
```

Make a request to the running container. This will return the default response: `Hello World!`.

```bash
curl -i http://localhost:8080
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Sun, 01 Feb 2026 08:38:40 GMT
Content-Length: 12

Hello World!
```

## Configuration

The response can be customized using environment variables:
- `SERVER_PORT`: sets the server listening port (default: `8080`)
- `CONTENT_TYPE`: sets the `Content-Type` header (default: `text/plain`)
- `STATUS_CODE`: sets the HTTP status code (default: `200`)
- `RESPONSE_BODY`: sets the response body (default: `Hello World!`)
- `WITH_ROBOTS_TXT`: if set to `true`, serves a `robots.txt` file disallowing all crawlers at the `/robots.txt` endpoint (default: `false`)

```bash
docker run --rm -p 8080:8088 -e SERVER_PORT=8088 -e CONTENT_TYPE=application/json -e STATUS_CODE=201 -e RESPONSE_BODY='{"status":"ok"}' -e WITH_ROBOTS_TXT=true sepetrov/hello-world:latest
2026/02/01 08:38:05 config: default content type: application/json
2026/02/01 08:38:05 config: default status code: 201
2026/02/01 08:38:05 config: default response body: {"status":"ok"}
2026/02/01 08:38:05 config: with robots.txt: true
2026/02/01 08:38:05 start listening on port 8088
10.0.2.100:44474 - curl/8.5.0 [01/Feb/2026:08:38:40 +0000] "GET / HTTP/1.1" 201 15
10.0.2.100:34934 - curl/8.5.0 [01/Feb/2026:08:38:50 +0000] "GET /robots.txt HTTP/1.1" 200 25
```

```bash
curl -i http://localhost:8080
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 01 Feb 2026 08:38:40 GMT
Content-Length: 15

{"status":"ok"}
```

```bash
curl -i http://localhost:8080/robots.txt
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Sun, 01 Feb 2026 08:38:50 GMT
Content-Length: 25

User-agent: *
Disallow: /
```

## Request-time Customisation

Alternatively, the response can be customised by the caller using query or POST parameters:
- `content_type`: sets the `Content-Type` header
- `status_code`: sets the HTTP status code
- `response_body`: sets the response body

```bash
curl -i 'http://localhost:8080?status_code=404&response_body=%7B%22status%22%3A%22not%20found%22%7D'                        
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Sun, 01 Feb 2026 08:38:40 GMT
Content-Length: 22

{"status":"not found"}
```

```bash
curl -i http://localhost:8080 -d 'status_code=404&response_body={"status":"not found"}'             
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Sun, 01 Feb 2026 08:38:40 GMT
Content-Length: 22

Not Found
```
