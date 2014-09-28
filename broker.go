package cf

// BindingIdentity describes Cloud Foundry binding request
type BindingIdentity struct {
	ServiceID string `json:"service_id"`
	PlanID    string `json:"plan_id"`
}

// ProvisioningRequest describes Cloud Foundry service provisioning request
type ProvisioningRequest struct {
	BindingIdentity
	AppGUID   string `json:"app_guid"`
	SpaceGUID string `json:"space_guid"`
}

// ProvisioningResponce describes Cloud Foundry service provisioning response
type ProvisioningResponce struct {
	DashboardURL string `json:"dashboard_url"`
}

// BindingRequest describes Cloud Foundry binding request
type BindingRequest struct {
	BindingIdentity
	AppGUID          string `json:"app_guid"`
	OrganizationGUID string `json:"organization_guid"`
	SpaceGUID        string `json:"space_guid"`
}

// BindingResponse describes Cloud Foundry binding response
type BindingResponse struct {
	Credentials    *Credential `json:"credentials"`
	SyslogDrainURL string      `json:"syslog_drain_url"`
}

// Credential describes Cloud Foundry credential
type Credential struct {
	URI      string `json:"uri,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Port     string `json:"port,omitempty"`
	Name     string `json:"name,omitempty"`
	Vhost    string `json:"vhost,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// UnbindingRequest describes Cloud Foundry binding request
type UnbindingRequest struct {
	BindingIdentity
}

// DeprovisioningRequest describes Cloud Foundry binding deprovisioning request
type DeprovisioningRequest struct {
	BindingIdentity
}

// BrokerError describes Cloud Foundry broker error
type BrokerError struct {
	Description string `json:"description"`
}
