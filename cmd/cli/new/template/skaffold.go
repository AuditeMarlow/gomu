package template

var SkaffoldCFG = `---

apiVersion: skaffold/v2beta21
kind: Config
metadata:
  name: {{.Alias}}
build:
  artifacts:
  - image: {{.Alias}}
deploy:
  kubectl:
    manifests:
    - skaffold/*.yaml
`

var SkaffoldDEP = `---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Alias}}
  labels:
    app: {{.Alias}}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{.Alias}}
  template:
    metadata:
      labels:
        app: {{.Alias}}
    spec:
      containers:
      - name: {{.Alias}}
        image: {{.Alias}}:latest
        args:
        - --server_name=go.micro.srv.{{lower .Alias}}
`
