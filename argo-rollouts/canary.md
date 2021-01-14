
# Argo Rollouts - Canary Rollout scenarios

## 1. Deploy a basic Canary Rollout

1. First deployment

    ```bash
    kustomize build canary/ | kubectl apply -f -
    ```

  `(i) INFO` Initial creations of any Rollout will immediately scale up the replicas to 100% (skipping any canary upgrade steps, analysis, etc...) since there was no upgrade that occurred.

1. Watch rollout in a separate terminal window

    ```bash
    kubectl argo rollouts get rollout rollout-canary --watch
    ```

1. Wait for all pods to be up and running

1. Try access the application

    - Using `kubectl port-forward` (not always working)
        ```bash
        kubectl port-forward svc/rollout-canary 8080
        ```
        Open the browser and go to http://localhost:8080

    - Using `minikube service`
        ```bash
        minikube service --url rollout-canary
        ```
        Copy/Paste the provided link in the browser

    You should see only `blue` points.

    `/!\ WARN` Keep the browser page open, we will use it again later!

1. Change Rollout to `yellow` version

    ```bash
    kubectl apply -f canary/rollout-yellow.yaml
    ```

    When the demo rollout reaches the second step, we can see from the plugin that the Rollout is in a paused state, and now has 1 of 5 replicas running the new version of the pod template, and 4 of 5 replicas running the old version. This equates to the 20% canary weight as defined by the `setWeight: 20` step.

1. Go back to the browser, now you should see `blue` and `yellow` points.

1. Promote the Rollout

    ```bash
    kubectl argo rollouts promote rollout-canary
    ```

    After promotion, Rollout will proceed to execute the remaining steps. The remaining rollout steps in our example are fully automated, so the Rollout will eventually complete steps until it has has fully transitioned to the new version.

## 2. Abort a basic Canary Rollout

1. Change Rollout to `red` version

    ```bash
    kubectl apply -f canary/rollout-red.yaml
    ```

1. Abort the Rollout

    ```bash
    kubectl argo rollouts abort rollout-canary
    ```

    When a rollout is aborted, it will scale up the "stable" version of the ReplicaSet (in this case the yellow image), and scale down any other versions. Although the stable version of the ReplicaSet may be running and is healthy, the overall rollout is still considered `Degraded`, since the desired version (the red image) is not the version which is actually running.

2. Change Rollout back to `yellow` version in order to make Rollout considered Healthy again and not Degraded

    ```bash
    kubectl apply -f canary/rollout-yellow.yaml
    ```

    After running this command, you should notice that the Rollout immediately becomes Healthy, and there is no activity with regards to new ReplicaSets becoming created.

---

## Links

- [official example](https://argoproj.github.io/getting-started/)
- [doc](https://argoproj.github.io/features/canary/)
- [video tutorial](https://www.youtube.com/watch?v=fviYWA2mcF8)
- [medium](https://medium.com/soluto-engineering/practical-canary-releases-in-kubernetes-with-argo-rollouts-933884133aea)
