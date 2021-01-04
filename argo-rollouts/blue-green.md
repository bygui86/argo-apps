
# Argo Rollouts - Blue/Green Rollout scenarios

## 1. Deploy a Blue/Green Rollout

1. Deploy `green` Rollout using Argo CD Application

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

1. Try access the `active (green)` application

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

1. Try access `preview (blue)` application

    In a terminal window run

    ```bash
    kubectl port-forward svc/rollout-bluegreen-preview 8081 -n default
    ```

    In another terminal window run (or directly open the browser)
  
    ```bash
    open http://localhost:8081
    ```

    You should see now `blue` points, while in the previous browser page you should still only `green` points.

1. Promote the Rollout

    ```bash
    kubectl argo rollouts promote rollout-bluegreen -n default
    ```

    After promotion, `blue` will be promoted as `active` and stable.

1. Try access the `active (now blue)` application

    In a terminal window run

    ```bash
    kubectl port-forward svc/rollout-bluegreen-active 8080 -n default
    ```

    In another terminal window run (or directly open the browser)

    ```bash
    open http://localhost:8080
    ```

    You should see only `blue` points.

## 2. Deploy a Blue/Green Rollout using Argo CD integration

`/!\ WARN` This scenario includes also Argo CD integration.

1. Deploy `green` Rollout using Argo CD Application

    ```bash
    kubectl apply -f argo-rollouts/rollout-bluegreen-app.yaml
    ```

    At the beginning, both services `rollout-bluegreen-active` and `rollout-bluegreen-preview` have the same endpoints, pointing to `argoproj/rollouts-demo:green` pods

1. Observe rollout in Argo CD UI

    ```bash
    kubectl port-forward svc/argocd-server -n argocd 8080:443
    ```

    or in a separate terminal window

    ```bash
    kubectl argo rollouts get rollout rollout-bluegreen -n bluegreen --watch
    ```

1. Wait for all pods to be up and running

1. Try access the `green` application

    In a terminal window run

    ```bash
    kubectl port-forward svc/rollout-bluegreen-active 8081:8080 -n bluegreen
    ```

    In another terminal window run (or directly open the browser)
  
    ```bash
    open http://localhost:8081
    ```

    You should see only `green` points.

1. Change container image tag parameter to trigger blue-green deployment process

    ```bash
    argocd app set rollout-bluegreen --kustomize-image argoproj/rollouts-demo:blue
    ```

    `(i) INFO` You should be able to perform this action also through Argo CD UI under `Applications / rollout-bluegreen / (rollout) rollout-bluegreen / Live manifest EDIT`

1. Observe blue/green process in Argo CD UI

    Now application runs `argoproj/rollouts-demo:green` and `argoproj/rollouts-demo:blue` simultaneously. The `argoproj/rollouts-demo:blue` is still considered `blue` available only via preview service `rollout-bluegreen-preview`, while `rollout-bluegreen-active` still serves `argoproj/rollouts-demo:green`

1. Try access the `blue` application

    In a terminal window run

    ```bash
    kubectl port-forward svc/rollout-bluegreen-preview 8082:8080 -n bluegreen
    ```

    In another terminal window run (or directly open the browser)

    ```bash
    open http://localhost:8082
    ```

    You should see now `blue` points, while in the previous browser page you should still only `green` points.

1. Promote rollout to `blue` 

    Through Argo CD UI
  
    `Applications / rollout-bluegreen / (rollout) rollout-bluegreen / right 3 dots / Promote-full`

    By patching Rollout resource using Argo CD CLI

    ```bash
    argocd app patch-resource rollout-bluegreen --kind Rollout --resource-name rollout-bluegreen --patch '{ "status": { "verifyingPreview": false } }' --patch-type 'application/merge-patch+json'
    ```

    or using Argo Rollouts kubectl plugin

    ```bash
    kubectl argo rollouts promote rollout-bluegreen -n bluegreen
    ```

    This promotes rollout to `blue` status and Rollout deletes old replica which runs `green`.
  After promotion, `blue` will be promoted as `active` and stable.

1. Try access the `active (now blue)` application

    In a terminal window run

    ```bash
    kubectl port-forward svc/rollout-bluegreen-active 8081 -n default
    ```

    In another terminal window run (or directly open the browser)
    
    ```bash
    open http://localhost:8081
    ```

    You should see only `blue` points.

---

## Links

- [official repo](https://github.com/bygui86/argocd-example-apps/tree/master/blue-green)
- [video tutorial](https://www.youtube.com/watch?v=krDxDz4V4Tg)
