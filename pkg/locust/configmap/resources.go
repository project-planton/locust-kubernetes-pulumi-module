package configmap

import (
	"github.com/pkg/errors"
	locustcontextstate "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/contextstate"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) (*pulumi.Context, error) {
	addedMainPyConfigMap, err := addMainPyConfigMap(ctx)
	if err != nil {
		return ctx, errors.Wrap(err, "failed to add main-py config map")
	}
	addedLibFilesConfigMap, err := addLibFilesConfigMap(ctx)
	if err != nil {
		return ctx, errors.Wrap(err, "failed to add lib-files config map")
	}

	var ctxState = ctx.Value(locustcontextstate.Key).(locustcontextstate.ContextState)
	addMainPyConfigMapToContext(&ctxState, addedMainPyConfigMap)
	addLibFileConfigMapToContext(&ctxState, addedLibFilesConfigMap)
	ctx = ctx.WithValue(locustcontextstate.Key, ctxState)
	return ctx, nil
}

func addMainPyConfigMap(ctx *pulumi.Context) (*kubernetescorev1.ConfigMap, error) {
	i := extractInput(ctx)

	// Create a ConfigMap for the main.py file
	addedMainPyConfigMap, err := kubernetescorev1.NewConfigMap(ctx, "main-py", &kubernetescorev1.ConfigMapArgs{
		Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
			Name:      pulumi.String("main-py"),
			Namespace: pulumi.String(i.namespaceName),
			Labels:    pulumi.ToStringMap(i.labels),
		}),
		Data: pulumi.StringMap{
			"main.py": pulumi.String(i.loadTest.MainPyContent),
		},
	}, pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "3m", Update: "3m", Delete: "3m"}),
		pulumi.Parent(i.namespace))
	if err != nil {
		return nil, err
	}
	return addedMainPyConfigMap, err
}

func addLibFilesConfigMap(ctx *pulumi.Context) (*kubernetescorev1.ConfigMap, error) {
	i := extractInput(ctx)

	addedLibFilesConfigMap, err := kubernetescorev1.NewConfigMap(ctx, "lib-files", &kubernetescorev1.ConfigMapArgs{
		Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
			Name:      pulumi.String("lib-files"),
			Namespace: pulumi.String(i.namespaceName),
			Labels:    pulumi.ToStringMap(i.labels),
		}),
		Data: pulumi.ToStringMap(i.loadTest.LibFilesContent),
	}, pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "3m", Update: "3m", Delete: "3m"}),
		pulumi.Parent(i.namespace))

	if err != nil {
		return nil, err
	}
	return addedLibFilesConfigMap, err
}

func addMainPyConfigMapToContext(existingConfig *locustcontextstate.ContextState, mainPyConfigMap *kubernetescorev1.ConfigMap) {
	if existingConfig.Status.AddedResources == nil {
		existingConfig.Status.AddedResources = &locustcontextstate.AddedResources{
			AddedMainPyConfigMap: mainPyConfigMap,
		}
		return
	}
	existingConfig.Status.AddedResources.AddedMainPyConfigMap = mainPyConfigMap
}

func addLibFileConfigMapToContext(existingConfig *locustcontextstate.ContextState, libFilesConfigMap *kubernetescorev1.ConfigMap) {
	if existingConfig.Status.AddedResources == nil {
		existingConfig.Status.AddedResources = &locustcontextstate.AddedResources{
			AddedLibFilesConfigMap: libFilesConfigMap,
		}
		return
	}
	existingConfig.Status.AddedResources.AddedLibFilesConfigMap = libFilesConfigMap
}
