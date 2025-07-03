package controlplane

// NewDefaultClient creates a new Client with the default control plane URL and
// provided security source.
func NewDefaultClient(opts ...ClientOption) (*Client, error) {
	return NewClient(DefaultControlPlane, opts...)
}
