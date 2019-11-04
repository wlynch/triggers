// +build tools

package tools

import (
	_ "github.com/tektoncd/plumbing"
	_ "k8s.io/code-generator/pkg/util"
	_ "knative.dev/pkg/codegen/cmd/injection-gen"
	_ "knative.dev/pkg/version"
)
