package gcp

import (
	"github.com/pkg/errors"
	locustkubernetescluster "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/cluster"
	locustconfigmap "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/configmap"
	locustcontextstate "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/contextstate"
	locustnamespace "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/namespace"
	locustnetwork "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/network"
	locustoutputs "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/outputs"
	locuststackmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/locustkubernetes/stack/model"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	WorkspaceDir     string
	Input            *locuststackmodel.LocustKubernetesStackInput
	KubernetesLabels map[string]string
}

func (resourceStack *ResourceStack) Resources(ctx *pulumi.Context) error {
	// reference to setup locust resources in locust https://github.com/deliveryhero/helm-charts/tree/master/stable/locust

	//load context config
	var ctxConfig, err = loadConfig(ctx, resourceStack)
	if err != nil {
		return errors.Wrap(err, "failed to initiate context config")
	}
	ctx = ctx.WithValue(locustcontextstate.Key, *ctxConfig)

	// Create the namespace resource
	ctx, err = locustnamespace.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to create namespace resource")
	}

	// Create a ConfigMap for the main.py file and lib files
	ctx, err = locustconfigmap.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to add config map")
	}

	// Deploying a Locust Helm chart from the Helm repository.
	err = locustkubernetescluster.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to add locust-kubernetes cluster resources")
	}

	ctx, err = locustnetwork.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to add locust-locust ingress resources")
	}

	err = locustoutputs.Export(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to export locust kubernetes outputs")
	}

	return nil
}
