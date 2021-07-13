# knative-hello-bnova

## Setup

### Build docker image and push to registry 
```
docker build -t bnova/knative-hello-bnova-image .
docker push bnova/knative-hello-bnova-image 
```

### Apply K8S resources
```
# create namespace
kubectl create ns knative-hello-bnova

# apply  
kubectl apply -f service.yaml
```

### Get application url
```
kubectl get ksvc knative-hello-bnova-service -n knative-hello-bnova  --output=custom-columns=NAME:.metadata.name,URL:.status.url

Output: 
NAME                          URL
knative-hello-bnova-service   http://knative-hello-bnova-service.knative-hello-bnova.1.2.3.4.sslip.io
```

## Quick Deployment to DigitalOcean's Kubernetes
[![Deploy to DO](https://www.deploytodo.com/do-btn-blue.svg)](https://cloud.digitalocean.com/apps/new?
repo=https://github.com/b-nova-techhub/knative-hello-bnova/tree/main)

## Contact
If you want to contact me you can reach me at [developer@b-nova.com](developer@b-nova.com).

## License
This project uses the following license: [MIT License](https://opensource.org/licenses/MIT)
. https://opensource.org/licenses/MIT
