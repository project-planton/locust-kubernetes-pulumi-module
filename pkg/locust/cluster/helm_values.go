package cluster

import (
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/kubernetes/containerresources"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/kubernetes/helm/mergemaps"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func getHelmChartValuesMap(i *input) pulumi.Map {

	//https://github.com/deliveryhero/helm-charts/tree/master/stable/locust#values
	var baseValues = pulumi.Map{
		"fullnameOverride": pulumi.String(i.resourceName),
		"master": pulumi.Map{
			"replicas":  pulumi.Int(i.masterContainerSpec.Replicas),
			"resources": containerresources.ConvertToPulumiMap(i.masterContainerSpec.Resources),
		},
		"worker": pulumi.Map{
			"replicas":  pulumi.Int(i.workerContainerSpec.Replicas),
			"resources": containerresources.ConvertToPulumiMap(i.workerContainerSpec.Resources),
		},
		"loadtest": pulumi.Map{
			"name":                        pulumi.String(i.loadTestName),
			"locust_locustfile_configmap": i.mainPyConfigMapName,
			"locust_lib_configmap":        i.libFilesConfigMapName,
		},
	}
	mergemaps.MergeMapToPulumiMap(baseValues, i.values)
	return baseValues
}
