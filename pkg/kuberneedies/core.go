// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

// the premier package for all your kuberneedies
package kuberneedies

import (
	"k8s.io/client-go/kubernetes"
)

var (
	client kubernetes.Interface
)

// InitializeCoreClient should be called once from each exe entry point of the controller or test suite
func InitializeCoreClient(cc kubernetes.Interface) {
	client = cc
}

// CoreClient retrieves the global state set up in InitializeCoreClient
func CoreClient() kubernetes.Interface {
	if client == nil {
		panic("you shouldnta done that")
	}

	return client
}
