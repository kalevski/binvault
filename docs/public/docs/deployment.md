# Deployment

[-- INFRASTRUCTURE DIAGRAM --]
[-- OVERVIEW --]

---

## 🖥️ Bare Metal
Deploying BinVault directly on a physical server:​

1. Download the Binary:

```bash
wget https://github.com/kalevski/binvault/releases/latest/download/binvault
chmod +x binvault
```

1. Run the Server:

```bash
./binvault --port 8080 --data-dir /var/binvault/data
```
Ensure the specified data-dir exists and is writable.

---

## 🐳 Docker

Running BinVault in a Docker container:​
```bash
docker run -d \
  -p 8080:8080 \
  -v /path/to/data:/data \
  kalevski/binvault
```
Replace /path/to/data with your desired host directory for persistent storage.

---

## 📦 Docker Compose

Using Docker Compose for easier management:

```yaml
version: '3.8'

services:
  binvault:
    image: kalevski/binvault
    ports:
      - "8080:8080"
    volumes:
      - ./data:/data
    restart: unless-stopped
```

---

## ☸️ Kubernetes

Deploying BinVault in a Kubernetes cluster:​

1. Create a Deployment YAML:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: binvault
spec:
  replicas: 1
  selector:
    matchLabels:
      app: binvault
  template:
    metadata:
      labels:
        app: binvault
    spec:
      containers:
      - name: binvault
        image: kalevski/binvault
        ports:
        - containerPort: 8080
        volumeMounts:
        - name: data
          mountPath: /data
      volumes:
      - name: data
        persistentVolumeClaim:
          claimName: binvault-pvc
```

Create Service YAML:
```yaml
apiVersion: v1
kind: Service
metadata:
  name: binvault-service
spec:
  selector:
    app: binvault
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: LoadBalancer
```

Apply the configurations:

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

Ensure you have a PersistentVolume and PersistentVolumeClaim named binvault-pvc configured for data persistence.​