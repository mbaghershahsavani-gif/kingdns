package controller

type KubernetesResource struct {
	Name      string
	Namespace string
}

func DiscoverService(name string, namespace string) KubernetesResource {
	return KubernetesResource{
		Name:      name,
		Namespace: namespace,
	}
}
