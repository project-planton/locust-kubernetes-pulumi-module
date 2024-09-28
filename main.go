package main

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/locust-kubernetes-pulumi-module/pkg"
	locustkubernetesmodel "github.com/plantoncloud/project-planton/apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/locustkubernetes"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &locustkubernetesmodel.LocustKubernetesStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
