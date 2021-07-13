# knative-hello-bnova

### Build docker image and push to registry 
```
docker build -t bnova/knative-hello-bnova .
docker push bnova/knative-hello-bnova 
```

### Apply K8S resources
```
kubectl apply -f service.yaml
```

### Get application url
```
kubectl get ksvc knative-hello-bnova -n knative-hello-bnova  --output=custom-columns=NAME:.metadata.name,URL:.status.url

Output: 
NAME                  URL
knative-hello-bnova   http://knative-hello-bnova.knative-hello-bnova.1.2.3.4.sslip.io
```

