
# Argo CD

## Tested scenarios

- [x] simple app with kustomize
- [x] custom app in golang
- [x] database
- [x] operator (e.g. prometheus)
- [x] namespaces first
- [x] [app-of-apps](https://argoproj.github.io/argo-cd/operator-manual/cluster-bootstrapping/)
- [x] [hooks](https://argoproj.github.io/argo-cd/user-guide/resource_hooks/)
  - [x] pre
  - [x] post
- [x] [sync wave](https://argoproj.github.io/argo-cd/user-guide/sync-waves/)
- [x] [metrics](https://argoproj.github.io/argo-cd/operator-manual/metrics/)
- [x] [private repo](https://argoproj.github.io/argo-cd/user-guide/private-repositories/)
- [x] [notifications](https://argoproj.github.io/argo-cd/operator-manual/notifications/)
  - [x] [argo cd notifications](https://argoproj-labs.github.io/argocd-notifications/)
- [ ] [declarative setup](https://argoproj.github.io/argo-cd/operator-manual/declarative-setup/)
  - [ ] [repository credentials](https://argoproj.github.io/argo-cd/operator-manual/declarative-setup/#repository-credentials)

## Install Argo CD CLI

See [here](https://argoproj.github.io/argo-cd/getting_started/#2-download-argo-cd-cli)

## Deploy Argo CD

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

## Configure private repo

1. Create an `Access token` in your git service

  - [GitHub](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line)
  - [GitLab](https://docs.gitlab.com/ee/user/project/deploy_tokens/)
  - [Bitbucket](https://confluence.atlassian.com/bitbucketserver/personal-access-tokens-939515499.html)

  For example GitHub access token must be created with complete `Repo` access (all ticks)

1. Using Argo CD CLI, execute following command

  ```bash
  argocd repo add https://github.com/bygui86/argo-private-app.git --username bygui86 --password <GITHUB_ACCESS_TOKEN>
  ```

1. Using Argo UI

    1. navigate to `Settings / Repositories`

    1. Click `Connect Repo using HTTPS` and enter credentials

        - Type: `git`
        - Repository URL: `https://github.com/bygui86/argo-private-app`
        - Username: `bygui86`
        - Password: `<GITHUB_ACCESS_TOKEN>`

        `(i) INFO` For some services (e.g. GitHub), you might have to specify your account name as the username instead of any string.

        `(i) INFO` For some services, you might have to specify `.git` at the end of repository URL.

## Deploy apps

```bash
kubectl apply -f argo-cd/app-of-apps.yaml
```

---

## Links

- [argocd-example-apps](https://github.com/bygui86/argocd-example-apps)
  - [kustomize app](https://github.com/bygui86/argocd-example-apps/tree/master/kustomize-guestbook)
  - [sync waves](https://github.com/bygui86/argocd-example-apps/tree/master/sync-waves)
  - [hooks](https://github.com/bygui86/argocd-example-apps/tree/master/pre-post-sync)
