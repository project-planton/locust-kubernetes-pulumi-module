package cluster

import (
	locustcontextstate "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/contextstate"
	locustkubernetesmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/locustkubernetes/model"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	resourceId            string
	resourceName          string
	namespaceName         string
	namespace             *kubernetescorev1.Namespace
	mainPyConfigMapName   pulumi.StringOutput
	libFilesConfigMapName pulumi.StringOutput
	loadTestName          string
	values                map[string]string
	workerContainerSpec   *locustkubernetesmodel.LocustKubernetesSpecContainerSpec
	masterContainerSpec   *locustkubernetesmodel.LocustKubernetesSpecContainerSpec
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxState = ctx.Value(locustcontextstate.Key).(locustcontextstate.ContextState)

	return &input{
		resourceId:            ctxState.Spec.ResourceId,
		resourceName:          ctxState.Spec.ResourceName,
		namespaceName:         ctxState.Spec.NamespaceName,
		namespace:             ctxState.Status.AddedResources.Namespace,
		mainPyConfigMapName:   ctxState.Status.AddedResources.AddedMainPyConfigMap.Metadata.Name().Elem(),
		libFilesConfigMapName: ctxState.Status.AddedResources.AddedLibFilesConfigMap.Metadata.Name().Elem(),
		loadTestName:          ctxState.Spec.LoadTest.Name,
		values:                ctxState.Spec.CustomHelmValues,
		workerContainerSpec:   ctxState.Spec.WorkerContainerSpec,
		masterContainerSpec:   ctxState.Spec.MasterContainerSpec,
	}
}
