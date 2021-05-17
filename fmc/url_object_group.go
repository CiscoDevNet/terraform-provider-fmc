package fmc

type URLObjectGroupUpdateInput struct {
	Type        string `json:"type"`
	Objects      interface{} `json:"objects"`
	Literals     interface{} `json:"literals"`
	Description string `json:"description"`
	Name       string `json:"name"`
	ID          string `json:"id"`
}

type URLObjectGroup struct {
	Type        string `json:"type"`
	Objects         interface{} `json:"objects"`
	Literals interface{}   `json:"literals"`
	Description string `json:"description"`
	Name       string `json:"name"`
}

type URLObjectGroupResponse struct {
	// Links struct {
	// 	Self string `json:"self"`
	// } `json:"links"`
	// Items []struct {
	Links struct {
		Self   string `json:"self"`
	} `json:"links"`
	Type        string `json:"type"`
	URL       string `json:"url"`
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
