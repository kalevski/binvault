# BinVault

**BinVault** is an open-source file storage and compression service designed for cloud-native environments. It efficiently compresses, stores, and serves images and other files, making it ideal for developers, hosting providers, and edge computing platforms.

[![forthebadge](https://forthebadge.com/images/featured/featured-built-with-love.svg)](https://forthebadge.com)


![GitHub release (latest by date)](https://img.shields.io/github/v/release/kalevski/binvault?style=for-the-badge)
![Docker Pulls](https://img.shields.io/docker/pulls/kalevski/binvault?style=for-the-badge)
![GitHub issues](https://img.shields.io/github/issues/kalevski/binvault?style=for-the-badge)
![GitHub](https://img.shields.io/github/license/kalevski/binvault?style=for-the-badge)

## 🚀 Features

- **Efficient Storage**: Organizes files in "buckets" with customizable permissions.
- **Built-in Nginx Server**: Serves compressed files with caching and gzip support.
- **REST API**: Fully functional API for uploading, retrieving, and managing files.
- **Worker Queue**: Asynchronous compression using worker pools.
- **SQLite Database**: Tracks files and metadata efficiently.
- **Cloud-Native Ready**: Supports Kubernetes, Docker, and containerized deployments.

---

## 📦 Installation
```sh
docker run -p 80:80 -p 8080:8080 -v $(pwd)/data:/app/data ghcr.io/kalevski/binvault:latest
```

---

## 📌 Usage

🚀 **Start using BinVault today for efficient and optimized file storage!**

| HTTP Method | Endpoint                          | Request Body | Response Body |
|-------------|-----------------------------------|--------------|---------------|
| GET         | /api/buckets                      | None         | JSON array of buckets |
| GET         | /api/buckets/:name                | None         | JSON object of the bucket |
| POST        | /api/buckets                      | JSON object  | JSON object of the created bucket |
| DELETE      | /api/buckets/:name                | None         | JSON object with deletion status |
| GET         | /api/buckets/:name/files          | None         | JSON array of files |
| GET         | /api/buckets/:name/files/:id      | None         | JSON object of the file |
| GET         | /api/buckets/:name/files/:id/content | None         | File content |
| POST        | /api/buckets/:name/files          | File data    | JSON object of the uploaded file |
| DELETE      | /api/buckets/:name/files/:id      | None         | JSON object with deletion status |

## 🛠 Configuration
BinVault uses environment variables for customization:
```sh
SERVER_PORT=8080
SERVER_HOST=localhost
DATA_PATH=./data
DB_NAME=_database
JWKS=
```

---

## 🤝 Contributing
We welcome contributions! Follow these steps:
1. Fork the repository
2. Create a feature branch (`git checkout -b feature-name`)
3. Commit changes (`git commit -m "Added new feature"`)
4. Push to your fork and submit a PR

For detailed contribution guidelines, see [CONTRIBUTING.md](CONTRIBUTING.md).

---
## 📜 License

BinVault is released under the **Apache 2.0 License**.

