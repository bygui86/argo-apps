
# Monitoring - Grafana - Dashboards

## Missing

| Folder     | Name               | Monitored technology | Status                                  |
|------------|--------------------|----------------------|-----------------------------------------|
| auditing   | Falco              | Falco                | MISSING deployments + MISSING dashboard |
| broker     | Kafka-WebView      | Kafka-WebView        | MISSING metrics + MISSING dashboard     |
| monitoring | K8s event exporter | K8s events           | MISSING deployments + MISSING dashboard |
| support    | Kubed              | Kubed                | MISSING dashboard                       |

### Issues

| Folder     | Name                    | Issue                                 | Action                                                 |
|------------|-------------------------|---------------------------------------|--------------------------------------------------------|
| brokers    | Strimzi operators       | `JVM` group of graphs                 | verify required metrics between prometheus and grafana |
| brokers    | Strimzi Kafka           | `Broker` variable                     | fix dashboard variables                                |
| brokers    | Strimzi Zookeeper       | `Broker` variable                     | fix dashboard variables                                |
| databases  | Instaclustr-Cassandra   | `JVM metrics by node` group of graphs | verify required metrics in prometheus                  |
| indexers   | Elasticsearch - Cluster | `ConfigMap` too big for provisioning  | split into multiple ConfigMap (?)                      |
| kubernetes | API-Server              | `ETCD` graphs                         | verify required metrics                                |
| kubernetes | Controller Manager      | totally not working                   | None, not compatible with GKE                          |
| kubernetes | Proxy                   | totally not working                   | None, not compatible with GKE                          |
| kubernetes | Scheduler               | totally not working                   | None, not compatible with GKE                          |
| monitoring | Node exporter           | `ConfigMap` too big for provisioning  | split into multiple ConfigMap (?)                      |
| monitoring | Node exporter           | dashboard variables                   | fix dashboard variables                                |

---

## Important queries

### Pod restart

Total restarts (currently in use)
```
sum( kube_pod_container_status_restarts_total{job="kube-state-metrics", namespace="hdp", pod=~"hdp-reader-.*"} )
```

Restart annotation (currently in use)
```
time() == BOOL timestamp(rate(kube_pod_container_status_restarts_total{job="kube-state-metrics", namespace="$namespace", pod="$pod"}[1m]) > 0)
```

Found on internet
```
sum( increase( kube_pod_container_status_restarts_total{container="$app"} [1m] ) ) by (app)
```
