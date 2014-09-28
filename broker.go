package cf

// ServiceIdentity describes Cloud Foundry based service identity
type ServiceIdentity struct {
	ServiceID string `json:"service_id"`
	PlanID    string `json:"plan_id"`
}

// ServiceCreationRequest describes Cloud Foundry service provisioning request
type ServiceCreationRequest struct {
	ServiceIdentity
	AppGUID   string `json:"app_guid"`
	SpaceGUID string `json:"space_guid"`
}

// ServiceCreationResponce describes Cloud Foundry service provisioning response
type ServiceCreationResponce struct {
	DashboardURL string `json:"dashboard_url"`
}

// ServiceBindingRequest describes Cloud Foundry service binding request
type ServiceBindingRequest struct {
	ServiceIdentity
	AppGUID          string `json:"app_guid"`
	OrganizationGUID string `json:"organization_guid"`
	SpaceGUID        string `json:"space_guid"`
}

// ServiceBindingResponse describes Cloud Foundry service binding response
type ServiceBindingResponse struct {
	Credentials    *Credential `json:"credentials"`
	SyslogDrainURL string      `json:"syslog_drain_url"`
}

// ServiceBindingDeletionRequest describes Cloud Foundry unbinding request
type ServiceBindingDeletionRequest struct {
	ServiceIdentity
}

// ServiceDeletionRequest describes Cloud Foundry de-previsioning request
type ServiceDeletionRequest struct {
	ServiceIdentity
}

// Credential describes Cloud Foundry service binding credential
type Credential struct {
	URI      string `json:"uri,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Port     string `json:"port,omitempty"`
	Name     string `json:"name,omitempty"`
	Vhost    string `json:"vhost,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// BrokerError describes Cloud Foundry broker error
type BrokerError struct {
	Description string `json:"description"`
}
