
# Monitoring - Prometheus

`/!\ WARN ` CRDs are in another folder ([prometheus-crds](../../11_required-crds/11_prometheus-crds)), because they have to be deployed in advance otherwise huge additional work is required right after Prometheus deployment

`/!\ WARN ` Remember to update Prometheus Roles and RoleBindings (folder specific-namespaces) every time there is a new Kubernetes Namespace

## Versions

### Main components

| Component           | Current     | Latest               |
|---------------------|-------------|----------------------|
| kube-prometheus     | release-0.6 | release-0.6 (v1.18+) |
| prometheus-operator | 0.44.0      | 0.44.0               |
| kube-state-metrics  | 1.9.7       | 1.9.7                |
| node-exporter       | 1.0.1       | 1.0.1                |
| prometheus-adapter  | 0.8.2       | 0.8.2                |
| prometheus          | 2.23.0      | 2.23.0               |
| alertmanager        | 0.21.0      | 0.21.0               |

### Collateral components

| Component                  | Current | Latest |
|----------------------------|---------|--------|
| prometheus-config-reloader | 0.44.0  | 0.44.0 |
| kube-rbac-proxy            | 0.8.0   | 0.8.0  |

## Service monitors

`WARN: Use following command to retrieve the right labels for the ServiceMonitor`
`k get prometheus mon-prometheus-operator-prometheus -o yaml | jq .spec.serviceMonitorSelector`

### Required labels

`NONE`

## Alert rules (Prometheus rules)

`WARN: Use following command to retrieve the right labels for the PrometheusRules`
`k get prometheus mon-prometheus-operator-prometheus -o yaml | jq .spec.ruleSelector`

### Required labels

```yaml
role: alert-rules
```

## Issues

- Kubernetes additional scrape configurations
- Grafana Kubernetes dashboards
  - kubernetes controller manager
  - kubernetes proxy
  - kubernetes scheduler
  - kubernetes use method cluster (partially)

## Links

- https://github.com/coreos/prometheus-operator
- https://github.com/coreos/kube-prometheus#customizing-prometheus-alertingrecording-rules-and-grafana-dashboards
- https://sysdig.com/blog/kubernetes-monitoring-prometheus-operator-part3/
- https://github.com/mateobur/prometheus-monitoring-guide
- https://github.com/kubernetes-monitoring/kubernetes-mixin

### Prometheus alerting

- https://pracucci.com/prometheus-understanding-the-delays-on-alerting.html
