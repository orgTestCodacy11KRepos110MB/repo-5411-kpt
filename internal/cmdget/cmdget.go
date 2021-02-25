// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cmdget contains the get command
package cmdget

import (
	"fmt"

	docs "github.com/GoogleContainerTools/kpt/internal/docs/generated/pkgdocs"
	"github.com/GoogleContainerTools/kpt/internal/util/cmdutil"
	"github.com/GoogleContainerTools/kpt/internal/util/get"
	"github.com/GoogleContainerTools/kpt/internal/util/parse"
	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/kyaml/kio/filters"
)

// NewRunner returns a command runner
func NewRunner(parent string) *Runner {
	r := &Runner{}
	c := &cobra.Command{
		Use:        "get REPO_URI[.git]/PKG_PATH[@VERSION] LOCAL_DEST_DIRECTORY",
		Args:       cobra.ExactArgs(2),
		Short:      docs.GetShort,
		Long:       docs.GetShort + "\n" + docs.GetLong,
		Example:    docs.GetExamples,
		RunE:       r.runE,
		PreRunE:    r.preRunE,
		SuggestFor: []string{"clone", "cp", "fetch"},
	}
	cmdutil.FixDocs("kpt", parent, c)
	r.Command = c
	c.Flags().StringVar(&r.FilenamePattern, "pattern", filters.DefaultFilenamePattern,
		`Pattern to use for writing files.  
May contain the following formatting verbs
%n: metadata.name, %s: metadata.namespace, %k: kind
`)
	return r
}

func NewCommand(parent string) *cobra.Command {
	return NewRunner(parent).Command
}

// Runner contains the run function
type Runner struct {
	Get             get.Command
	Command         *cobra.Command
	FilenamePattern string
}

func (r *Runner) preRunE(_ *cobra.Command, args []string) error {
	t, err := parse.GitParseArgs(args)
	if err != nil {
		return err
	}
	r.Get.GitLock = t.GitLock

	_, absDestPath, err := cmdutil.ResolveAbsAndRelPaths(t.Destination)
	if err != nil {
		return err
	}

	r.Get.Destination = absDestPath
	return nil
}

func (r *Runner) runE(c *cobra.Command, args []string) error {
	fmt.Fprintf(c.OutOrStdout(), "fetching package %s from %s to %s\n",
		r.Get.Directory, r.Get.Repo, r.Get.Destination)
	if err := r.Get.Run(); err != nil {
		return err
	}

	return nil
}