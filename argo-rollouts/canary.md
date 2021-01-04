
# Argo Rollouts - Canary Rollout scenarios

## 1. Deploy a basic Canary Rollout

1. First deployment

  ```bash
  kubectl apply -f argo-rollouts/canary/service.yaml
  kubectl apply -f argo-rollouts/canary/rollout-blue.yaml
  ```

  `(i) INFO` Initial creations of any Rollout will immediately scale up the replicas to 100% (skipping any canary upgrade steps, analysis, etc...) since there was no upgrade that occurred.

1. Observe rollout in a separate terminal window

  ```bash
  kubectl argo rollouts get rollout rollout-canary -n argo-rollouts --watch
  ```

1. Wait for all pods to be up and running

1. Try access the application

  In a terminal window run
  ```bash
  kubectl port-forward svc/rollout-canary 8080 -n argo-rollouts
  ```

  In another terminal window run (or directly open the browser)
  ```bash
  open http://localhost:8080
  ```

  You should see only `blue` points.

1. Change Rollout to `yellow` version

  ```bash
  kubectl apply -f argo-rollouts/canary/rollout-yellow.yaml
  ```

  When the demo rollout reaches the second step, we can see from the plugin that the Rollout is in a paused state, and now has 1 of 5 replicas running the new version of the pod template, and 4 of 5 replicas running the old version. This equates to the 20% canary weight as defined by the `setWeight: 20` step.

1. Try access the application

  In a terminal window run
  ```bash
  kubectl port-forward svc/rollout-canary 8080 -n argo-rollouts
  ```

  In the same browser page you should see now `blue` and `yellow` points.

1. Promote the Rollout

  ```bash
  kubectl argo rollouts promote rollout-canary -n argo-rollouts
  ```

  After promotion, Rollout will proceed to execute the remaining steps. The remaining rollout steps in our example are fully automated, so the Rollout will eventually complete steps until it has has fully transitioned to the new version.

## 2. Abort a basic Canary Rollout

1. Change Rollout to `red` version

  ```bash
  kubectl apply -f argo-rollouts/canary/rollout-red.yaml
  ```

1. Abort the Rollout

  ```bash
  kubectl argo rollouts abort rollout-canary -n argo-rollouts
  ```

  When a rollout is aborted, it will scale up the "stable" version of the ReplicaSet (in this case the yellow image), and scale down any other versions. Although the stable version of the ReplicaSet may be running and is healthy, the overall rollout is still considered `Degraded`, since the desired version (the red image) is not the version which is actually running.

1. Change Rollout back to `yellow` version in order to make Rollout considered Healthy again and not Degraded

  ```bash
  kubectl apply -f argo-rollouts/canary/rollout-yellow.yaml
  ```

  After running this command, you should notice that the Rollout immediately becomes Healthy, and there is no activity with regards to new ReplicaSets becoming created.

---

## Links

- [official example](https://argoproj.github.io/argo-rollouts/getting-started/)
- [doc](https://argoproj.github.io/argo-rollouts/features/canary/)
- [video tutorial](https://www.youtube.com/watch?v=fviYWA2mcF8)
- [medium](https://medium.com/soluto-engineering/practical-canary-releases-in-kubernetes-with-argo-rollouts-933884133aea)
