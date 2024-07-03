package gcp

import (
	"github.com/pkg/errors"
	environmentblueprinthostnames "github.com/plantoncloud/environment-pulumi-blueprint/pkg/gcpgke/endpointdomains/hostnames"
	locustcontextstate "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/contextstate"
	locustnetutilshostname "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/network/ingress/netutils/hostname"
	locustnetutilsservice "github.com/plantoncloud/locust-kubernetes-pulumi-blueprint/pkg/locust/network/ingress/netutils/service"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	"github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func loadConfig(ctx *pulumi.Context, resourceStack *ResourceStack) (*locustcontextstate.ContextState, error) {

	kubernetesProvider, err := pulumikubernetesprovider.GetWithStackCredentials(ctx, resourceStack.Input.CredentialsInput)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup kubernetes provider")
	}

	var resourceId = resourceStack.Input.ResourceInput.Metadata.Id
	var resourceName = resourceStack.Input.ResourceInput.Metadata.Name
	var environmentInfo = resourceStack.Input.ResourceInput.Spec.EnvironmentInfo
	var isIngressEnabled = false

	if resourceStack.Input.ResourceInput.Spec.Ingress != nil {
		isIngressEnabled = resourceStack.Input.ResourceInput.Spec.Ingress.IsEnabled
	}

	var endpointDomainName = ""
	var envDomainName = ""
	var ingressType = kubernetesworkloadingresstype.KubernetesWorkloadIngressType_unspecified
	var internalHostname = ""
	var externalHostname = ""
	var certSecretName = ""

	if isIngressEnabled {
		endpointDomainName = resourceStack.Input.ResourceInput.Spec.Ingress.EndpointDomainName
		envDomainName = environmentblueprinthostnames.GetExternalEnvHostname(environmentInfo.EnvironmentName, endpointDomainName)
		ingressType = resourceStack.Input.ResourceInput.Spec.Ingress.IngressType

		internalHostname = locustnetutilshostname.GetInternalHostname(resourceId, environmentInfo.EnvironmentName, endpointDomainName)
		externalHostname = locustnetutilshostname.GetExternalHostname(resourceId, environmentInfo.EnvironmentName, endpointDomainName)
	}

	return &locustcontextstate.ContextState{
		Spec: &locustcontextstate.Spec{
			KubeProvider:        kubernetesProvider,
			ResourceId:          resourceId,
			ResourceName:        resourceName,
			MasterContainerSpec: resourceStack.Input.ResourceInput.Spec.MasterContainer,
			WorkerContainerSpec: resourceStack.Input.ResourceInput.Spec.WorkerContainer,
			CustomHelmValues:    resourceStack.Input.ResourceInput.Spec.HelmValues,
			LoadTest:            resourceStack.Input.ResourceInput.Spec.LoadTest,
			Labels:              resourceStack.KubernetesLabels,
			WorkspaceDir:        resourceStack.WorkspaceDir,
			NamespaceName:       resourceId,
			EnvironmentInfo:     resourceStack.Input.ResourceInput.Spec.EnvironmentInfo,
			IsIngressEnabled:    isIngressEnabled,
			IngressType:         ingressType,
			EndpointDomainName:  endpointDomainName,
			EnvDomainName:       envDomainName,
			InternalHostname:    internalHostname,
			ExternalHostname:    externalHostname,
			KubeServiceName:     locustnetutilsservice.GetKubeServiceName(resourceName),
			KubeLocalEndpoint:   locustnetutilsservice.GetKubeServiceNameFqdn(resourceName, resourceId),
			CertSecretName:      certSecretName,
		},
		Status: &locustcontextstate.Status{},
	}, nil
}