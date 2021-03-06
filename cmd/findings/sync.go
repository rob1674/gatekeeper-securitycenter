// Copyright 2020 Google LLC
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

package findings

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/GoogleCloudPlatform/gatekeeper-securitycenter/cmd/flag"
	"github.com/GoogleCloudPlatform/gatekeeper-securitycenter/pkg/logging"
	"github.com/GoogleCloudPlatform/gatekeeper-securitycenter/pkg/sync"
)

var (
	syncFlags = flag.New(googleServiceAccount, kubeconfig, dryRun, source, clusterName)

	syncCmd = &cobra.Command{
		Use:   "sync",
		Short: "Run a one-off sync",
		PreRunE: func(_ *cobra.Command, _ []string) error {
			return syncFlags.Validate()
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return syncRun(cmd.Context())
		},
	}
)

func init() {
	syncFlags.AddToFlagSet(syncCmd.Flags())
}

// syncRun runs a one-off sync
func syncRun(ctx context.Context) error {
	log := logging.CreateStdLog("sync")
	client, err := sync.NewClient(ctx, log, kubeconfig.Value(), dryRun.Value(), source.Value(), clusterName.Value(), googleServiceAccount.Value())
	if err != nil {
		return err
	}
	return client.Sync(ctx)
}
