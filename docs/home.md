# BinVault

A **lightweight**, open-source file server designed to **upload**, **compress**, and **serve** static files effortlessly.

**BinVault** is a high-performance file storage solution that offers a straightforward HTTP API for uploading and organizing files into buckets. Uploaded files can be automatically compressed and stored on disk, ensuring instant and reliable static delivery.

Developed in **Go**, BinVault is distributed as a single executable, facilitating hassle-free deployment. You can run it:

- 🐳 Inside a **Docker** container
- ☁️ With orchestration tools like **Kubernetes**, **Docker Compose**, or **Swarm**

### 🚀 Key Features
- 🌍 Simplified and portable deployment
- 🗜️ Automatic file processing during uploads
- 🔐 Seamless integration with your authentication service or internal APIs
- ⚙️ Compatible with any external CLI tools

### 📦 Use Cases

#### File Upload Backend

Utilize BinVault as an origin server to handle and compress user-uploaded images within custom workflows.

**Example scenario:**
- Upload user files via the API
- Store compressed thumbnails and original files
- Serve them through `/bucket/file.jpg` routes

✅ Ideal for applications that manage images or files.

---

#### CI/CD Artifact Storage

Use BinVault as a lightweight alternative to object storage for retaining build artifacts or test reports.

**Example scenario:**
- Upload `.zip` or `.json` test results from CI/CD pipelines like GitHub Actions or GitLab CI
- Access them via public or token-protected buckets

✅ Perfect for development teams and automation pipelines.

---

> Whether you're powering image delivery, managing static assets, or building your own media pipeline — **BinVault** is ready to serve.
