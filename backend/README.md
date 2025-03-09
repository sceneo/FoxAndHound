# Test build process locally

## Build Backend

```bash
cd backend
docker build . -f ../infra/docker/build_backend.Dockerfile -t foxandhound_backend_local
docker tag foxandhound_backend_local johannesrosskopp/my_private_repository:foxandhound_backend_dev_latest
docker push johannesrosskopp/my_private_repository:foxandhound_backend_dev_latest
```

for a release version also tag and push an image with a version like foxandhound_backend_dev_1_0_0

## Run databse locally

```bash
cd database
docker-compose up
```

## Run backend container with host network

```bash
docker run --init --network="host" -e STAGE=dev foxandhound_backend_local
```

When localhost port 8080 is used you can do the following

Find you local machines ip address, e.g. 

```bash
ip addr show
``` 

and save it as 

```bash
LOCAL_IP=XXX
```

then use

```bash
docker run \
    -e DB_HOST=192.168.178.32 \
    -e STAGE=dev \
    -p 8085:8080 \
    foxandhound_backend_local
```

This will allow the image to access your local machine running the sql databse and map container 8080 to local 8085 (choose any free port instead of 8085).
The backend is the reachable under `localhost:8085`, e.g. `curl localhost:8085/api/rating-cards`.