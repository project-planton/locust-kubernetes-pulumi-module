package configmap

import (
	locustcontextstate "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/contextstate"
	locustkubernetesmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/locustkubernetes/model"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	loadTest      *locustkubernetesmodel.LocustKubernetesSpecLoadTestSpec
	namespace     *kubernetescorev1.Namespace
	namespaceName string
	labels        map[string]string
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxState = ctx.Value(locustcontextstate.Key).(locustcontextstate.ContextState)

	return &input{
		loadTest:      ctxState.Spec.LoadTest,
		namespace:     ctxState.Status.AddedResources.Namespace,
		namespaceName: ctxState.Spec.NamespaceName,
		labels:        ctxState.Spec.Labels,
	}
}
