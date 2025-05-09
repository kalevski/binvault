# 📦 BinVault

**BinVault** is a lightweight, open-source file server designed to **upload**, **compress**, and **serve** static files effortlessly.  
Built in Go and optimized for containerized environments, it’s ideal for developers, CI/CD pipelines, and edge platforms.

![GitHub release](https://img.shields.io/github/v/release/kalevski/binvault?style=for-the-badge)
![GitHub issues](https://img.shields.io/github/issues/kalevski/binvault?style=for-the-badge)
![License](https://img.shields.io/github/license/kalevski/binvault?style=for-the-badge)

---

## 🚀 Features

- 🗃️ Bucket-based file organization (public/private)
- 🛠️ Extensible CLI-based file processing
- 🌐 REST API for uploads, downloads, and management
- 🔐 Optional JWT authorization (RSA, JWKS, PEM)
- ☁️ Ready for Docker, Kubernetes, and Swarm deployments

---

## 🧪 Quickstart

Get started with Docker:

```bash
docker pull ghcr.io/kalevski/binvault:quickstart
docker run -d -p 8080:80 ghcr.io/kalevski/binvault:quickstart
```

Then open: [http://localhost:8080](http://localhost:8080)

[📖 Full Quickstart Guide](https://binvault.io/docs.html#/quickstart)

---

## 🛠 How It Works

BinVault exposes a REST API to upload and serve files. Files are grouped into buckets and can be processed using configurable CLI commands. Public buckets are accessible via NGINX; private buckets require authorization.

[📖 Learn More](https://binvault.io/docs.html#/how-works)

---

## 📚 REST API Overview

BinVault provides CRUD endpoints for managing files and buckets.

| Method | Endpoint                         | Description            |
|--------|----------------------------------|------------------------|
| POST   | `/api/buckets`                  | Create a bucket        |
| GET    | `/api/buckets`                  | List all buckets       |
| GET    | `/api/buckets/{name}`           | Get bucket details     |
| DELETE | `/api/buckets/{name}`           | Delete a bucket        |
| GET    | `/api/buckets/{name}/files`     | List files in a bucket |
| POST   | `/api/buckets/{name}/files`     | Upload a file          |

[📖 Full API Reference](https://binvault.io/docs.html#/api)

---

## ⚙️ Configuration

Customize behavior using environment variables:

```bash
SERVER_PORT=8080
DATA_PATH=./data
PROCESSOR_CONFIG_PATH=./processors.cfg
RSA_PUBLIC_KEY=./keys/public.pem
JWT_CLAIM_ID=id
```

[🔧 Configuration Details](https://binvault.io/docs.html#/how-works?id=%f0%9f%94%90-authorization)

---

## 🤝 Contributing

1. Fork this repo
2. Create a new branch: `git checkout -b feature/your-feature`
3. Commit your changes
4. Open a pull request

See [CONTRIBUTING.md](./CONTRIBUTING.md) for details.

---

## 📜 License

Released under the **Apache 2.0 License**.