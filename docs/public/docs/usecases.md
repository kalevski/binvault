# 📦 Use Cases for BinVault

BinVault is a flexible, developer-friendly solution for managing static files at scale. Below are some common ways it can be used across different environments and stacks.

---
## Self-Hosted Image Proxy

Use BinVault as an origin server to handle and compress user-uploaded images in custom workflows.

**Example scenario:**
- Upload user avatars or uploads via API
- Store compressed thumbnails and originals
- Serve them via `/bucket/file.jpg` routes

✅ Great for user profile systems, forums, or SaaS platforms

---

## CI/CD Artifact Storage

Use BinVault as a lightweight alternative to object storage for storing build artifacts or test reports.

**Example scenario:**
- Upload `.zip` or `.json` test results from GitHub Actions or GitLab CI
- Access via a public or token-protected bucket
- Use file versioning to keep deployment history

✅ Perfect for dev teams and automation pipelines

---


## Private File Storage

Create protected buckets with restricted access for private content or internal tools.

**Example scenario:**
- Use API keys to upload files
- Limit read/write permissions per bucket
- Serve private documents or assets securely

✅ Useful for internal apps, B2B platforms, or client dashboards

---

## Developer File Gateway

Let other services push or pull files from BinVault using a RESTful interface.

**Example scenario:**
- Upload files from a Node.js or Python app
- Use signed URLs for secure temporary access

✅ Ideal for microservices needing shared file storage

---

## Image CDN Backend

Use BinVault to store, compress, and serve images directly or behind a CDN like Cloudflare or Fastly.

**Example scenario:**
- Upload product photos to a bucket via API
- BinVault automatically compresses and versions them
- Serve them with proper caching and content-type headers

✅ Ideal for e-commerce, media platforms, and content-heavy websites

---

## Frontend Build Hosting

Host your static site or frontend app directly from BinVault.

**Example scenario:**
- Push your `build/` output from a CI/CD pipeline (coming soon)
- Serve versioned bundles via unique URLs
- Set cache headers for long-term browser caching

✅ Great for SPAs, marketing sites, and dashboards

---

> 💡 Have a cool use case for BinVault? [Submit a feature request](https://github.com/kalevski/binvault/issues/new?template=feature_request.md)!