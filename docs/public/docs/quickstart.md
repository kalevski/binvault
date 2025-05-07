# Quickstart

**_This is a short tutorial covering a simple scenario for using BinVault as a storage service for any type of files._**

This tutorial does not cover setting up file processors or configuring authorization. 

- To learn more about file processors, refer to [this guide](guide_processors.md).
- To learn more about authorization, refer to [this guide](guide_authorization.md).

### Prerequisites
Before you begin, ensure that [Docker](https://www.docker.com/) is installed on your machine.

For a bare-metal setup without Docker, follow [this guide](setup_baremetal.md).

---

### Quickstart with Prebuilt Docker Image

To simplify the setup process, we provide a prebuilt Docker image that includes the following components:

1. **BinVault Server**: A REST API server for uploading and managing files.
2. **BinVault CLI**: Command-line utilities for managing the BinVault service.
3. **NGINX**: An HTTP server for serving public buckets and files.
4. **ImageMagick**: A library demonstrating how images can be compressed during upload.

All these components are preconfigured and included in the Docker image, which can be pulled using the following command:

```bash
docker pull ghcr.io/kalevski/binvault:quickstart
```

#### Starting the Service

Once the image is downloaded, you can start the service with:

```bash
docker run -d --name my_vault -p 8080:80 ghcr.io/kalevski/binvault:quickstart
```

After starting the container, verify that the service is running by accessing the homepage at [http://localhost:8080](http://localhost:8080). The page should display the BinVault interface, similar to the following:

![BinVault Homepage](path/to/uploaded-image.png)

### Uploading Your First File

To upload your first file, you need to create a bucket. Use the following `curl` command to create a public bucket named `my_bucket`:

```bash
curl --request POST \
  --url http://localhost:8080/api/buckets \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "web",
	"visibility": "public"
}'
```

Once the bucket is created, you can start uploading files to it. Refer to the API 

To upload your first file to my_bucket you can use 
```bash
curl --request POST \
  --url http://localhost:8080/api/buckets/my_bucket/files \
  --header 'Content-Type: multipart/form-data' \
  --form file={path_to_your_file} \
  --form strict=false
```
