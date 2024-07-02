package loadbalancer

import (
	locustcontextstate "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/contextstate"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/cloudaccount/enums/kubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	kubeProvider kubernetesprovider.KubernetesProvider
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxConfig = ctx.Value(locustcontextstate.Key).(locustcontextstate.ContextState)

	return &input{
		kubeProvider: ctxConfig.Spec.EnvironmentInfo.KubernetesProvider,
	}
}
