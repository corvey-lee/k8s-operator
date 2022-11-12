<h1>K8S Custom Controller</h1>

As pre-requisite, operator-sdk need initial go modules.

```go 
go mod init k8s-operator
go mod tidy
```

```shell
operator-sdk init --domain k8s-operator.com
operator-sdk create api --group redis-cluster --version v1 --kind RedisCluster --resource --controller
make manifests
```

``make manifests`` should be done after modifying the ``*_types.go`` file

```go 
type RedisClusterSpec struct {
	Image            string   `json:"image,omitempty"`
	MasterNodeIPList []string `json:"master_node_ip_list,omitempty"`
}

type RedisClusterStatus struct {
	ClusterState    string `json:"cluster_state,omitempty"`
	NumOfRedisNodes string `json:"cluster_size,omitempty"`
}
```

Following ``yaml`` file was generated through ``make manifests``

```yaml
spec:
  description: RedisClusterSpec defines the desired state of RedisCluster
  properties:
    image:
      type: string
    master_node_ip_list:
      items:
        type: string
      type: array
  type: object
status:
  description: RedisClusterStatus defines the observed state of RedisCluster
  properties:
    cluster_size:
      type: string
    cluster_state:
      description: 'INSERT ADDITIONAL STATUS FIELD - define observed state
        of cluster Important: Run "make" to regenerate code after modifying
        this file'
      type: string
  type: object
```

To manage multi CR(``RedisCluster``, ``RedisNode``) through one Custom Controller
```shell
operator-sdk edit --multigroup=true
operator-sdk create api --group redis-node --version v1 --kind RedisNode --resource --controller
```


