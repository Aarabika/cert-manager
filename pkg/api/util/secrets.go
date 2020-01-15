package util

import "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1"

func GetSecretsNamespace(crt *v1alpha1.Certificate) string {
	namespace, ok := crt.Annotations[v1alpha1.SecretsNamespaceAnnotationKey]

	if !ok {
		return crt.Namespace
	}

	return namespace
}
