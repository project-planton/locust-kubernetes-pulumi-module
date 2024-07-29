package pkg

import (
	"fmt"
	"github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/outputs"
	locustkubernetesmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/locustkubernetes/model"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	IngressCertClusterIssuerName string
	IngressCertSecretName        string
	IngressExternalHostname      string
	IngressHostnames             []string
	IngressInternalHostname      string
	KubePortForwardCommand       string
	KubeServiceFqdn              string
	KubeServiceName              string
	Namespace                    string
	LocustKubernetes             *locustkubernetesmodel.LocustKubernetes
}

func initializeLocals(ctx *pulumi.Context, stackInput *locustkubernetesmodel.LocustKubernetesStackInput) *Locals {
	locals := &Locals{}
	//assign value for the locals variable to make it available across the project
	locals.LocustKubernetes = stackInput.ApiResource

	locustKubernetes := stackInput.ApiResource

	//decide on the namespace
	locals.Namespace = locustKubernetes.Metadata.Id
	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))

	locals.KubeServiceName = locustKubernetes.Metadata.Name

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, locustKubernetes.Metadata.Name)

	//export kube-port-forward command
	ctx.Export(outputs.KubePortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if locustKubernetes.Spec.Ingress == nil ||
		!locustKubernetes.Spec.Ingress.IsEnabled ||
		locustKubernetes.Spec.Ingress.EndpointDomainName == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", locustKubernetes.Metadata.Id,
		locustKubernetes.Spec.Ingress.EndpointDomainName)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", locustKubernetes.Metadata.Id,
		locustKubernetes.Spec.Ingress.EndpointDomainName)

	locals.IngressHostnames = []string{
		locals.IngressExternalHostname,
		locals.IngressInternalHostname,
	}

	//export ingress hostnames
	ctx.Export(outputs.IngressExternalHostname, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(outputs.IngressInternalHostname, pulumi.String(locals.IngressInternalHostname))

	//note: a ClusterIssuer resource should have already exist on the kubernetes-cluster.
	//this is typically taken care of by the kubernetes cluster administrator.
	//if the kubernetes-cluster is created using Planton Cloud, then the cluster-issuer name will be
	//same as the ingress-domain-name as long as the same ingress-domain-name is added to the list of
	//ingress-domain-names for the GkeCluster/EksCluster/AksCluster spec.
	locals.IngressCertClusterIssuerName = locustKubernetes.Spec.Ingress.EndpointDomainName

	locals.IngressCertSecretName = locustKubernetes.Metadata.Id

	return locals
}