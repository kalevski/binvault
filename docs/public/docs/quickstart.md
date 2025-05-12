# 🚀 Quickstart Guide

_A step-by-step tutorial to get **BinVault** up and running as a storage service for any type of file._

> **Note**: This guide does **not** cover file processors or authorization setup.  
> • For file processing details, see [File Processor Setup](guide_processors.md).  
> • For authentication and access control, see [Authorization Guide](guide_authorization.md).

---

### ✅ Prerequisites

Ensure the following is installed on your system:

- 🐳 [Docker](https://www.docker.com/)  
  _(For bare-metal installation, refer to the [Bare-Metal Setup Guide](setup_baremetal.md))_

---

### 🐳 Quickstart Using Docker

The prebuilt **:quickstart** Docker image includes:

- **BinVault Server** – REST API for file uploads and management  
- **BinVault CLI** – Command-line tools  
- **NGINX** – Static file server for public buckets  
- **ImageMagick** – Demonstrates image compression during uploads

#### 🔹 Step 1: Pull the Docker Image

```bash
docker pull ghcr.io/kalevski/binvault:quickstart
```

---

#### 🔹 Step 2: Start the Service

```bash
docker run -d \
  --name my_vault \
  -p 8080:80 \
  ghcr.io/kalevski/binvault:quickstart
```

Once running, open your browser and navigate to [http://localhost:8080](http://localhost:8080).

You should see the BinVault default page.

![BinVault Homepage](path/to/uploaded-image.png)

---

### 📤 Upload Your First File

#### 🔹 Step 1: Create a Bucket
Run the following `curl` command to create a public bucket named `my_bucket`:
```bash
curl --request POST http://localhost:8080/api/buckets \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "my_bucket",
    "visibility": "public"
  }'
```

#### 🔹 Step 2: Upload a File
```bash
curl --request POST http://localhost:8080/api/buckets/my_bucket/files \
  --header 'Content-Type: multipart/form-data' \
  --form file=@{path_to_your_file} \
  --form strict=false
```

📌 Replace `{path_to_your_file}` with the actual file path (e.g., `./logo.png`).

---

### 🌐 Access Your Uploaded File

Your uploaded file is now accessible at:  
`http://localhost:8080/my_bucket/{filename.extension}`

🎉 Congratulations! You've successfully installed BinVault, created a bucket, and uploaded your first file.

---

### 📚 What’s Next?

Discover [How it Works](how-works.md).
