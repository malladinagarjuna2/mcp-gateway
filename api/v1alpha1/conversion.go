package v1alpha1

import (
	"encoding/json"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	mcpv1 "github.com/Kuadrant/mcp-gateway/api/v1"
)

// ExperimentalDiscoveryAnnotation stores experimental discovery configuration for v1alpha1.
const ExperimentalDiscoveryAnnotation = "mcp.kuadrant.io/experimental-discovery"

// ExperimentalExtensionAnnotation stores experimental extension configuration for v1alpha1.
const ExperimentalExtensionAnnotation = "mcp.kuadrant.io/experimental-extension"

type experimentalDiscoveryData struct {
	// tokenURLElicitation stores the experimental token URL elicitation config.
	// +optional
	TokenURLElicitation *TokenURLElicitationConfig `json:"tokenURLElicitation,omitempty"`
	// userSpecificList stores the experimental user specific list policy.
	// +optional
	UserSpecificList UserSpecificListPolicy `json:"userSpecificList,omitempty"`
}

type experimentalExtensionData struct {
	// urlElicitation stores the experimental URL elicitation policy.
	// +optional
	URLElicitation URLElicitationPolicy `json:"urlElicitation,omitempty"`
}

// ConvertTo converts this MCPServerRegistration to the Hub version (v1).
func (r *MCPServerRegistration) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*mcpv1.MCPServerRegistration)

	b, err := json.Marshal(r)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return err
	}

	expData := experimentalDiscoveryData{
		TokenURLElicitation: r.Spec.TokenURLElicitation,
		UserSpecificList:    r.Spec.UserSpecificList,
	}
	if expData.TokenURLElicitation != nil || expData.UserSpecificList != "" {
		expB, err := json.Marshal(expData)
		if err == nil {
			if dst.Annotations == nil {
				dst.Annotations = make(map[string]string)
			}
			dst.Annotations[ExperimentalDiscoveryAnnotation] = string(expB)
		}
	}
	return nil
}

// ConvertFrom converts from the Hub version (v1) to this MCPServerRegistration.
func (r *MCPServerRegistration) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*mcpv1.MCPServerRegistration)

	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, r); err != nil {
		return err
	}

	if data, ok := src.Annotations[ExperimentalDiscoveryAnnotation]; ok {
		var expData experimentalDiscoveryData
		if err := json.Unmarshal([]byte(data), &expData); err == nil {
			r.Spec.TokenURLElicitation = expData.TokenURLElicitation
			r.Spec.UserSpecificList = expData.UserSpecificList
		}

		delete(r.Annotations, ExperimentalDiscoveryAnnotation)
		if len(r.Annotations) == 0 {
			r.Annotations = nil
		}
	}
	return nil
}

// ConvertTo converts this MCPGatewayExtension to the Hub version (v1).
func (e *MCPGatewayExtension) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*mcpv1.MCPGatewayExtension)

	b, err := json.Marshal(e)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return err
	}

	expData := experimentalExtensionData{
		URLElicitation: e.Spec.URLElicitation,
	}
	if expData.URLElicitation != "" {
		expB, err := json.Marshal(expData)
		if err == nil {
			if dst.Annotations == nil {
				dst.Annotations = make(map[string]string)
			}
			dst.Annotations[ExperimentalExtensionAnnotation] = string(expB)
		}
	}
	return nil
}

// ConvertFrom converts from the Hub version (v1) to this MCPGatewayExtension.
func (e *MCPGatewayExtension) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*mcpv1.MCPGatewayExtension)

	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, e); err != nil {
		return err
	}

	if data, ok := src.Annotations[ExperimentalExtensionAnnotation]; ok {
		var expData experimentalExtensionData
		if err := json.Unmarshal([]byte(data), &expData); err == nil {
			e.Spec.URLElicitation = expData.URLElicitation
		}

		delete(e.Annotations, ExperimentalExtensionAnnotation)
		if len(e.Annotations) == 0 {
			e.Annotations = nil
		}
	}
	return nil
}

// ConvertTo converts this MCPVirtualServer to the Hub version (v1).
func (v *MCPVirtualServer) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*mcpv1.MCPVirtualServer)
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

// ConvertFrom converts from the Hub version (v1) to this MCPVirtualServer.
func (v *MCPVirtualServer) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*mcpv1.MCPVirtualServer)
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}
