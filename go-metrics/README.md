
# Go exporting metrics sample project

## Run

### Local
```
go run main.go
```

### Docker
```
docker build . -t bygui86/go-metrics:latest
docker run -d --name go-metrics \
	-e KUBE_HOST="0.0.0.0" \
	-e MONITOR_HOST="0.0.0.0" \
	-e REST_HOST="0.0.0.0" \
	bygui86/go-metrics:latest
```

### Kubernetes
```
kubectl apply -f kube/
```

---

## APIs

### Echo server
```
:8080/echo
:8080/echo/{msg}
```

### Prometheus metrics
```
:9090/metrics
```

### Kubernetes probes
```
:9091/live
:9091/ready
```

---

## Please note

Here we use `promauto` module instead of normal `prometheus` one, so we can avoid to manually register the Prometheus collector with kind of following command

```
prometheus.MustRegister(myCustomMetric)
```

---

## Versions

### version 0.0.1
- [x] simple rest apis
- [x] default metrics

### version 0.0.2
- [x] multistage dockerfile
- [x] kubernetes manifests

### version 0.0.3
- [x] custom metrics

### version 0.0.4
- [x] improve monitoring-package segregation

---

## Links
* https://prometheus.io/docs/guides/go-application/
* https://scot.coffee/2018/12/monitoring-go-applications-with-prometheus/
* https://levelup.gitconnected.com/multi-stage-docker-builds-with-go-modules-df23b7f91a67
