
# Argo CD Notifications

## Description

Argo CD Notifications continuously monitors Argo CD applications and provides a flexible way to notify users about important changes in the application state.

Using a flexible mechanism of triggers and templates you can configure when the notification should be sent as well as notification content.

Argo CD Notifications includes the catalog of useful triggers and templates. So you can just use them instead of reinventing new ones.

## Prerequisites

- Kustomize installed locally
- Argo CD already deployed

## Instructions

1. Configure Slack following instructions [here](https://argoproj-labs.github.io/argocd-notifications/services/slack/)

1. Install Triggers and Templates from the catalog

  ```bash
  kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj-labs/argocd-notifications/release-v1.0/catalog/install.yaml
  ```

1. Put token in secret

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
  kubectl patch app <my-app> -n argocd -p '{"metadata": {"annotations": {"notifications.argoproj.io/subscribe.on-sync-succeeded.slack":"<SLACK_CHANNEL>"}}}' --type merge
  ```

1. Try syncing an application and get the notification once sync is completed.

## Links

- https://argoproj-labs.github.io/argocd-notifications/
