/*
Copyright 2021 The NitroCI Authors.

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
	pkgCCPlugins "github.com/nitroci/nitroci-cobra-core/pkg/core/plugins"
	pkg "github.com/nitroci/nitroci-cobra-plugin-core/pkg/core/plugins"
	pkgCPlugins "github.com/nitroci/nitroci-core/pkg/core/plugins"

	"github.com/spf13/cobra"
)

var (
	environmentsFlags map[string]interface{}
)

var environmentsCmd = &cobra.Command{
	Use:   "environments",
	Short: "Interact with the environments",
	Long:  `Interact with the environments`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.PluginModule.Environments(runtimeContext, args, environmentsFlags)
	},
}

func initEnvironmentsConfig() {
	pluginModel, _ := pkgCPlugins.LoadPluginFile("./manifest.yml")
	environmentsFlags = pkgCCPlugins.LoadMapFromFlags(environmentsCmd, pluginModel.Operations.Environments.Flags)
}

func init() {
	cobra.OnInitialize(initEnvironmentsConfig)
	rootCmd.AddCommand(environmentsCmd)
	pluginModel, _ := pkgCPlugins.LoadPluginFile("./manifest.yml")
	pkgCCPlugins.LoadFlags(environmentsCmd, pluginModel.Operations.Environments.Flags)
}
