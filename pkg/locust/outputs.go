package gcp

import (
	"context"
	"github.com/pkg/errors"
	"github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/outputs"
	"github.com/plantoncloud/stack-job-runner-golang-sdk/pkg/stack/output/backend"

	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/enums/stackjoboperationtype"

	locustkubernetesmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/locustkubernetes/model"
	locustkubernetesstackmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/locustkubernetes/stack/model"
)

func Outputs(ctx context.Context,
	input *locustkubernetesstackmodel.LocustKubernetesStackInput) (*locustkubernetesmodel.LocustKubernetesStatusStackOutputs, error) {
	stackOutput, err := backend.StackOutput(input.StackJob)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get stack output")
	}
	return OutputMapTransformer(stackOutput, input), nil
}

func OutputMapTransformer(stackOutput map[string]interface{},
	input *locustkubernetesstackmodel.LocustKubernetesStackInput) *locustkubernetesmodel.LocustKubernetesStatusStackOutputs {
	if input.StackJob.Spec.OperationType != stackjoboperationtype.StackJobOperationType_apply || stackOutput == nil {
		return &locustkubernetesmodel.LocustKubernetesStatusStackOutputs{}
	}
	return &locustkubernetesmodel.LocustKubernetesStatusStackOutputs{
		Namespace:          backend.GetVal(stackOutput, outputs.GetNamespaceNameOutputName()),
		Service:            backend.GetVal(stackOutput, outputs.GetKubeServiceNameOutputName()),
		PortForwardCommand: backend.GetVal(stackOutput, outputs.GetKubePortForwardCommandOutputName()),
		KubeEndpoint:       backend.GetVal(stackOutput, outputs.GetKubeEndpointOutputName()),
		ExternalHostname:   backend.GetVal(stackOutput, outputs.GetExternalClusterHostnameOutputName()),
		InternalHostname:   backend.GetVal(stackOutput, outputs.GetInternalClusterHostnameOutputName()),
	}
}
