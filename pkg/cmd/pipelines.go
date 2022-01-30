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
	pkgCCPPlugins "github.com/nitroci/nitroci-cobra-plugin-core/pkg/core/plugins"
	pkgCPlugins "github.com/nitroci/nitroci-core/pkg/core/plugins"

	"github.com/spf13/cobra"
)

var (
	pipelinesFlags map[string]interface{}
)

var pipelinesCmd = &cobra.Command{
	Use:   "pipelines",
	Short: "Interact with the pipelines",
	Long:  `Interact with the pipelines`,
	Run: func(cmd *cobra.Command, args []string) {
		pkgCCPPlugins.PluginModule.Pipelines(runtimeContext, args, pipelinesFlags)
	},
}

func initPipelinesConfig() {
	pluginModel, _ := pkgCPlugins.LoadPluginFile("./manifest.yml")
	pipelinesFlags = pkgCCPlugins.LoadMapFromFlags(pipelinesCmd, pluginModel.Operations.Pipelines.Flags)
}

func init() {
	cobra.OnInitialize(initPipelinesConfig)
	rootCmd.AddCommand(pipelinesCmd)
	pluginModel, _ := pkgCPlugins.LoadPluginFile("./manifest.yml")
	pkgCCPlugins.LoadFlags(pipelinesCmd, pluginModel.Operations.Pipelines.Flags)
}
