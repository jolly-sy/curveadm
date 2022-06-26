/*
 *  Copyright (c) 2021 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: CurveAdm
 * Created Date: 2021-10-15
 * Author: Jingli Chen (Wine93)
 */

package command

import (
	"fmt"

	"github.com/opencurve/curveadm/cli/command/playground"
	"github.com/opencurve/curveadm/cli/cli"
	"github.com/opencurve/curveadm/cli/command/cluster"
	"github.com/opencurve/curveadm/cli/command/config"
	"github.com/opencurve/curveadm/cli/command/plugin"
	"github.com/opencurve/curveadm/cli/command/target"
	"github.com/opencurve/curveadm/internal/tools"
	cliutil "github.com/opencurve/curveadm/internal/utils"
	"github.com/spf13/cobra"
)

var curveadmExample = `Examples:
  $ curveadm cluster add c1      # Add a cluster named 'c1'
  $ curveadm deploy              # Deploy current cluster
  $ curveadm stop                # Stop current cluster service
  $ curveadm clean               # Clean current cluster
  $ curveadm enter 6ff561598c6f  # Enter specified service container
  $ curveadm -u                  # Upgrade curveadm itself to the latest version`

type rootOptions struct {
	upgrade bool
}

func addSubCommands(cmd *cobra.Command, curveadm *cli.CurveAdm) {
	cmd.AddCommand(
		cluster.NewClusterCommand(curveadm),       // curveadm cluster ...
		config.NewConfigCommand(curveadm),         // curveadm config ...
		target.NewTargetCommand(curveadm),         // curveadm target ...
		plugin.NewPluginCommand(curveadm),         // curveadm plugin ...
		playground.NewPlaygroundCommand(curveadm), // curveadm playground ...

		NewDeployCommand(curveadm),   // curveadm deploy
		NewStartCommand(curveadm),    // curveadm start
		NewStopCommand(curveadm),     // curveadm stop
		NewRestartCommand(curveadm),  // curveadm restart
		NewReloadCommand(curveadm),   // curveadm reload
		NewStatusCommand(curveadm),   // curveadm status
		NewCleanCommand(curveadm),    // curveadm clean
		NewUpgradeCommand(curveadm),  // curveadm upgrade
		NewScaleOutCommand(curveadm), // curveadm scale-out
		NewMigrateCommand(curveadm),  // curveadm migrate
		NewEnterCommand(curveadm),    // curveadm enter
		NewMountCommand(curveadm),    // curveadm mount
		NewUmountCommand(curveadm),   // curveadm umount
		NewCheckCommand(curveadm),    // curveadm check
		NewSupportCommand(curveadm),  // curveadm support
		NewFormatCommand(curveadm),   // curveadm format
		NewMapCommand(curveadm),      // curveadm map
		NewUnmapCommand(curveadm),    // curveadm unmap
		NewAuditCommand(curveadm),    // curveadm audit

		NewCompletionCommand(), // curveadm completion
	)
}

func setupRootCommand(cmd *cobra.Command) {
	cmd.SetVersionTemplate("CurveAdm v{{.Version}}\n")
	cliutil.SetFlagErrorFunc(cmd)
	cliutil.SetHelpTemplate(cmd)
	cliutil.SetUsageTemplate(cmd)
}

func NewCurveAdmCommand(curveadm *cli.CurveAdm) *cobra.Command {
	var options rootOptions

	cmd := &cobra.Command{
		Use:     "curveadm [OPTIONS] COMMAND [ARGS...]",
		Short:   "Deploy and manage CurveBS/CurveFS cluster",
		Version: cli.Version,
		Example: curveadmExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if options.upgrade {
				return tools.Upgrade(curveadm)
			} else if len(args) == 0 {
				return cliutil.ShowHelp(curveadm.Err())(cmd, args)
			}

			return fmt.Errorf("curveadm: '%s' is not a curveadm command.\n"+
				"See 'curveadm --help'", args[0])
		},
		SilenceUsage:          true, // silence usage when an error occurs
		DisableFlagsInUseLine: true,
	}

	cmd.Flags().BoolP("version", "v", false, "Print version information and quit")
	cmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	cmd.Flags().BoolVarP(&options.upgrade, "upgrade", "u", false, "Upgrade curveadm itself to the latest version")

	addSubCommands(cmd, curveadm)
	setupRootCommand(cmd)

	return cmd
}
