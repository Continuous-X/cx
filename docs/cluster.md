# basic installation
## windows local
### openshift cluster
- https://cloud.redhat.com/openshift/install/crc/installer-provisioned
- RED HAT DEVELOPER - [Red Hat CodeReady Containers](https://access.redhat.com/documentation/en-us/red_hat_codeready_containers)
### docker
- https://docs.docker.com/docker-for-windows/install/
    - https://hub.docker.com/editions/community/docker-ce-desktop-windows/
- Use the WSL 2 based engine
    - https://docs.docker.com/docker-for-windows/wsl/
    - https://docs.microsoft.com/de-de/windows/wsl/install-win10
        - Step 4 (wsl_update_x64.msi)
#### local docker registry
- https://docs.docker.com/registry/deploying/
- ``docker run -d -e REGISTRY_HTTP_ADDR=0.0.0.0:5001 -p 5001:5001 --restart=always --name registry registry:2``
### argo cd
- ``oc create namespace argocd``
- ``mkdir argocd``
- ``cd argocd``
- ``wget https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml``
- ``oc apply -n argocd -f ./install.yaml``
- ``oc get pods -n argocd``
- ``oc get pods -l=app.kubernetes.io/name=argocd-dex-server``
- ``oc -n argocd get pod -l "app.kubernetes.io/name=argocd-server" -o jsonpath='{.items[*].metadata.name}'``
- ``oc -n argocd patch deployment argocd-server -p '{"spec":{"template":{"spec":{"$setElementOrder/containers":[{"name":"argocd-server"}],"containers":[{"command":["argocd-server","--insecure","--staticassets","/shared/app"],"name":"argocd-server"}]}}}}'``
- ``oc -n argocd create route edge argocd-server --service=argocd-server --port=http --insecure-policy=Redirect``
- ``echo https://$(oc get routes argocd-server -o=jsonpath='{ .spec.host }')``