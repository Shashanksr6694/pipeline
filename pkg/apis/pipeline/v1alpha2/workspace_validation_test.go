/*
Copyright 2019 The Tekton Authors

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

package v1alpha2

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestWorkspaceBindingValidateValid(t *testing.T) {
	for _, tc := range []struct {
		name    string
		binding *WorkspaceBinding
	}{{
		name: "Valid PVC",
		binding: &WorkspaceBinding{
			Name: "beth",
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: "pool-party",
			},
		},
	}, {
		name: "Valid emptyDir",
		binding: &WorkspaceBinding{
			Name:     "beth",
			EmptyDir: &corev1.EmptyDirVolumeSource{},
		},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.binding.Validate(context.Background()); err != nil {
				t.Errorf("didnt expect error for valid binding but got: %v", err)
			}
		})
	}

}

func TestWorkspaceBindingValidateInvalid(t *testing.T) {
	for _, tc := range []struct {
		name    string
		binding *WorkspaceBinding
	}{{
		name:    "no binding provided",
		binding: nil,
	}, {
		name: "Provided both pvc and emptydir",
		binding: &WorkspaceBinding{
			Name:     "beth",
			EmptyDir: &corev1.EmptyDirVolumeSource{},
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: "pool-party",
			},
		},
	}, {
		name: "Provided neither pvc nor emptydir",
		binding: &WorkspaceBinding{
			Name: "beth",
		},
	}, {
		name: "Provided pvc without claim name",
		binding: &WorkspaceBinding{
			Name:                  "beth",
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{},
		},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.binding.Validate(context.Background()); err == nil {
				t.Errorf("expected error for invalid binding but didn't get any!")
			}
		})
	}
}
