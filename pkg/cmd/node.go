/*
Copyright © 2021 Microshift Contributors

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
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/openshift/microshift/pkg/node"
)

// nodeCmd represents the node command
var NodeCmd = &cobra.Command{
	Use:   "node",
	Short: "openshift node start",
	Long:  `openshift node start`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return startNode(args)
	},
}

func startNode(args []string) error {
	if err := startNodeOnly(); err != nil {
		return err
	}
	select {}
}

func startNodeOnly() error {
	if err := node.StartKubelet(); err != nil {
		return err
	}
	node.StartKubeProxy()
	return nil
}
