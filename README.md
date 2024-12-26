# Kubernetes Application Migration Tool

## Overview  
This tool facilitates the migration of applications from one Kubernetes cluster to another by transferring persistent volumes and application configurations using Kubernetes APIs and GRPC.

---

## Prerequisites  

### Source Cluster  
- Ensure a valid **VolumeSnapshotClass** exists.

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

To deploy the Writer Server, apply the `writer-deployment.yaml` file to the **target Kubernetes cluster**:

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
   kubectl get service writer-server-service -n <namespace>
   ```

### Provide IPs to the Orchestrator

Use the IPs obtained above in the GRPC request sent to the orchestrator.

#### Example GRPC Command:

```bash
grpcurl -d '{"reader_ip": "source-cluster-ip:50051", "writer_ip": "target-cluster-ip:50051"}' -plaintext orchestrator-service-ip:port
```

---

## References

- [Kubernetes Documentation](https://kubernetes.io/docs/concepts)
- [Nutanix Data Services for Kubernetes](https://portal.nutanix.com/page/documents/details?targetId=Nutanix-Data-Services-for-Kubernetes-v1_2)
- [Learn Kubernetes - Udemy Course](https://nutanix.udemy.com/course/learn-kubernetes)
