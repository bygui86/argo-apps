
# Argo Rollouts - Experiment scenario

## Description

This example demonstrates an experiment which starts two ReplicaSets with different images and run for 5 minutes once they both become available.

While the experiment is running, it runs a job-based analysis, which performs HTTP benchmarking against each replicasets via their respective services.

`($) SPOIL` The experiment uses intentionally "bad" version of the rollouts-demo image, which has a high error rate and will cause the experiment to fail.

## Instructions

1. Deploy the Experiment

    ```bash
    kustomize build experiment/ | kubectl apply -f -
    ```

1. Watch the Experiment in a separate terminal window

    ```bash
    kubectl argo rollouts get experiment experiment-demo --watch
    ```

1. Wait for the Experiment to be completed

1. The outcome should be like following

    - `purple` analysis should be 100% successfull

    - `orage` analysis should be 100% failing (due to the "bad" version of the rollouts-demo image used)

    ```
    Name:            experiment-demo
    Namespace:       default
    Status:          ✔ Successful

    NAME                                                            KIND         STATUS        AGE    INFO
    Σ experiment-demo                                               Experiment   ✔ Successful  17m
    ├──⧉ experiment-demo-orange                                     ReplicaSet   • ScaledDown  17m
    ├──⧉ experiment-demo-purple                                     ReplicaSet   • ScaledDown  17m
    ├──α experiment-demo-purple                                     AnalysisRun  ✔ Successful  17m    ✔ 69
    │  └──⊞ 9f677bed-5458-4a3d-9f06-fba58867e821.http-benchmark.69  Job          ✔ Successful  2m20s
    └──α experiment-demo-orange                                     AnalysisRun  ✔ Successful  17m    ✖ 69
    └──⊞ 0b04112b-de07-4e0c-97f2-293e3131d779.http-benchmark.69  Job          ✖ Failed      2m30s
    ```

---

## Links

- [doc](https://argoproj.github.io/argo-rollouts/features/experiment/)
