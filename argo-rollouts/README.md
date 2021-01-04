
# Argo Rollouts

## Tested scenarios

- [X] rollouts
  - [x] [basic canary](https://argoproj.github.io/argo-rollouts/getting-started/)
  - [x] [blue/green](https://argoproj.github.io/argo-rollouts/features/bluegreen/)
- [ ] [Argo CD integration](https://www.youtube.com/watch?v=35Qimb_AZ8U)
  - [ ] [blue/green using Argo CD](https://github.com/bygui86/argocd-example-apps/tree/master/blue-green)
- [ ] [Metrics](https://argoproj.github.io/argo-rollouts/features/controller-metrics/)
- [ ] [Analysis based on metrics](https://argoproj.github.io/argo-rollouts/features/analysis/#background-analysis)
- [ ] [Kustomize integration](https://argoproj.github.io/argo-rollouts/features/kustomize/)

## Install Argo Rollouts kubectl plugin

See [here](https://argoproj.github.io/argo-rollouts/installation/#kubectl-plugin-installation)

`(i) INFO` The Argo Rollouts kubectl plugin allows you to visualize the Rollout, its related resources (ReplicaSets, Pods, AnalysisRuns), and presents live state changes as they occur.

`/!\ WARN` Argo Rollouts kubectl plugin is required to observe rollouts.

## Deploy Argo Rollouts

```bash
kubectl create namespace argo-rollouts
kubectl apply -n argo-rollouts -f https://raw.githubusercontent.com/argoproj/argo-rollouts/stable/manifests/install.yaml
```

## Scenarios

`(i) INFO` All scenarios use `argoproj/rollouts-demo` container image, see [here](container-images.md) for all available tags.

### [Canary Rollouts](canary.md)

### [Blue/Green Rollouts](blue-green.md)
