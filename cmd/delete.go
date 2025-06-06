// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/spf13/cobra"
)

// NewCmdDelete creates a new cobra.Command for the delete subcommand.
func NewCmdDelete(options *[]crane.Option) *cobra.Command {
	return &cobra.Command{
		Use:   "delete IMAGE",
		Short: "Delete an image reference from its registry",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			digest, err := crane.Digest(args[0], *options...)
			if err != nil {
				return err
			}
			parsed, err := name.ParseReference(args[0])
			if err != nil {
				return err
			}
			ref := fmt.Sprintf("%s/%s@%s",
				parsed.Context().Registry.Name(),
				parsed.Context().RepositoryStr(), digest)
			return crane.Delete(ref, *options...)
		},
	}
}
