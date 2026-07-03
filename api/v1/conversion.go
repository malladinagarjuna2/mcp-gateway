package v1

// Hub marks MCPServerRegistration as a conversion hub.
func (*MCPServerRegistration) Hub() {}

// Hub marks MCPGatewayExtension as a conversion hub.
func (*MCPGatewayExtension) Hub() {}

// Hub marks MCPVirtualServer as a conversion hub.
func (*MCPVirtualServer) Hub() {}
