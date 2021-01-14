package main

import (
	"fmt"

	kopsapi "k8s.io/kops/pkg/apis/kops"
)

var DefaultCluster = kopsapi.Cluster{}

func main() {
	fmt.Printf("cluster %v\n", DefaultCluster)
}
