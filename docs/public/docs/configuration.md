# Configuration
BinVault allows extensive customization through environment variables. Below is a categorized list of the available environment variables:

### General

| **Env Variable**        | **Default Value** | **Description**                              |
|--------------------------|-------------------|----------------------------------------------|
| `SERVER_PORT`           | `8080`            | The port number for the server.             |
| `SERVER_HOST`           | `localhost`       | The hostname for the server.                |
| `DATA_PATH`             | `./data`          | The path to store application data.         |
| `DB_NAME`               | `database.db`     | The name of the database file.              |
| `PROCESSOR_CONFIG_PATH` | `./processors.cfg`| The path to the processor configuration file.|

## Authorization - Secured Endpoints

NOTE: JWT_CLIAM_ID


**RSA**

| **Env Variable**   | **Default Value** | **Description**                          |
|---------------------|-------------------|------------------------------------------|
| `RSA_PUBLIC_KEY`   | `""`              | The path to the SSH public key.          |
| `RSA_PRIVATE_KEY`  | `""`              | The path to the SSH private key.         |
| `JWT_CLAIM_ID`     | `id`              | The claim ID for JSON Web Tokens (JWT).  |

**JWKS**

| **Env Variable**   | **Default Value** | **Description**                          |
|---------------------|-------------------|------------------------------------------|
| `JWKS_URL`         | `""`              | The URL for the JSON Web Key Set (JWKS). |
| `JWKS_KID`         | `main`            | The key ID for the JWKS.                 |
| `JWT_CLAIM_ID`     | `id`              | The claim ID for JSON Web Tokens (JWT).  |

**PEM**

| **Env Variable**       | **Default Value** | **Description**                          |
|-------------------------|-------------------|------------------------------------------|
| `PEM_PRIVATE_FILENAME` | `key.pem`         | The filename for the private PEM key.    |
| `PEM_PUBLIC_FILENAME`  | `key_pub.pem`     | The filename for the public PEM key.     |
| `JWT_CLAIM_ID`         | `id`              | The claim ID for JSON Web Tokens (JWT).  |

## Configure File Processors (processors.cfg)

Simple configuration

```cfg
default=mv {{.Source}} {{.Target}}
```