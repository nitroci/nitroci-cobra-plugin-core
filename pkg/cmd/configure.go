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
	coreCbrPlugins "github.com/nitroci/nitroci-cobra-core/pkg/plugins"
	"github.com/nitroci/nitroci-cobra-plugin-core/pkg/plugins"

	"github.com/spf13/cobra"
)

var (
	configureFlagShow, configureFlagRaw bool
	configureFlags                      map[string]interface{}
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configurations management",
	Long:  `Configurations management`,
	Run: func(cmd *cobra.Command, args []string) {
		plugins.PluginModule.Configure(runtimeContext, args, configureFlags)
	},
}

func initConfigureConfig() {
	pluginModel, _ := corePlugins.LoadPluginFile("./manifest.yml")
	configureFlags = coreCbrPlugins.LoadMapFromFlags(configureCmd, pluginModel.Operations.Configure.Flags)
}

func init() {
	cobra.OnInitialize(initConfigureConfig)
	rootCmd.AddCommand(configureCmd)
	configureCmd.Flags().BoolVarP(&configureFlagShow, "show", "s", false, "show configurations")
	configureCmd.Flags().BoolVarP(&configureFlagRaw, "raw", "r", false, "output raw configurations")
	pluginModel, _ := corePlugins.LoadPluginFile("./manifest.yml")
	coreCbrPlugins.LoadFlags(configureCmd, pluginModel.Operations.Configure.Flags)
}
