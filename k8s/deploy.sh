docker build . -f Dockerfile.api -t eu.gcr.io/kite-prod-297314/api:$(git rev-parse --short HEAD)
docker push eu.gcr.io/kite-prod-297314/api:$(git rev-parse --short HEAD)
