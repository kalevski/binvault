# Quickstart

**_This is a short tutorial covering a simple scenario for using BinVault as a storage service for any type of files._**

This tutorial does not cover setting up file processors or configuring authorization. 

- To learn more about file processors, refer to [this guide](guide_processors.md).
- To learn more about authorization, refer to [this guide](guide_authorization.md).

### Prerequisites
Before you begin, ensure that [Docker](https://www.docker.com/) is installed on your machine.

For a bare-metal setup without Docker, follow [this guide](setup_baremetal.md).

---

For quick start we created a image that includes

1. binvault server - Server that exposes rest API for uploading / managing files
1. binvault cli - bundled utilities for managing binvault service
1. NGINX - HTTP server for serving public buckets & files 
1. imagemagick - library as an example how images can be compressed when they are uploaded

All of this are included and already setuped in image that can be downloaded using

`docker pull ghcr.io/kalevski/binvault:quickstart`

