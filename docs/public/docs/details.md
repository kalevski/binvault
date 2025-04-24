# How it Works

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