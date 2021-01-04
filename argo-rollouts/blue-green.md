
# Argo Rollouts - Blue/Green Rollout scenarios

## 1. Deploy a Blue/Green Rollout

1. Deploy `blue` Rollout using Argo CD Application

  ```bash
  kubectl apply -f argo-rollouts/blue-green/service.yaml
  kubectl apply -f argo-rollouts/blue-green/rollout-green.yaml
  ```

  At the beginning, both services `rollout-bluegreen-active` and `rollout-bluegreen-preview` have the same endpoints, pointing to `argoproj/rollouts-demo:green` pods

1. Observe rollout in a separate terminal window

  ```bash
  kubectl argo rollouts get rollout rollout-bluegreen -n default --watch
  ```

1. Wait for all pods to be up and running

1. Try access the application

  In a terminal window run
  ```bash
  kubectl port-forward svc/rollout-bluegreen-active 8080 -n default
  ```

  In another terminal window run (or directly open the browser)
  ```bash
  open http://localhost:8080
  ```

  You should see only `green` points.

1. Deploy `blue` Rollout using Argo CD Application

  ```bash
  kubectl apply -f argo-rollouts/blue-green/rollout-blue.yaml
  ```

1. Wait for all pods to be up and running

1. Observe status using Argo Rollouts kubectl plugin in the other window previously opened

  Now application runs `argoproj/rollouts-demo:green` and `argoproj/rollouts-demo:blue` simultaneously. The `argoproj/rollouts-demo:blue` is still considered `blue` available only via preview service `rollout-bluegreen-preview`, while `rollout-bluegreen-active` still serves `argoproj/rollouts-demo:green`

1. Try access the applications

  In a terminal window run
  ```bash
  kubectl port-forward svc/rollout-bluegreen-active 8080 -n default
  ```

  In another terminal window run (or directly open the browser)
  ```bash
  open http://localhost:8080
  ```

  You should still see only `green` points.

  Now close port-forward to active service and run
  ```bash
  kubectl port-forward svc/rollout-bluegreen-preview 8080 -n default
  ```

  In the same browser page you should see now only `blue` points.

1. Promote the Rollout

  ```bash
  kubectl argo rollouts promote rollout-bluegreen -n default
  ```

  After promotion, `blue` will be promoted as `active` and stable.

1. Try access the application

  In a terminal window run
  ```bash
  kubectl port-forward svc/rollout-bluegreen-active 8080 -n default
  ```

  In another terminal window run (or directly open the browser)
  ```bash
  open http://localhost:8080
  ```

  You should see only `blue` points.

## 2. Deploy a Blue/Green Rollout using Argo CD - `WIP`

`/!\ WARN` This scenario includes also Argo CD integration.

1. Deploy `blue` Rollout using Argo CD Application

  ```bash
  kubectl apply -f argo-rollouts/rollout-bluegreen-app.yaml
  ```

1. Observe rollout in a separate terminal window

  ```bash
  kubectl argo rollouts get rollout rollout-bluegreen -n bluegreen --watch
  ```

`TBD`

1. Promote the Rollout

  ```bash
  kubectl argo rollouts promote rollout-bluegreen -n bluegreen
  ```

  After promotion, Rollout will proceed to execute the remaining steps. The remaining rollout steps in our example are fully automated, so the Rollout will eventually complete steps until it has has fully transitioned to the new version.

---

## Links

- [official repo](https://github.com/bygui86/argocd-example-apps/tree/master/blue-green)
- [video tutorial](https://www.youtube.com/watch?v=krDxDz4V4Tg)
