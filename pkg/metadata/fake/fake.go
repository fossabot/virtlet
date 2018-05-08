/*
Copyright 2018 Mirantis

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

package fake

import (
	"strconv"

	"github.com/Mirantis/virtlet/pkg/metadata/types"
	"github.com/Mirantis/virtlet/pkg/utils"
)

const (
	samplePodNsUuid       = "cd6bc1b7-a9e2-4739-ade9-3e2447d28a90"
	sampleContainerNsUuid = "a51f5bed-db9c-49b1-a1b6-9989d46d637b"
)

type ContainerTestConfig struct {
	Name                string
	SandboxId           string
	ContainerId         string
	Image               string
	RootImageVolumeName string
	Labels              map[string]string
	Annotations         map[string]string
}

func GetSandboxes(sandboxCount int) []*types.PodSandboxConfig {
	sandboxes := []*types.PodSandboxConfig{}
	for i := 0; i < sandboxCount; i++ {
		name := "testName_" + strconv.Itoa(i)
		sandboxConfig := &types.PodSandboxConfig{
			Name:         name,
			Uid:          utils.NewUUID5(samplePodNsUuid, name),
			Namespace:    "default",
			Attempt:      uint32(0),
			Hostname:     "localhost",
			LogDirectory: "/var/log/test_log_directory",
			Labels: map[string]string{
				"foo":  "bar",
				"fizz": "buzz",
			},
			Annotations: map[string]string{
				"hello": "world",
				"virt":  "let",
			},
		}

		sandboxes = append(sandboxes, sandboxConfig)
	}

	return sandboxes
}

func GetContainersConfig(sandboxConfigs []*types.PodSandboxConfig) []*ContainerTestConfig {
	containers := []*ContainerTestConfig{}
	for _, sandbox := range sandboxConfigs {
		name := "container-for-" + sandbox.Name
		containerConf := &ContainerTestConfig{
			Name:                name,
			SandboxId:           sandbox.Uid,
			Image:               "testImage",
			RootImageVolumeName: "sample_name",
			ContainerId:         utils.NewUUID5(sampleContainerNsUuid, name),
			Labels:              map[string]string{"foo": "bar", "fizz": "buzz"},
			Annotations:         map[string]string{"hello": "world", "virt": "let"},
		}
		containers = append(containers, containerConf)
	}

	return containers
}
