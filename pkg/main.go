package main

import (
	"fmt"

	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

var DefaultCluster = kops.Cluster{}

func main() {
	fmt.Printf("cluster %v. version %v\n", DefaultCluster, cloudup.OldestRecommendedKubernetesVersion)
}
