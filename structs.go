package govrealize

import "time"

// vrealizeRequestResponse represents a response to a request to create a machine
type vrealizeRequestResponse struct {
	Type          string `json:"@type"`
	ID            string `json:"id"`
	IconID        string `json:"iconId"`
	Version       int    `json:"version"`
	RequestNumber int    `json:"requestNumber"`
	State         string `json:"state"`
	Description   string `json:"description"`
	Reasons       string `json:"reasons"`
	RequestedFor  string `json:"requestedFor"`
	RequestedBy   string `json:"requestedBy"`
	Organization  struct {
		TenantRef      string `json:"tenantRef"`
		TenantLabel    string `json:"tenantLabel"`
		SubtenantRef   string `json:"subtenantRef"`
		SubtenantLabel string `json:"subtenantLabel"`
	} `json:"organization"`
	RequestCompletion struct {
		RequestCompletionState string `json:"requestCompletionState"`
		CompletionDetails      string `json:"completionDetails"`
	} `json:"requestCompletion"`
	RequestorEntitlementID string      `json:"requestorEntitlementId"`
	PreApprovalID          interface{} `json:"preApprovalId"`
	PostApprovalID         interface{} `json:"postApprovalId"`
	DateCreated            time.Time   `json:"dateCreated"`
	LastUpdated            time.Time   `json:"lastUpdated"`
	DateSubmitted          time.Time   `json:"dateSubmitted"`
	Quote                  interface{} `json:"quote"`
	RequestData            struct {
		Entries []struct {
			Key   string `json:"key"`
			Value struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"value"`
		} `json:"entries"`
	} `json:"requestData"`
	RetriesRemaining         int    `json:"retriesRemaining"`
	RequestedItemName        string `json:"requestedItemName"`
	RequestedItemDescription string `json:"requestedItemDescription"`
	StateName                string `json:"stateName"`
	Phase                    string `json:"phase"`
	ApprovalStatus           string `json:"approvalStatus"`
	WaitingStatus            string `json:"waitingStatus"`
	ExecutionStatus          string `json:"executionStatus"`
	CatalogItemRef           struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	} `json:"catalogItemRef"`
}

// vrealizeRequestResourcesResponse represents a response to a request for a requests resources
type vrealizeRequestResourcesResponse struct {
	Links    []interface{} `json:"links"`
	Content  []*Machine    `json:"content"`
	Metadata struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
		Offset        int `json:"offset"`
	} `json:"metadata"`
}

type vrealizeResourceActionRequest struct {
	Type              string                                         `json:"@type"`
	ResourceRef       vrealizeResourceActionRequestResourceRef       `json:"resourceRef"`
	ResourceActionRef vrealizeResourceActionRequestResourceActionRef `json:"resourceActionRef"`
	Organization      vrealizeResourceActionRequestOrganization      `json:"organization"`
	State             string                                         `json:"state"`
	RequestNumber     int                                            `json:"requestNumber"`
	RequestData       vrealizeResourceActionRequestRequestData       `json:"requestData"`
}

type vrealizeResourceActionRequestResourceRef struct {
	ID string `json:"id"`
}

type vrealizeResourceActionRequestResourceActionRef struct {
	ID string `json:"id"`
}

type vrealizeResourceActionRequestOrganization struct {
	TenantRef    string `json:"tenantRef"`
	SubtenantRef string `json:"subtenantRef"`
}

type vrealizeResourceActionRequestRequestData struct {
	Entries []interface{} `json:"entries"`
}
