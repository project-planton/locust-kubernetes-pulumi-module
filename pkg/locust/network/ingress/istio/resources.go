package istio

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/network/ingress/istio/gateway"
	"github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/network/ingress/istio/virtualservice"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) error {
	if err := gateway.Resources(ctx); err != nil {
		return errors.Wrap(err, "failed to add gateway resources")
	}
	if err := virtualservice.Resources(ctx); err != nil {
		return errors.Wrap(err, "failed to add virtual resources")
	}
	return nil
}
