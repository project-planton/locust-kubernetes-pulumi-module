package ingress

import (
	locustcontextstate "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/contextstate"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	ingressType kubernetesworkloadingresstype.KubernetesWorkloadIngressType
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxConfig = ctx.Value(locustcontextstate.Key).(locustcontextstate.ContextState)

	return &input{
		ingressType: ctxConfig.Spec.IngressType,
	}
}
