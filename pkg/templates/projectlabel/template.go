package projectlabel

import (
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/projectlabel/internal/params"
)

const (
	labelname = "project"
)

func init() {
	templates.Register(check.Template{
		HumanName:   "Project Label on Deployments",
		Key:         "project-label",
		Description: "Flag Deployments not having a project label",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{objectkinds.DeploymentLike},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(_ params.Params) (check.Func, error) {
			return func(_ lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
				labels := object.K8sObject.GetLabels()

				if containsKey(labels, labelname) {
					return []diagnostic.Diagnostic{{Message: "object missing project label"}}
				}
				return nil
			}, nil
		}),
	})
}

func containsKey[K comparable, V any](m map[K]V, key K) bool {
	for k := range m {
		if k == key {
			return true
		}
	}
	return false
}
