/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package cmd

import (
	"github.com/spf13/cobra"
)

// startCmd is starting node
var startCmd = &cobra.Command{
	Use:    "start",
	Short:  "Starting node",
	PreRun: loadConfigWKey,
	RunE: func(cmd *cobra.Command, args []string) error {
		return loadStartRun()
	},
}
