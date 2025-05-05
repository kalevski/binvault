# How it Works

1. Exposes Rest API for uploading files
1. Exposes Rest API for downloading files
1. Uses background workers to process the files (based on configuration)
1. Saves the files to disk on desired location
1. Files can be served using any server

---

1. Group of users uploading files
1. RestAPI with Secured layer accepts the files
1. Box processing the files and sorting them in 2 boxes
1. One box is exposed publicly via Nginx
1. Another is served through API

[-- INFRASTRUCTURE DIAGRAM --]
[-- OVERVIEW --]

---

- Docker
- Docker Compose
- Kubernetes

## Docker

Running BinVault in a Docker container:​
```bash
docker run -d \
  -p 8080:8080 \
  -v /path/to/data:/data \
  kalevski/binvault
```
Replace /path/to/data with your desired host directory for persistent storage.