package gateway

import (
	locustcontextstate "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/contextstate"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	workspaceDir     string
	kubeProvider     *pulumikubernetes.Provider
	resourceName     string
	resourceId       string
	labels           map[string]string
	externalHostname string
	envDomainName    string
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxState = ctx.Value(locustcontextstate.Key).(locustcontextstate.ContextState)

	return &input{
		workspaceDir:     ctxState.Spec.WorkspaceDir,
		kubeProvider:     ctxState.Spec.KubeProvider,
		resourceName:     ctxState.Spec.ResourceName,
		resourceId:       ctxState.Spec.ResourceId,
		labels:           ctxState.Spec.Labels,
		externalHostname: ctxState.Spec.ExternalHostname,
		envDomainName:    ctxState.Spec.EnvDomainName,
	}
}
