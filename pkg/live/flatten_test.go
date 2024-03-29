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

package live

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Test_Flatten(t *testing.T) {
	tests := []struct {
		name string
		in   []*unstructured.Unstructured
		want []*unstructured.Unstructured
	}{
		{
			name: "simple list",
			in:   simpleList,
			want: simpleList,
		},
		{
			name: "list with list",
			in: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"kind":  "List",
						"items": anySlice(simpleList),
					},
				},
			},
			want: simpleList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Flatten(tt.in)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func anySlice(in []*unstructured.Unstructured) (out []interface{}) {
	for _, o := range in {
		out = append(out, o.Object)
	}
	return
}

var simpleList = []*unstructured.Unstructured{
	{
		Object: map[string]interface{}{
			"kind":    "ConfigMap",
			"version": "v1",
		},
	},
	{
		Object: map[string]interface{}{
			"kind":    "Pod",
			"version": "v1",
		},
	},
}
