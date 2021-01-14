
# Argo Rollouts

## Tested scenarios

- [X] rollouts
  - [x] [basic canary](https://argoproj.github.io/argo-rollouts/getting-started/)
  - [x] [blue/green](https://argoproj.github.io/argo-rollouts/features/bluegreen/)
- [x] [Argo CD integration](https://www.youtube.com/watch?v=35Qimb_AZ8U)
  - [x] [blue/green using Argo CD](https://github.com/bygui86/argocd-example-apps/tree/master/blue-green)
- [x] [Expose metrics](https://argoproj.github.io/argo-rollouts/features/controller-metrics/)
- [x] [Custom metrics example](https://argoproj.github.io/argo-rollouts/features/analysis/)
- [x] [Experiment](https://argoproj.github.io/argo-rollouts/features/experiment/)

## Install Argo Rollouts kubectl plugin

See [here](https://argoproj.github.io/argo-rollouts/installation/#kubectl-plugin-installation)

`(i) INFO` The Argo Rollouts kubectl plugin allows you to visualize the Rollout, its related resources (ReplicaSets, Pods, AnalysisRuns), and presents live state changes as they occur.

`/!\ WARN` Argo Rollouts kubectl plugin is required to observe rollouts.

## Deploy Argo Rollouts

1. Create `argo-rollouts` namespace

    ```bash
    kubectl create namespace argo-rollouts
    ```

1. Deploy Argo Rollouts

    ```bash
    kubectl apply -n argo-rollouts -f https://raw.githubusercontent.com/argoproj/argo-rollouts/stable/manifests/install.yaml
    ```

## Scenarios

`(i) INFO` All scenarios use `argoproj/rollouts-demo` container image, see [here](container-images.md) for all available tags.

### [Canary Rollouts](canary.md)

### [Blue/Green Rollouts (including Argo CD integration)](blue-green.md)

### [Expose metrics](expose-metrics.md)

### [Experiment](experiment.md)

---

## Links

- [argo-rollouts-examples](https://github.com/bygui86/rollouts-demo)
