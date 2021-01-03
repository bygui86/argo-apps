
# Argo Rollouts

## Tested scenarios

- [ ] rollouts
  - [x] [basic canary](https://argoproj.github.io/argo-rollouts/getting-started/)
  - [ ] [blue/green](https://argoproj.github.io/argo-rollouts/features/bluegreen/)
  - [ ] [advanced canary](https://argoproj.github.io/argo-rollouts/features/canary/)
- [ ] [Metrics](https://argoproj.github.io/argo-rollouts/features/controller-metrics/)
- [ ] [Analysis based on metrics](https://argoproj.github.io/argo-rollouts/features/analysis/#background-analysis)
- [ ] [Argo CD integration](https://argoproj.github.io/argo-rollouts/FAQ/#how-does-argo-rollouts-integrate-with-argo-cd)
  - [ ] [Argo CD resource health](https://argoproj.github.io/argo-cd/operator-manual/health/)
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

## 1. Deploy a Rollout with Canary strategy

1. Observe rollout in a separate terminal window

  ```bash
  kubectl argo rollouts get rollout rollouts-demo --watch
  ```

1. First deployment

  ```bash
  kubectl apply -f argo-rollouts/basic-canary/rollout-blue.yaml
  kubectl apply -f argo-rollouts/basic-canary/service.yaml
  ```

  `(i) INFO` Initial creations of any Rollout will immediately scale up the replicas to 100% (skipping any canary upgrade steps, analysis, etc...) since there was no upgrade that occurred.

1. Change Rollout to `yellow` version

  ```bash
  kubectl apply -f argo-rollouts/basic-canary/rollout-yellow.yaml
  ```

  When the demo rollout reaches the second step, we can see from the plugin that the Rollout is in a paused state, and now has 1 of 5 replicas running the new version of the pod template, and 4 of 5 replicas running the old version. This equates to the 20% canary weight as defined by the `setWeight: 20` step.

1. Promote the Rollout

  ```bash
  kubectl argo rollouts promote rollouts-demo
  ```

  After promotion, Rollout will proceed to execute the remaining steps. The remaining rollout steps in our example are fully automated, so the Rollout will eventually complete steps until it has has fully transitioned to the new version.

## 2. Abort a Rollout with Canary strategy

1. Change Rollout to `red` version

  ```bash
  kubectl apply -f argo-rollouts/basic-canary/rollout-red.yaml
  ```

1. Abort the Rollout

  ```bash
  kubectl argo rollouts abort rollouts-demo
  ```

  When a rollout is aborted, it will scale up the "stable" version of the ReplicaSet (in this case the yellow image), and scale down any other versions. Although the stable version of the ReplicaSet may be running and is healthy, the overall rollout is still considered `Degraded`, since the desired version (the red image) is not the version which is actually running.

1. Change Rollout back to `yellow` version in order to make Rollout considered Healthy again and not Degraded

  ```bash
  kubectl apply -f argo-rollouts/basic-canary/rollout-yellow.yaml
  ```

  After running this command, you should notice that the Rollout immediately becomes Healthy, and there is no activity with regards to new ReplicaSets becoming created.

### TBD

`/!\ WARN` go-k8s-probes requires PostgreSQL deployed in `postgres` namespace.

1. Deploy PostgreSQL

  ```bash
  kubectl apply -f argo-rollouts/custom/postgresql.yaml
  ```

1. Deploy app

  ```bash
  kubectl apply -f argo-rollouts/custom/go-k8s-probes.yaml
  ```

1. Go to [repo](https://github.com/bygui86/go-k8s-probes)

1. Uncomment version `v1.1` in `kube/kustomization.yaml` and push

---

## Links

- [blue/green](https://github.com/bygui86/argocd-example-apps/tree/master/blue-green)
- [tutorial](https://medium.com/soluto-engineering/practical-canary-releases-in-kubernetes-with-argo-rollouts-933884133aea)
