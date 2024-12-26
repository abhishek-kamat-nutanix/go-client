# Kubernetes Application Migration Tool

## Overview  
This tool facilitates the migration of applications from one Kubernetes cluster to another by transferring persistent volumes and application configurations using Kubernetes APIs and GRPC.

---

## Prerequisites  

### Source Cluster  
- Ensure a valid **VolumeSnapshotClass** exists. A sample is given which uses Nutanix CSI Driver [VSC](https://github.com/abhishek-kamat-nutanix/go-client/blob/master/k8s/vsc.yaml)

### Target Cluster  
- A default **VolumeSnapshotClass** must be available.  
- A default **StorageClass** must be available.  

---

## Installation  

### Install Necessary Roles and Bindings  
Apply the following RBAC configurations on the **target cluster**:  
- [Cluster Role](https://github.com/abhishek-kamat-nutanix/go-client/blob/master/k8s/clusterrole.yaml)  
- [Cluster Role Binding](https://github.com/abhishek-kamat-nutanix/go-client/blob/master/k8s/clusterrolebinding.yaml)  
- [Role](https://github.com/abhishek-kamat-nutanix/go-client/blob/master/k8s/role.yaml)  
- [Role Binding](https://github.com/abhishek-kamat-nutanix/go-client/blob/master/k8s/rolebinding.yaml)  

Deploy each using the following commands:

```bash
kubectl apply -f <path-to-role.yaml>
kubectl apply -f <path-to-rolebinding.yaml>
kubectl apply -f <path-to-clusterrole.yaml>
kubectl apply -f <path-to-clusterrolebinding.yaml>
```

# Writer Server Deployment and Usage

## Deploy the Writer Server

To deploy the Writer Server, apply the `writer-deployment.yaml` file on the **target Kubernetes cluster**: [writer-deployment.yaml](https://github.com/abhishek-kamat-nutanix/read-write-grpc/blob/master/writer-deployment.yaml)

```bash
kubectl apply -f https://github.com/abhishek-kamat-nutanix/read-write-grpc/blob/master/writer-deployment.yaml
```

---

## Usage Instructions

### Retrieve Service IPs

1. On the **source cluster**, get the IP of the `reader-server` service:

   ```bash
   kubectl get service reader-server-service -n <namespace>
   ```

2. On the **target cluster**, get the IP of the `writer-server` service:

   ```bash
   kubectl get service grpc-server-service -n <namespace>
   ```

### Provide IPs to the Orchestrator

Use the IPs obtained above in the GRPC request sent to the orchestrator.

#### Example GRPC Command:

```bash
grpcurl -d '{
  "namespace": "wordpress",
  "resources": "deployments,secrets,svc",
  "labels": "app=wordpress",
  "serverAddr": "10.15.170.48:50051",
  "readerAddr": "10.15.168.215:30051"
}' -plaintext localhost:50051 move.MoveService/MigrateApp
```
