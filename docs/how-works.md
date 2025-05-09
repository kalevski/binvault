# How it Works

![BinVault diagram](https://github-production-user-asset-6210df.s3.amazonaws.com/10467454/441719425-0b6634fa-f0ff-40f3-8139-3e46b2a3aa4d.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVCODYLSA53PQK4ZA%2F20250508%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20250508T130947Z&X-Amz-Expires=300&X-Amz-Signature=51369108d1722477cac91a1f48f32d549763e53b0012280c3aa1fc738c7c7964&X-Amz-SignedHeaders=host)

---

**BinVault** is a lightweight web server distributed as a single binary. By default, it exposes a RESTful API on port **8080**.

The API provides endpoints to manage two core resource types:

#### 📦 Buckets

Buckets are logical containers used to organize and manage files.

Each bucket can be configured as either:

- 🔓 **Public**
  - A symbolic link is created in the public directory.
  - Files are accessible via static file servers (e.g., NGINX) without authentication.
- 🔒 **Private**
  - Files are only accessible via the API, with appropriate permissions.

This public/private distinction allows flexible control over how content is served and secured.

#### 📁 Files

**Files** are the primary data objects stored inside buckets. You can:

- **Uploaded**: Add files to a bucket via the API.
- **Downloaded**: Retrieve files via the API or static endpoint (if public).

---

### ⚙️ Processing

**BinVault** supports file processing at a **global level**.
You can configure processing rules based on file extensions, with a default fallback if no rule is defined for a specific extension.

#### 📄 Configuration File

To enable processing, provide the path to a configuration file using the environment variable:

```bash
PROCESSOR_CONFIG_PATH=/path/to/processors.cfg
```

In this case, files are not altered during upload—they are simply moved to the serving directory without processing.



The processor configuration file maps file extensions to CLI commands. A minimal example might look like:

```processors.cfg
default=mv {{.Source}} {{.Target}}
```

> To learn more about using processors for tasks like resizing, compressing, and more, check out this [guide](guide_processors.md).

---

### 🔐 Authorization

Authorization in **BinVault** is **optional** and disabled by default. If enabled, it secures the API using **JWT (JSON Web Tokens)** to authorize users.

This system allows you to control access to buckets, files, and administrative actions using signed tokens—either from your own service or a third-party identity provider.

#### How It Works

1. Clients send requests to protected endpoints with a JWT in the Authorization header.

2. BinVault validates the token using one of the supported verification methods:
  - 🔑 RSA public key
  - 🌐 JWKS (JSON Web Key Set)
  - 📄 PEM-formatted keys

3. A specific claim from the token (default: id) is extracted to identify the user or service.

> For example: If your JWT contains a claim named username, and that value uniquely identifies the user in your system, you can set JWT_CLAIM_ID=username to use it as the owner of the uploaded file in BinVault.

To learn more about authorization and how to configure it, check out this [guide](guide_authorization.md).

### Internal database

BinVault includes a built-in **SQLite** database to manage and persist metadata related to your files and buckets. While the actual files are stored on disk, the database tracks their **structure**, **ownership**, and **status**.