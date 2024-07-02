package network

import (
	locustcontextstate "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/contextstate"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	isIngressEnabled   bool
	endpointDomainName string
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxConfig = ctx.Value(locustcontextstate.Key).(locustcontextstate.ContextState)

	return &input{
		isIngressEnabled:   ctxConfig.Spec.IsIngressEnabled,
		endpointDomainName: ctxConfig.Spec.EndpointDomainName,
	}
}
