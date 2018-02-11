/*
Copyright 2018 The Kubernetes Authors.

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

package framework

import (
	"testing"

	"github.com/kubernetes-sig-testing/frameworks/integration"
)

var (
	etcd     *integration.Etcd
	refCount int
)

func SetUpEtcd(t *testing.T) string {
	if etcd == nil {
		etcd = &integration.Etcd{}
		err := etcd.Start()
		if err != nil {
			etcd = nil
			t.Fatalf("Error starting etcd: %v", err)
		}
	}
	refCount += 1
	return etcd.URL.String()
}

func TearDownEtcd(t *testing.T) {
	if etcd != nil {
		refCount -= 1
		if refCount <= 0 {
			err := etcd.Stop()
			if err != nil {
				t.Errorf("Error stopping etcd: %v", err)
			}
			etcd = nil
		}
	}
}