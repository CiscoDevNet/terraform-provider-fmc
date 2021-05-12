package fmc

type NetworkObjectUpdateInput struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Type        string `json:"type"`
	ID          string `json:"id"`
}

type NetworkObject struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type NetworkObjectResponse struct {
	// Links struct {
	// 	Self string `json:"self"`
	// } `json:"links"`
	// Items []struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Name        string `json:"name"`
	ID          string `json:"id"`
	Metadata    struct {
		Timestamp int `json:"timestamp"`
		LastUser  struct {
			Name string `json:"name"`
		} `json:"lastUser"`
		Domain struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"domain"`
		IPType     string `json:"ipType"`
		ParentType string `json:"parentType"`
	} `json:"metadata"`
	// } `json:"items"`
}

// {
//     "name": "net1",
//     "value": "1.0.0.0/24",
//     "overridable": false,
//     "description": "Network obj 1",
//     "type": "Network"
// }
// POST /fmc_config/v1/domain/DomainUUID/object/networks
