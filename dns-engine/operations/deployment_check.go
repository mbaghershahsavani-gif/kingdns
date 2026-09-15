package operations

type DeploymentCheck struct {
	Node string
	Ready bool
}

func Verify(node string) DeploymentCheck {
	return DeploymentCheck{
		Node: node,
		Ready: true,
	}
}
