package ingress

import (
	"github.com/pkg/errors"
	locustistio "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/network/ingress/istio"
	locustloadbalancer "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/network/ingress/loadbalancer"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) (newCtx *pulumi.Context, err error) {
	i := extractInput(ctx)
	switch i.ingressType {
	case kubernetesworkloadingresstype.KubernetesWorkloadIngressType_load_balancer:
		ctx, err = locustloadbalancer.Resources(ctx)
		if err != nil {
			return ctx, errors.Wrap(err, "failed to add load balancer resources")
		}
	case kubernetesworkloadingresstype.KubernetesWorkloadIngressType_ingress_controller:
		if err = locustistio.Resources(ctx); err != nil {
			return ctx, errors.Wrap(err, "failed to add istio resources")
		}
	}
	return ctx, nil
}
