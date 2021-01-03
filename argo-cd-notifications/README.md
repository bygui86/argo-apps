
# Argo CD Notifications

## Description

Argo CD Notifications continuously monitors Argo CD applications and provides a flexible way to notify users about important changes in the application state.

Using a flexible mechanism of triggers and templates you can configure when the notification should be sent as well as notification content.

Argo CD Notifications includes the catalog of useful triggers and templates. So you can just use them instead of reinventing new ones.

## Prerequisites

- Argo CD already deployed
- Kustomize installed locally

## Instructions

1. Configure Slack following instructions [here](https://argoproj-labs.github.io/argocd-notifications/services/slack/)

1. Deploy Argo CD Notifications

  ```bash
  kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj-labs/argocd-notifications/release-1.0/manifests/install.yaml
  ```

  `(i) INFO` This command will deploy also an **empty ConfigMap** and an **empty Secret**.

1. Download Triggers and Templates ConfigMap from the catalog

  ```bash
  curl -s https://raw.githubusercontent.com/argoproj-labs/argocd-notifications/release-1.0/catalog/install.yaml -o configmap.yaml
  ```

1. Edit ConfigMap to add Slack:

  ```yaml
  # ...
  data:
    service.slack: |
      token: $slack-token
    # ...
  ```

2. Put token in secret

  ```bash
  export SLACK_TOKEN=<SLACK_TOKEN>
  sed -i "s|%SLACK_TOKEN%|$SLACK_TOKEN|g" secret.yaml
  ```

1. Configure Slack integration

  ```bash
  kustomize build . | kubectl apply -f -
  ```

1. Subscribe to notifications by `adding the notifications.argoproj.io/subscribe.on-sync-succeeded.slack` annotation to the Argo CD application or project:

  ```bash
  kubectl patch app app-of-apps -n argocd -p '{"metadata": {"annotations": {"notifications.argoproj.io/subscribe.on-sync-succeeded.slack":"<SLACK_CHANNEL>"}}}' --type merge
  ```

1. Try syncing an application and get the notification once sync is completed.

## Annotation on Application

`/!\ WARN` If you add the proper annotation to an `app-of-apps`, you won't receive notifications about single sub-applications. If you want to receive notifications also on sub-applications, you have to add the same annotation on those as well.

---

## Links

- https://argoproj-labs.github.io/argocd-notifications/
