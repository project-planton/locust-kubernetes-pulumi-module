package cluster

import (
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) error {
	err := addHelmChart(ctx)
	if err != nil {
		return err
	}
	return nil
}

func addHelmChart(ctx *pulumi.Context) error {
	i := extractInput(ctx)

	// Deploying a Locust Helm chart from the Helm repository.
	_, err := helmv3.NewChart(ctx, i.resourceId, helmv3.ChartArgs{
		Chart:     pulumi.String("locust"),
		Version:   pulumi.String("0.31.5"), // Use the Helm chart version you want to install
		Namespace: pulumi.String(i.namespaceName),
		Values:    getHelmChartValuesMap(i),
		//if you need to add the repository, you can specify `repo url`:
		FetchArgs: helmv3.FetchArgs{
			Repo: pulumi.String("https://charts.deliveryhero.io"), // The URL for the Helm chart repository
		},
	}, pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "3m", Update: "3m", Delete: "3m"}), pulumi.Parent(i.namespace))
	if err != nil {
		return err
	}
	return nil
}
