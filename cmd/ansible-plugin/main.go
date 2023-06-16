// Copyright 2019 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/kubebuilder/v3/pkg/cli"
	cfgv3 "sigs.k8s.io/kubebuilder/v3/pkg/config/v3"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
	v2alpha "sigs.k8s.io/kubebuilder/v3/pkg/plugins/common/kustomize/v2-alpha"

	ansiblev1 "github.com/everettraven/ansible-external-poc/internal/plugins/ansible/v1"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	c := GetPluginsCLIAndRoot()
	return c.Run()
}

// GetPluginsCLIAndRoot returns the plugins based CLI configured to use operator-sdk as the root command
// This CLI can run kubebuilder commands and certain SDK specific commands that are aligned for
// the kubebuilder project layout
func GetPluginsCLIAndRoot() *cli.CLI {
	ansibleBundle, _ := plugin.NewBundle("ansible", plugin.Version{Number: 1}, v2alpha.Plugin{}, ansiblev1.Plugin{})
	c, err := cli.New(
		cli.WithCommandName("ansible"),
		cli.WithVersion("v1"),
		cli.WithPlugins(
			ansibleBundle,
		),
		cli.WithDefaultPlugins(cfgv3.Version, ansibleBundle),
		cli.WithDefaultProjectVersion(cfgv3.Version),
		cli.WithCompletion(),
	)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
