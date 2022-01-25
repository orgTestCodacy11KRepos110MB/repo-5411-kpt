// Copyright 2022 Google LLC
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

package kptlib

import (
	"io"

	"github.com/GoogleContainerTools/kpt/internal/pkg"
	kptfilev1 "github.com/GoogleContainerTools/kpt/pkg/api/kptfile/v1"
	"github.com/GoogleContainerTools/kpt/pkg/kptfile/kptfileutil"
)

func ParseKptFile(r io.Reader) (*kptfilev1.KptFile, error) {
	kf, err := pkg.DecodeKptfile(r)
	if err != nil {
		return nil, err
	}
	return kf, nil
}

func UpdateUpstreamLockFromGit(path string, spec *GitRepoSpec) error {
	return kptfileutil.UpdateUpstreamLockFromGit(path, spec)
}
