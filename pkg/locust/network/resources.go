package network

import (
	"github.com/pkg/errors"
	locustingress "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/network/ingress"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) (*pulumi.Context, error) {
	i := extractInput(ctx)
	if !i.isIngressEnabled || i.endpointDomainName == "" {
		return ctx, nil
	}
	if ctx, err := locustingress.Resources(ctx); err != nil {
		return ctx, errors.Wrap(err, "failed to add network resources")
	}
	return ctx, nil
}
