
# argo-apps

Applications to test Argo project.

## TODOs

- [x] simple app with kustomize
- [x] custom app in golang
- [x] database
- [x] operator (e.g. prometheus)
- [x] namespaces first
- [x] [app-of-apps](https://argoproj.github.io/argo-cd/operator-manual/cluster-bootstrapping/)
- [ ] [hooks](https://argoproj.github.io/argo-cd/user-guide/resource_hooks/)
  - [ ] pre
  - [ ] post
- [x] [sync wave](https://argoproj.github.io/argo-cd/user-guide/sync-waves/)
- [ ] [metrics](https://argoproj.github.io/argo-cd/operator-manual/metrics/) - `WIP`
- [x] [private repo](https://argoproj.github.io/argo-cd/user-guide/private-repositories/)
- [ ] [notifications](https://argoproj.github.io/argo-cd/operator-manual/notifications/)
  - [ ] [argocd notifications](https://github.com/argoproj-labs/argocd-notifications)
  - [ ] [argo kube notifier](https://github.com/argoproj-labs/argo-kube-notifier)
  - [ ] [kube watch](https://github.com/bitnami-labs/kubewatch)
- [ ] [declarative setup](https://argoproj.github.io/argo-cd/operator-manual/declarative-setup/)
  - [ ] [repository credentials](https://argoproj.github.io/argo-cd/operator-manual/declarative-setup/#repository-credentials)

## Integrations

- [ ] argo rollouts
  - [ ] blue/green
  - [ ] canary
- [ ] tekton
- [ ] multi-cluster

## Examples

- [kustomize app](https://github.com/bygui86/argocd-example-apps/tree/master/kustomize-guestbook)
- [sync waves](https://github.com/bygui86/argocd-example-apps/tree/master/sync-waves)
- [hooks](https://github.com/bygui86/argocd-example-apps/tree/master/pre-post-sync)
- [blue/green](https://github.com/bygui86/argocd-example-apps/tree/master/blue-green)
