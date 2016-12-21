package govrealize

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"time"
)

const (
	machineBasePath          = "catalog-service/api/consumer"
	entitledCatalogItemsPath = machineBasePath + "/entitledCatalogItems"
	requestsPath             = machineBasePath + "/requests"
	resourcesPath            = machineBasePath + "/resources"
	machinePath              = resourcesPath + "/types/Infrastructure.Machine"
)

// MachineService is an interface for interfacing with the storage endpoints of the vRealize API.
type MachineService interface {
	ListMachines(*ListOptions) ([]Machine, *Response, error)
	GetMachine(string) (*Machine, *Response, error)
	CreateMachine(*MachineCreateRequest) (*Machine, *Response, error)
	DestroyMachine(string) (bool, *Response, error)
}

// MachineServiceOp handles communication with the machines related methods of the vRealize API.
type MachineServiceOp struct {
	client *Client
}

var _ MachineService = &MachineServiceOp{}

// FIXME: Implement remainder of machine struct
// Various properties have not been commented as they are not required for this library's original use case.

// Machine represents a vRealize machine. Commented properties left to be implmented in future.
type Machine struct {
	ID string `json:"id"`
	// IconID          string `json:"iconId"`
	ResourceTypeRef MachineResourceTypeRef `json:"resourceTypeRef"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	Status          string                 `json:"status"`
	CatalogItem     MachineCatalogItem     `json:"catalogItem"`
	RequestID       string                 `json:"requestId"`
	// ProviderBinding struct {
	// 	BindingID   string `json:"bindingId"`
	// 	ProviderRef struct {
	// 		ID    string `json:"id"`
	// 		Label string `json:"label"`
	// 	} `json:"providerRef"`
	// } `json:"providerBinding"`
	// Owners []struct {
	// 	TenantName string `json:"tenantName"`
	// 	Ref        string `json:"ref"`
	// 	Type       string `json:"type"`
	// 	Value      string `json:"value"`
	// } `json:"owners"`
	Organization MachineOrganization `json:"organization"`
	DateCreated  time.Time           `json:"dateCreated"`
	LastUpdated  time.Time           `json:"lastUpdated"`
	HasLease     bool                `json:"hasLease"`
	// Lease       struct {
	// 	Start time.Time `json:"start"`
	// } `json:"lease"`
	LeaseForDisplay interface{} `json:"leaseForDisplay"`
	// HasCosts        bool        `json:"hasCosts"`
	// Costs           struct {
	// 	LeaseRate struct {
	// 		Type string `json:"type"`
	// 		Cost struct {
	// 			Type         string  `json:"type"`
	// 			CurrencyCode string  `json:"currencyCode"`
	// 			Amount       float64 `json:"amount"`
	// 		} `json:"cost"`
	// 		Basis struct {
	// 			Type   string `json:"type"`
	// 			Unit   string `json:"unit"`
	// 			Amount int    `json:"amount"`
	// 		} `json:"basis"`
	// 	} `json:"leaseRate"`
	// } `json:"costs"`
	// CostToDate struct {
	// 	Type         string  `json:"type"`
	// 	CurrencyCode string  `json:"currencyCode"`
	// 	Amount       float64 `json:"amount"`
	// } `json:"costToDate"`
	// TotalCost      interface{}   `json:"totalCost"`
	// ChildResources []interface{} `json:"childResources"`
	Operations []MachineOperation `json:"operations"`
	// Forms struct {
	// 	CatalogResourceInfoHidden bool `json:"catalogResourceInfoHidden"`
	// 	Details                   struct {
	// 		Type             string      `json:"type"`
	// 		ExtensionID      string      `json:"extensionId"`
	// 		ExtensionPointID interface{} `json:"extensionPointId"`
	// 	} `json:"details"`
	// } `json:"forms"`
	ResourceData MachineResourceData `json:"resourceData"`
}

// MachineResourceTypeRef represents a resourceTypeRef
type MachineResourceTypeRef struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// MachineCatalogItem represents a catalogueItem
type MachineCatalogItem struct {
	ID string `json:"id"`
	// 	Label string `json:"label"`
}

// MachineOrganization represents an organization
type MachineOrganization struct {
	TenantRef    string `json:"tenantRef"`
	SubtenantRef string `json:"subtenantRef"`
}

// MachineOperation represents an operation
type MachineOperation struct {
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	IconID         string      `json:"iconId"`
	Type           string      `json:"type"`
	ID             string      `json:"id"`
	ExtensionID    string      `json:"extensionId"`
	ProviderTypeID interface{} `json:"providerTypeId"`
	BindingID      interface{} `json:"bindingId"`
	HasForm        bool        `json:"hasForm"`
	FormScale      interface{} `json:"formScale"`
}

// MachineResourceData represents resourceData
type MachineResourceData struct {
	Entries []MachineDataEntry `json:"entries"`
}

// MachineDataEntry represents a dataEntry
type MachineDataEntry struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// DestroyActionID returns the ID of the Destroy action
func (m Machine) DestroyActionID() string {
	for _, operation := range m.Operations {
		if operation.Name == "Destroy" {
			return operation.ID
		}
	}

	return ""
}

// ReconfigureActionID returns the ID of the Reconfigure action
func (m Machine) ReconfigureActionID() string {
	for _, operation := range m.Operations {
		if operation.Name == "Reconfigure" {
			return operation.ID
		}
	}

	return ""
}

// IPAddress returns the IP Address of the machine
func (m Machine) IPAddress() string {
	ipAddr := ""
	for _, dataEntry := range m.ResourceData.Entries {
		if dataEntry.Key == "NETWORK_LIST" {
			items := dataEntry.Value.(map[string]interface{})["items"].([]interface{})
			for _, item := range items {
				if item.(map[string]interface{})["componentTypeId"] == "com.vmware.csp.component.iaas.proxy.provider" {
					values := item.(map[string]interface{})["values"]
					for _, entry := range values.(map[string]interface{})["entries"].([]interface{}) {
						if entry.(map[string]interface{})["key"] == "NETWORK_ADDRESS" {
							ipAddr = entry.(map[string]interface{})["value"].(map[string]interface{})["value"].(string)
						}
					}
				}
			}
		}
	}
	return ipAddr
}

// MachineName returns the name of the machine
func (m Machine) MachineName() string {

	for _, dataEntry := range m.ResourceData.Entries {
		if dataEntry.Key == "MachineName" {
			value := dataEntry.Value.(map[string]interface{})["value"].(string)
			return value
		}
	}

	return ""
}

// Link is a representation of a link
type Link struct {
	Type string `json:"@type"`
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

// MachineRequestDataEntry represents a requestDataEntry
type MachineRequestDataEntry struct {
	Key   string                       `json:"key"`
	Value MachineRequestDataEntryValue `json:"value"`
}

// MachineRequestDataEntryValue represents a dataEntryValue
type MachineRequestDataEntryValue struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// MachineCreateRequest represents a request to create a machine
type MachineCreateRequest struct {
	Type           string                             `json:"@type"`
	CatalogItemRef MachineCreateRequestCatalogItemRef `json:"catalogItemRef"`
	Organization   MachineCreateRequestOrganization   `json:"organization"`
	State          string                             `json:"state"`
	RequestNumber  int                                `json:"requestNumber"`
	RequestData    MachineCreateRequestRequestData    `json:"requestData"`
}

// MachineCreateRequestCatalogItemRef represents a catalogItemRef
type MachineCreateRequestCatalogItemRef struct {
	ID string `json:"id"`
}

// MachineCreateRequestOrganization represents an organization
type MachineCreateRequestOrganization struct {
	TenantRef    string `json:"tenantRef"`
	SubtenantRef string `json:"subtenantRef"`
}

// MachineCreateRequestRequestData represents a requestData
type MachineCreateRequestRequestData struct {
	Entries []MachineRequestDataEntry `json:"entries"`
}

// MachineCreateResponse represents a response to a request to create a machine
type MachineCreateResponse struct {
	Location url.URL
}

func (mcr MachineCreateResponse) String() string {
	return fmt.Sprintf("%s", mcr.Location.String())
}

type machinesRoot struct {
	Machines []Machine `json:"content"`
	Links    []Link    `json:"links"`
}

// ListMachines lists all machines.
func (svc *MachineServiceOp) ListMachines(opt *ListOptions) ([]Machine, *Response, error) {
	path, err := addOptions(machinePath, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := svc.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	machines := new(machinesRoot)
	resp, err := svc.client.Do(req, machines)
	if err != nil {
		return nil, resp, err
	}

	if l := machines.Links; l != nil {
		resp.Links = l
	}

	return machines.Machines, resp, nil
}

// CreateMachine creates a machine.
func (svc *MachineServiceOp) CreateMachine(createRequest *MachineCreateRequest) (*Machine, *Response, error) {

	path := requestsPath
	req, err := svc.client.NewRequest("POST", path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(req, nil)
	if err != nil {
		return nil, nil, err
	}

	loc, err := resp.Location()
	if err != nil {
		return nil, nil, err
	}

	vrr := new(vrealizeRequestResponse)

	successRequestCompletionState := "SUCCESSFUL"
	failedRequestCompletionState := "FAILED"

	// FIXME: If the createRequest fails or has an invalid response this should
	// be handled with the relevant information

	// Check request until requestCompletion.requestCompletionState == 'SUCCESSFUL'
	i := 1
	for i < 360 {

		fmt.Println("[INFO] Checking machine request status")

		vrr = new(vrealizeRequestResponse)
		req, err = svc.client.NewRequest("GET", loc.String(), nil)
		if err != nil {
			return nil, nil, err
		}
		resp, err = svc.client.Do(req, vrr)

		fmt.Printf("[INFO] Checking machine request: %s.\n", vrr.ID)
		if reflect.DeepEqual(vrr.RequestCompletion.RequestCompletionState, successRequestCompletionState) {
			fmt.Println("[INFO] Machine request status SUCCESSFUL. Machine created.")
			break
		}

		if reflect.DeepEqual(vrr.RequestCompletion.RequestCompletionState, failedRequestCompletionState) {
			fmt.Println("[ERROR] Machine request status FAILED. Machine not created.")
			failedErr := errors.New(vrr.RequestCompletion.CompletionDetails)
			return nil, nil, failedErr
		}

		i = i + 1
		time.Sleep(10 * time.Second)
	}

	if !reflect.DeepEqual(vrr.RequestCompletion.RequestCompletionState, successRequestCompletionState) {
		fmt.Println("[ERROR] Machine request status TIMEOUT. Machine not created.")
		timeoutErr := errors.New("Machine request timedout. Machine not created.")
		return nil, nil, timeoutErr
	}

	// Get resource (using request ID) and return
	path = fmt.Sprintf("%s/%s/resources", requestsPath, vrr.ID)

	vrrr := new(vrealizeRequestResourcesResponse)
	req, err = svc.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err = svc.client.Do(req, vrrr)

	if err != nil {
		return nil, resp, err
	}

	fmt.Printf("%+v", vrrr.Content)

	machine := vrrr.Content[0]

	fmt.Printf("[INFO] Machine %s online.\n", machine.ID)

	return machine, resp, nil
}

// GetMachine retrieves an individual machine.
func (svc *MachineServiceOp) GetMachine(id string) (*Machine, *Response, error) {

	path := fmt.Sprintf("%s/%s", resourcesPath, id)

	req, err := svc.client.NewRequest("GET", path, nil)

	if err != nil {
		return nil, nil, err
	}

	machine := new(Machine)
	resp, err := svc.client.Do(req, machine)
	if err != nil {
		return nil, resp, err
	}

	return machine, resp, nil
}

// DestroyMachine Destroys a machine.
func (svc *MachineServiceOp) DestroyMachine(id string) (bool, *Response, error) {

	machine, _, err := svc.GetMachine(id)

	if err != nil {
		return false, nil, err
	}

	destroyRequest := vrealizeResourceActionRequest{
		Type: "ResourceActionRequest",
		ResourceRef: vrealizeResourceActionRequestResourceRef{
			ID: machine.ID,
		},
		ResourceActionRef: vrealizeResourceActionRequestResourceActionRef{
			ID: machine.DestroyActionID(),
		},
		Organization: vrealizeResourceActionRequestOrganization{
			TenantRef:    machine.Organization.TenantRef,
			SubtenantRef: machine.Organization.SubtenantRef,
		},
		State:         "SUBMITTED",
		RequestNumber: 0,
	}

	path := requestsPath
	req, err := svc.client.NewRequest("POST", path, destroyRequest)
	if err != nil {
		return false, nil, err
	}
	resp, err := svc.client.Do(req, nil)
	loc, _ := resp.Location()

	vrr := new(vrealizeRequestResponse)
	req, err = svc.client.NewRequest("GET", loc.String(), nil)
	if err != nil {
		return false, nil, err
	}
	resp, err = svc.client.Do(req, vrr)

	successRequestCompletionState := "SUCCESSFUL"
	failedRequestCompletionState := "FAILED"

	// Check request until requestCompletion.requestCompletionState == 'SUCCESSFUL'
	i := 1
	for i < 180 {

		fmt.Println("[INFO] Checking machine destroy request status")

		vrr = new(vrealizeRequestResponse)
		req, err = svc.client.NewRequest("GET", loc.String(), nil)
		if err != nil {
			return false, nil, err
		}
		resp, err = svc.client.Do(req, vrr)

		fmt.Printf("[INFO] Checking machine request: %s.\n", vrr.ID)

		if reflect.DeepEqual(vrr.RequestCompletion.RequestCompletionState, successRequestCompletionState) {
			fmt.Println("[INFO] Machine request status SUCCESSFUL. Machine destroyed.")
			break
		}

		if reflect.DeepEqual(vrr.RequestCompletion.RequestCompletionState, failedRequestCompletionState) {
			fmt.Println("[ERROR] Machine request status FAILED. Machine not created.")
			failedErr := errors.New(vrr.RequestCompletion.CompletionDetails)
			return false, nil, failedErr
		}

		i = i + 1
		time.Sleep(10 * time.Second)
	}

	if reflect.DeepEqual(vrr.RequestCompletion.RequestCompletionState, successRequestCompletionState) {
		fmt.Printf("[INFO] Machine %s destroyed.\n", machine.ID)

		return true, resp, nil
	}

	fmt.Printf("[ERROR] Machine %s not destroyed.\n", machine.ID)

	return false, resp, nil
}
