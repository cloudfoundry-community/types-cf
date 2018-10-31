package cf

// Catalog describes Cloud Foundry catalog
type Catalog struct {
	Services []*Service `json:"services"`
}

// Service describes Cloud Foundry service
type Service struct {
	ID                   string       `json:"id"`
	Name                 string       `json:"name"`
	Description          string       `json:"description"`
	Bindable             bool         `json:"bindable"`
	InstancesRetrievable bool         `json:"instances_retrievable"`
	BindingsRetrievable  bool         `json:"bindings_retrievable"`
	Tags                 []string     `json:"tags,omitempty"`
	PlanUpdatable        bool         `json:"plan_updateable"`
	Plans                []*Plan      `json:"plans"`
	Requires             []string     `json:"requires,omitempty"`
	Metadata             *ServiceMeta `json:"metadata,omitempty"`
	Dashboard            *Dashboard   `json:"dashboard_client,omitempty"`
}

// Plan describes Cloud Foundry plan structure
type Plan struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Free        bool      `json:"free,omitempty"`
	Bindable    bool      `json:"bindable,omitempty"`
	Metadata    *PlanMeta `json:"metadata,omitempty"`
}

// PlanMeta describers Cloud Foundry plan meta-data
type PlanMeta struct {
	Bullets     []string `json:"bullets"`
	Costs       string   `json:"costs"`
	DisplayName string   `json:"displayName"`
}

// Cost describers Cloud Foundry plan Cost
type Cost struct {
	Amount *Amount `json:"amount"`
	Unit   string  `json:"unit"`
}

// Amount describers Cloud Foundry cost amount
type Amount struct {
	Usd float32 `json:"usd"`
}

// Dashboard describes Cloud Foundry dashboard
type Dashboard struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
	URI    string `json:"redirect_uri"`
}

// ServiceMeta describers Cloud Foundry service meta-data
type ServiceMeta struct {
	DisplayName         string `json:"displayName"`
	ImageURL            string `json:"imageUrl"`
	Description         string `json:"longDescription"`
	ProviderDisplayName string `json:"providerDisplayName"`
	DocURL              string `json:"documentationUrl"`
	SupportURL          string `json:"supportUrl"`
}
