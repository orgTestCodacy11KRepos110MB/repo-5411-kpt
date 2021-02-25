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

package cmdsearch

var searchReplaceCases = []test{
	{
		name: "search by value",
		args: []string{"--by-value", "3"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo: 3
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.replicas
value: 3

${baseDir}/${filePath}
fieldPath: spec.foo
value: 3

Matched 2 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo: 3
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
 `,
	},
	{
		name: "search replace by value",
		args: []string{"--by-value", "3", "--put-value", "4"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
foo:
  bar: 3
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.replicas
value: 4

${baseDir}/${filePath}
fieldPath: foo.bar
value: 4

Mutated 2 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 4
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
foo:
  bar: 4
 `,
	},
	{
		name: "search replace by value to different type 1",
		args: []string{"--by-value", "3", "--put-value", "four"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
foo:
  bar: 3
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.replicas
value: four

${baseDir}/${filePath}
fieldPath: foo.bar
value: four

Mutated 2 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: four
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
foo:
  bar: four
 `,
	},
	{
		name: "search replace by value to different type 2",
		args: []string{"--by-value", "four", "--put-value", "4"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo: four
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.foo
value: 4

Mutated 1 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo: 4
 `,
	},
	{
		name: "search replace multiple deployments",
		args: []string{"--by-value", "3", "--put-value", "4"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql-deployment
spec:
  replicas: 3
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.replicas
value: 4

${baseDir}/${filePath}
fieldPath: spec.replicas
value: 4

Mutated 2 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 4
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql-deployment
spec:
  replicas: 4
 `,
	},
	{
		name: "search replace multiple deployments different value",
		args: []string{"--by-value", "3", "--put-value", "4"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql-deployment
spec:
  replicas: 5
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.replicas
value: 4

Mutated 1 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 4
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql-deployment
spec:
  replicas: 5
 `,
	},
	{
		name: "search by regex",
		args: []string{"--by-value-regex", "nginx-*"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
 `,
		out: `${baseDir}/${filePath}
fieldPath: metadata.name
value: nginx-deployment

Matched 1 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
 `,
	},
	{
		name: "search replace by regex",
		args: []string{"--by-value-regex", "nginx-*", "--put-value", "ubuntu-deployment"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
 `,
		out: `${baseDir}/${filePath}
fieldPath: metadata.name
value: ubuntu-deployment

Mutated 1 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ubuntu-deployment
spec:
  replicas: 3
 `,
	},
	{
		name: "search replace by regex helm template and empty values",
		args: []string{"--by-value-regex", "nginx-*", "--put-value", "ubuntu-deployment"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: {? {replicas: ''} : ''}
  foo:
 `,
		out: `${baseDir}/${filePath}
fieldPath: metadata.name
value: ubuntu-deployment

Mutated 1 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ubuntu-deployment
spec:
  replicas: {? {replicas: ''} : ''}
  foo:
 `,
	},
	{
		name: "search by path",
		args: []string{"--by-path", "spec.replicas"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo: 3
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.replicas
value: 3

Matched 1 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo: 3
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
 `,
	},
	{
		name: "search by array path",
		args: []string{"--by-path", "spec.foo[1]"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo:
  - a
  - b
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.foo[1]
value: b

Matched 1 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo:
  - a
  - b
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
 `,
	},
	{
		name: "search by array path all elements",
		args: []string{"--by-path", "spec.foo[*]"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo:
  - a
  - b
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.foo[0]
value: a

${baseDir}/${filePath}
fieldPath: spec.foo[1]
value: b

Matched 2 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo:
  - a
  - b
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
 `,
	},
	{
		name: "search replace by array path regex",
		args: []string{"--by-path", "spec.foo[1]", "--put-value", "c"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo:
  - a
  - b
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.foo[1]
value: c

Mutated 1 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo:
  - a
  - c
 `,
	},
	{
		name: "search replace by array path out of bounds",
		args: []string{"--by-path", "spec.foo[2]", "--put-value", "c"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo:
  - a
  - b
 `,
		out: `Mutated 0 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo:
  - a
  - b
 `,
	},
	{
		name: "search replace by array objects path",
		args: []string{"--by-path", "spec.foo[1].c", "--put-value", "thing-new"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo:
  - c: thing0
  - c: thing1
  - c: thing2
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.foo[1].c
value: thing-new

Mutated 1 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo:
  - c: thing0
  - c: thing-new
  - c: thing2
 `,
	},
	{
		name: "replace by path and value",
		args: []string{"--by-path", "spec.replicas", "--by-value", "3", "--put-value", "4"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  foo: 3
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
 `,
		out: `${baseDir}/${filePath}
fieldPath: spec.replicas
value: 4

Mutated 1 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 4
  foo: 3
---
apiVersion: apps/v1
kind: Service
metadata:
  name: nginx-service
 `,
	},
	{
		name: "add non-existing field",
		args: []string{"--by-path", "metadata.namespace", "--put-value", "myspace"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
 `,
		out: `${baseDir}/${filePath}
fieldPath: metadata.namespace
value: myspace

Mutated 1 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  namespace: myspace
spec:
  replicas: 3
 `,
	},
	{
		name: "put literal error",
		args: []string{"--put-value", "something"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
 `,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
 `,
		errMsg: `at least one of ["by-value", "by-value-regex", "by-path"] must be provided`,
	},
	{
		name: "put value by regex capture groups",
		args: []string{"--by-value-regex", `(\w+)-dev-(\w+)-us-east-1-(\w+)`,
			"--put-value", "${1}-prod-${2}-us-central-1-${3}"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: foo1-dev-bar1-us-east-1-baz1
  namespace: foo2-dev-bar2-us-east-1-baz2
 `,
		out: `${baseDir}/${filePath}
fieldPath: metadata.name
value: foo1-prod-bar1-us-central-1-baz1

${baseDir}/${filePath}
fieldPath: metadata.namespace
value: foo2-prod-bar2-us-central-1-baz2

Mutated 2 field(s)
`,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: foo1-prod-bar1-us-central-1-baz1
  namespace: foo2-prod-bar2-us-central-1-baz2
 `,
	},
	{
		name: "error when both by-value and by-regex provided",
		args: []string{"--by-value", "nginx-deployment", "--by-value-regex", "nginx-*"},
		input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
 `,
		expectedResources: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
 `,
		errMsg: `only one of ["by-value", "by-value-regex"] can be provided`,
	},
}