# Configuration

BinVault supports extensive customization through environment variables, allowing you to control server behavior, file processing, and authorization.

Below is a categorized list of configuration options:

### 🌐 General

Configure core aspects of the BinVault server, such as port, hostname, and data paths.

| **Env Variable**        | **Default Value** | **Description**                              |
|--------------------------|-------------------|----------------------------------------------|
| `SERVER_PORT`           | `8080`            | The port number for the server.             |
| `SERVER_HOST`           | `localhost`       | The hostname for the server.                |
| `DATA_PATH`             | `./data`          | The path to store application data.         |
| `DB_NAME`               | `database.db`     | The name of the database file.              |
| `PROCESSOR_CONFIG_PATH` | `./processors.cfg`| The path to the processor configuration file.|

### 🔐 Authorization (Secured Endpoints)

Enable JWT-based authentication using RSA keys, JWKS, or PEM files. These variables control how user identity is verified and extracted.


#### **RSA**
Use this method when managing your own RSA key pair locally.

| **Env Variable**   | **Default Value** | **Description**                          |
|---------------------|-------------------|------------------------------------------|
| `RSA_PUBLIC_KEY`   | `""`              | The path to the SSH public key.          |
| `JWT_CLAIM_ID`     | `id`              | The claim ID for JSON Web Tokens (JWT).  |

#### **JWKS**
Use this method to verify tokens from third-party identity providers via a remote JWKS endpoint.

| **Env Variable**   | **Default Value** | **Description**                          |
|---------------------|-------------------|------------------------------------------|
| `JWKS_URL`         | `""`              | The URL for the JSON Web Key Set (JWKS). |
| `JWKS_KID`         | `main`            | The key ID for the JWKS.                 |
| `JWT_CLAIM_ID`     | `id`              | The claim ID for JSON Web Tokens (JWT).  |

#### **PEM**
Use this method with local PEM-formatted key files (X.509).

| **Env Variable**       | **Default Value** | **Description**                          |
|-------------------------|-------------------|------------------------------------------|
| `PEM_PUBLIC_FILENAME`  | `""`     | The filename for the public PEM key.     |
| `JWT_CLAIM_ID`         | `id`              | The claim ID for JSON Web Tokens (JWT).  |