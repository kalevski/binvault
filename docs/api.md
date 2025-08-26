# Rest API

BinVault provides Rest API for managing buckets and files
The API exposes endpoints for crud operations for both bucket and files


#### Create bucket

Creates a new bucket for organizing uploaded files. Buckets can be either `public` or `private`.


**POST** `/api/buckets`

**Request Headers**
- **Content-Type**: `application/json`
- **Authorization**: `Bearer {your_token}`

**Request Body**

```json
{
  "name": "my_new_bucket",
  "visibility": "public"
}
```

**Example**

```bash
curl --request POST http://localhost:8080/api/buckets \
  --header "Authorization: Bearer ${your_token}" \
  --header "Content-Type: application/json" \
  --data '{
    "name": "my_new_bucket",
    "visibility": "public"
  }'
```

---

#### List all buckets

Returns a list of all buckets accessible to the authenticated user.

**GET** `/api/buckets`

**Request Headers**
- **Authorization**: `Bearer {token}`

**Example**

```bash
curl --request GET \
  --url http://localhost:8080/api/buckets \
  --header "Authorization: Bearer {token}"
```

---

#### Fetch single bucket

Retrieves metadata for a specific bucket.

**GET** `/api/buckets/{bucket_name}`

**Request Headers**
- **Authorization**: `Bearer {token}`

**Example**

```bash
curl --request GET \
  --url http://localhost:8080/api/buckets/my_new_bucket \
  --header "Authorization: Bearer {token}"
```

---

#### Delete bucket

Deletes a specific bucket and all its associated files.

**DELETE** `/api/buckets/{bucket_name}`

**Request Headers**
- **Authorization**: `Bearer {token}`

**Example**

```bash
curl --request DELETE \
  --url http://localhost:8080/api/buckets/my_new_bucket \
  --header "Authorization: Bearer {token}"
```

---

#### List Files in Bucket

Returns a list of files stored in the specified bucket.

**GET** `/api/buckets/{bucket_name}/files`

**Request Headers**
- **Authorization**: `Bearer {token}`

**Example**

```bash
curl --request GET \
  --url http://localhost:8080/api/buckets/my_new_bucket/files \
  --header "Authorization: Bearer {token}"
```

---

#### Upload file

Uploads a file to the specified bucket.  
If file processing is enabled, the file will be transformed according to the config.

**POST** `/api/buckets/{bucket_name}/files`

**Request Headers**
- **Authorization**: `Bearer {token}`
- **Content-Type**: `multipart/form-data`

**Form Fields**
- `file`: The file to upload  
- `strict`: Set to `false` to skip validation errors

**Example**

```bash
curl --request POST \
  --url http://localhost:8080/api/buckets/my_new_bucket/files \
  --header "Authorization: Bearer {token}" \
  --header "Content-Type: multipart/form-data" \
  --form file=@{filepath} \
  --form strict=false
```