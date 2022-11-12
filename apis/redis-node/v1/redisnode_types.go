/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RedisNodeSpec defines the desired state of RedisNode
type RedisNodeSpec struct {
	HostIP string `json:"host_ip,omitempty"`
}

// RedisNodeStatus defines the observed state of RedisNode
type RedisNodeStatus struct {
	Role                  string `json:"role,omitempty"`
	MemoryUsage           string `json:"memory_usage_human,omitempty"`
	CPUUsage              string `json:"used_cpu_user,omitempty"`
	NumOfConnectedClients string `json:"connected_client,omitempty"`
	StartSlot             string `json:"start_slot,omitempty"`
	EndSlot               string `json:"end_slot,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RedisNode is the Schema for the redisnodes API
type RedisNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisNodeSpec   `json:"spec,omitempty"`
	Status RedisNodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RedisNodeList contains a list of RedisNode
type RedisNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisNode{}, &RedisNodeList{})
}
