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
	corePlugins "github.com/nitroci/nitroci-core/pkg/core/plugins"
	"github.com/nitroci/nitroci-cobra-plugin-core/pkg/plugins"

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
		plugins.PluginModule.Pipelines(runtimeContext, args, pipelinesFlags)
	},
}

func initPipelinesConfig() {
	pluginModel, _ := intPlugins.LoadPluginFile("./manifest.yml")
	pipelinesFlags = LoadMapFromFlags(pipelinesCmd, pluginModel.Operations.Pipelines.Flags)
}

func init() {
	cobra.OnInitialize(initPipelinesConfig)
	rootCmd.AddCommand(pipelinesCmd)
	pluginModel, _ := intPlugins.LoadPluginFile("./manifest.yml")
	LoadFlags(pipelinesCmd, pluginModel.Operations.Pipelines.Flags)
}
