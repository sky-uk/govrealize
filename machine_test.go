package govrealize

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestMachine_ListMachines(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/catalog-service/api/consumer/resources/types/Infrastructure.Machine", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, jBlobListMachines)
	})

	machines, _, err := client.Machine.ListMachines(nil)

	if err != nil {
		t.Errorf("Machine.ListMachines returned error: %v", err)
	}

	dateCreated, _ := time.Parse(time.RFC3339, "2016-07-15T11:14:06.542Z")
	lastUpdated, _ := time.Parse(time.RFC3339, "2016-07-15T11:14:09.468Z")

	expected := []Machine{
		{
			ID: "536ee25b-94b1-48b1-b5e3-b15c52690c4c",
			ResourceTypeRef: MachineResourceTypeRef{
				ID:    "Infrastructure.Virtual",
				Label: "Virtual Machine",
			},
			Name:   "vm001448",
			Status: "ACTIVE",
			CatalogItem: MachineCatalogItem{
				ID: "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
			},
			RequestID: "161c66a0-e300-4703-aaa4-aac1029c9f34",
			Organization: MachineOrganization{
				TenantRef:    "vsphere.local",
				SubtenantRef: "f04f060d-73be-48a3-b82c-20cb98efd2d2",
			},
			DateCreated: dateCreated,
			LastUpdated: lastUpdated,
			HasLease:    true,
			LeaseForDisplay: map[string]interface{}{
				"type":   "timeSpan",
				"unit":   "DAYS",
				"amount": 5.0,
			},
			Operations: []MachineOperation{{
				Name:           "Connect by Using SSH",
				Description:    "Infrastructure connect using SSH",
				IconID:         "cafe_default_icon_genericResourceOperation",
				Type:           "EXTENSION",
				ID:             "0c369513-715c-42af-a9c7-260031244dae",
				ExtensionID:    "csp.places.iaas.item.window.ConnectViaSsh",
				ProviderTypeID: nil,
				BindingID:      nil,
				HasForm:        false,
				FormScale:      nil,
			}, {
				Name:           "Connect to Remote Console",
				Description:    "Infrastructure connect using remote console",
				IconID:         "cafe_default_icon_genericResourceOperation",
				Type:           "EXTENSION",
				ID:             "fcbf9796-cabe-4d92-ad5c-ebb7843787ab",
				ExtensionID:    "csp.places.iaas.item.window.ConnectViaVmrc",
				ProviderTypeID: nil,
				BindingID:      nil,
				HasForm:        false,
				FormScale:      nil,
			}, {
				Name:           "Destroy",
				Description:    "Destroy a virtual machine.",
				IconID:         "virtualDestroy.png",
				Type:           "ACTION",
				ID:             "5883bbea-cd9a-4bf3-a2b3-f6cc20135435",
				ExtensionID:    "",
				ProviderTypeID: "com.vmware.csp.iaas.blueprint.service",
				BindingID:      "Infrastructure.Virtual.Action.Destroy",
				HasForm:        false,
				FormScale:      nil,
			}, {
				Name:           "Reconfigure",
				Description:    "Reconfigure a machine.",
				IconID:         "machineReconfigure.png",
				Type:           "ACTION",
				ID:             "4ae8147e-ce45-4b25-a8a7-3c0f02281363",
				ExtensionID:    "",
				ProviderTypeID: "com.vmware.csp.iaas.blueprint.service",
				BindingID:      "Infrastructure.Machine.Action.Reconfigure",
				HasForm:        true,
				FormScale:      "BIG",
			}},
			ResourceData: MachineResourceData{
				Entries: []MachineDataEntry{
					{
						Key: "NETWORK_LIST",
						Value: map[string]interface{}{
							"items": []interface{}{
								map[string]interface{}{
									"typeFilter": interface{}(nil),
									"values": map[string]interface{}{
										"entries": []interface{}{
											map[string]interface{}{
												"key": "NETWORK_ADDRESS", "value": map[string]interface{}{
													"type":  "string",
													"value": "10.10.10.20",
												},
											}, map[string]interface{}{
												"key": "NETWORK_MAC_ADDRESS",
												"value": map[string]interface{}{
													"type":  "string",
													"value": "00:50:50:aa:01:1c",
												},
											}, map[string]interface{}{
												"key": "NETWORK_NAME",
												"value": map[string]interface{}{
													"type":  "string",
													"value": "local.network",
												},
											},
										},
									},
									"type":            "complex",
									"componentTypeId": "com.vmware.csp.component.iaas.proxy.provider",
									"componentId":     interface{}(nil),
									"classId":         "dynamicops.api.model.NetworkViewModel",
								},
							},
							"type":          "multiple",
							"elementTypeId": "COMPLEX",
						},
					},
					{
						Key:   "MachineGroupName",
						Value: map[string]interface{}{"type": "string", "value": "Elements"},
					},
					{
						Key:   "MachineName",
						Value: map[string]interface{}{"type": "string", "value": "vm001234"},
					},
					{
						Key:   "Expire",
						Value: map[string]interface{}{"type": "boolean", "value": true},
					},
				},
			},
		},
		{
			ID: "96d414c6-295e-4e3a-ac59-eb9456c1e1d1",
			ResourceTypeRef: MachineResourceTypeRef{
				ID:    "Infrastructure.Virtual",
				Label: "Virtual Machine",
			},
			Name:   "vm001448",
			Status: "ACTIVE",
			CatalogItem: MachineCatalogItem{
				ID: "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
			},
			RequestID: "161c66a0-e300-4703-aaa4-aac1029c9f34",
			Organization: MachineOrganization{
				TenantRef:    "vsphere.local",
				SubtenantRef: "f04f060d-73be-48a3-b82c-20cb98efd2d2",
			},
			DateCreated: dateCreated,
			LastUpdated: lastUpdated,
			HasLease:    true,
			LeaseForDisplay: map[string]interface{}{
				"type":   "timeSpan",
				"unit":   "DAYS",
				"amount": 5.0,
			},
			Operations: []MachineOperation{{
				Name:           "Connect by Using SSH",
				Description:    "Infrastructure connect using SSH",
				IconID:         "cafe_default_icon_genericResourceOperation",
				Type:           "EXTENSION",
				ID:             "0c369513-715c-42af-a9c7-260031244dae",
				ExtensionID:    "csp.places.iaas.item.window.ConnectViaSsh",
				ProviderTypeID: nil,
				BindingID:      nil,
				HasForm:        false,
				FormScale:      nil,
			}, {
				Name:           "Connect to Remote Console",
				Description:    "Infrastructure connect using remote console",
				IconID:         "cafe_default_icon_genericResourceOperation",
				Type:           "EXTENSION",
				ID:             "fcbf9796-cabe-4d92-ad5c-ebb7843787ab",
				ExtensionID:    "csp.places.iaas.item.window.ConnectViaVmrc",
				ProviderTypeID: nil,
				BindingID:      nil,
				HasForm:        false,
				FormScale:      nil,
			}, {
				Name:           "Destroy",
				Description:    "Destroy a virtual machine.",
				IconID:         "virtualDestroy.png",
				Type:           "ACTION",
				ID:             "5883bbea-cd9a-4bf3-a2b3-f6cc20135436",
				ExtensionID:    "",
				ProviderTypeID: "com.vmware.csp.iaas.blueprint.service",
				BindingID:      "Infrastructure.Virtual.Action.Destroy",
				HasForm:        false,
				FormScale:      nil,
			}, {
				Name:           "Reconfigure",
				Description:    "Reconfigure a machine.",
				IconID:         "machineReconfigure.png",
				Type:           "ACTION",
				ID:             "4ae8147e-ce45-4b25-a8a7-3c0f02281364",
				ExtensionID:    "",
				ProviderTypeID: "com.vmware.csp.iaas.blueprint.service",
				BindingID:      "Infrastructure.Machine.Action.Reconfigure",
				HasForm:        true,
				FormScale:      "BIG",
			}},
			ResourceData: MachineResourceData{
				Entries: []MachineDataEntry{
					{
						Key: "NETWORK_LIST",
						Value: map[string]interface{}{
							"items": []interface{}{
								map[string]interface{}{
									"typeFilter": interface{}(nil),
									"values": map[string]interface{}{
										"entries": []interface{}{
											map[string]interface{}{
												"key": "NETWORK_ADDRESS",
												"value": map[string]interface{}{
													"type":  "string",
													"value": "10.70.10.10",
												},
											}, map[string]interface{}{
												"key": "NETWORK_MAC_ADDRESS",
												"value": map[string]interface{}{
													"type":  "string",
													"value": "00:50:50:aa:01:1c",
												},
											}, map[string]interface{}{
												"key": "NETWORK_NAME",
												"value": map[string]interface{}{
													"type":  "string",
													"value": "local.network",
												},
											},
										},
									},
									"type":            "complex",
									"componentTypeId": "com.vmware.csp.component.iaas.proxy.provider",
									"componentId":     interface{}(nil),
									"classId":         "dynamicops.api.model.NetworkViewModel",
								},
							},
							"type":          "multiple",
							"elementTypeId": "COMPLEX",
						},
					},
					{
						Key:   "MachineGroupName",
						Value: map[string]interface{}{"type": "string", "value": "Elements"},
					},
					{
						Key:   "MachineName",
						Value: map[string]interface{}{"type": "string", "value": "vm005678"},
					},
					{
						Key:   "Expire",
						Value: map[string]interface{}{"type": "boolean", "value": true},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(machines, expected) {
		t.Errorf("Machine.ListMachines returned %#v, expected %#v", machines, expected)
	}

	if machines[0].DestroyActionID() != "5883bbea-cd9a-4bf3-a2b3-f6cc20135435" {
		t.Errorf("Machine.DestroyActionID returned %#v, expected %#v", machines[0].DestroyActionID(), "5883bbea-cd9a-4bf3-a2b3-f6cc20135435")
	}

	if machines[0].ReconfigureActionID() != "4ae8147e-ce45-4b25-a8a7-3c0f02281363" {
		t.Errorf("Machine.ReconfigureActionID returned %#v, expected %#v", machines[0].ReconfigureActionID(), "4ae8147e-ce45-4b25-a8a7-3c0f02281363")
	}

	if machines[0].IPAddress() != "10.10.10.20" {
		t.Errorf("Machine.IPAddress returned %#v, expected %#v", machines[0].IPAddress(), "10.10.10.20")
	}

	if machines[1].DestroyActionID() != "5883bbea-cd9a-4bf3-a2b3-f6cc20135436" {
		t.Errorf("Machine.DestroyActionID returned %#v, expected %#v", machines[1].DestroyActionID(), "5883bbea-cd9a-4bf3-a2b3-f6cc20135436")
	}

	if machines[1].ReconfigureActionID() != "4ae8147e-ce45-4b25-a8a7-3c0f02281364" {
		t.Errorf("Machine.ReconfigureActionID returned %#v, expected %#v", machines[1].ReconfigureActionID(), "4ae8147e-ce45-4b25-a8a7-3c0f02281364")
	}

	if machines[1].IPAddress() != "10.70.10.10" {
		t.Errorf("Machine.IPAddress returned %#v, expected %#v", machines[1].IPAddress(), "10.70.10.10")
	}
}

func TestMachine_Get(t *testing.T) {

	setup()

	defer teardown()

	dateCreated, _ := time.Parse(time.RFC3339, "2016-07-15T12:55:28.369Z")
	lastUpdated, _ := time.Parse(time.RFC3339, "2016-07-15T12:55:32.507Z")

	expected := &Machine{
		ID: "b4ce1b01-365a-4964-963a-90263c16ebda",
		ResourceTypeRef: MachineResourceTypeRef{
			ID:    "Infrastructure.Virtual",
			Label: "Virtual Machine",
		},
		Name:   "vm001451",
		Status: "ACTIVE",
		CatalogItem: MachineCatalogItem{
			ID: "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
		},
		RequestID: "9f3f8c07-16a3-4f68-a34b-73f756ae960a",
		Organization: MachineOrganization{
			TenantRef:    "vsphere.local",
			SubtenantRef: "f04f060d-73be-48a3-b82c-20cb98efd2d2",
		},
		DateCreated: dateCreated,
		LastUpdated: lastUpdated,
		HasLease:    true,
		LeaseForDisplay: map[string]interface{}{
			"amount": 5.0,
			"type":   "timeSpan",
			"unit":   "DAYS",
		},
		Operations: []MachineOperation{{
			Name:           "Connect by Using SSH",
			Description:    "Infrastructure connect using SSH",
			IconID:         "cafe_default_icon_genericResourceOperation",
			Type:           "EXTENSION",
			ID:             "0c369513-715c-42af-a9c7-260031244dae",
			ExtensionID:    "csp.places.iaas.item.window.ConnectViaSsh",
			ProviderTypeID: nil,
			BindingID:      nil,
			HasForm:        false,
			FormScale:      nil,
		}, {
			Name:           "Connect to Remote Console",
			Description:    "Infrastructure connect using remote console",
			IconID:         "cafe_default_icon_genericResourceOperation",
			Type:           "EXTENSION",
			ID:             "fcbf9796-cabe-4d92-ad5c-ebb7843787ab",
			ExtensionID:    "csp.places.iaas.item.window.ConnectViaVmrc",
			ProviderTypeID: nil,
			BindingID:      nil,
			HasForm:        false,
			FormScale:      nil,
		}, {
			Name:           "Destroy",
			Description:    "Destroy a virtual machine.",
			IconID:         "virtualDestroy.png",
			Type:           "ACTION",
			ID:             "5883bbea-cd9a-4bf3-a2b3-f6cc20135435",
			ExtensionID:    "",
			ProviderTypeID: "com.vmware.csp.iaas.blueprint.service",
			BindingID:      "Infrastructure.Virtual.Action.Destroy",
			HasForm:        false,
			FormScale:      nil,
		}, {
			Name:           "Reconfigure",
			Description:    "Reconfigure a machine.",
			IconID:         "machineReconfigure.png",
			Type:           "ACTION",
			ID:             "4ae8147e-ce45-4b25-a8a7-3c0f02281363",
			ExtensionID:    "",
			ProviderTypeID: "com.vmware.csp.iaas.blueprint.service",
			BindingID:      "Infrastructure.Machine.Action.Reconfigure",
			HasForm:        true,
			FormScale:      "BIG",
		}},
		ResourceData: MachineResourceData{
			Entries: []MachineDataEntry{
				{
					Key: "NETWORK_LIST",
					Value: map[string]interface{}{
						"items": []interface{}{
							map[string]interface{}{
								"typeFilter": interface{}(nil), "values": map[string]interface{}{
									"entries": []interface{}{
										map[string]interface{}{
											"key": "NETWORK_ADDRESS",
											"value": map[string]interface{}{
												"type": "string", "value": "10.10.10.20",
											},
										}, map[string]interface{}{
											"key": "NETWORK_MAC_ADDRESS",
											"value": map[string]interface{}{
												"type":  "string",
												"value": "00:50:50:aa:01:1c",
											},
										}, map[string]interface{}{
											"key": "NETWORK_NAME",
											"value": map[string]interface{}{
												"type":  "string",
												"value": "local.network",
											},
										},
									},
								},
								"type":            "complex",
								"componentTypeId": "com.vmware.csp.component.iaas.proxy.provider",
								"componentId":     interface{}(nil),
								"classId":         "dynamicops.api.model.NetworkViewModel",
							},
						},
						"type":          "multiple",
						"elementTypeId": "COMPLEX",
					},
				},
				{
					Key:   "MachineGroupName",
					Value: map[string]interface{}{"type": "string", "value": "Elements"},
				},
				{
					Key:   "MachineName",
					Value: map[string]interface{}{"type": "string", "value": "vm001234"},
				},
				{
					Key:   "Expire",
					Value: map[string]interface{}{"type": "boolean", "value": true},
				},
			},
		},
	}

	mux.HandleFunc("/catalog-service/api/consumer/resources/b4ce1b01-365a-4964-963a-90263c16ebda", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, jBlobResource)
	})

	machine, _, err := client.Machine.GetMachine("b4ce1b01-365a-4964-963a-90263c16ebda")

	if err != nil {
		t.Errorf("Machine.GetMachine returned error: %v", err)
	}

	if !reflect.DeepEqual(machine, expected) {
		t.Errorf("Machine.GetMachine returned %#v, expected %#v", machine, expected)
	}

	if machine.DestroyActionID() != "5883bbea-cd9a-4bf3-a2b3-f6cc20135435" {
		t.Errorf("Machine.DestroyActionID returned %#v, expected %#v", machine.DestroyActionID(), "5883bbea-cd9a-4bf3-a2b3-f6cc20135435")
	}

	if machine.ReconfigureActionID() != "4ae8147e-ce45-4b25-a8a7-3c0f02281363" {
		t.Errorf("Machine.ReconfigureActionID returned %#v, expected %#v", machine.ReconfigureActionID(), "4ae8147e-ce45-4b25-a8a7-3c0f02281363")
	}

	if machine.IPAddress() != "10.10.10.20" {
		t.Errorf("Machine.IPAddress returned %#v, expected %#v", machine.IPAddress(), "10.10.10.20")
	}
}

func TestMachine_Create(t *testing.T) {
	setup()
	defer teardown()

	mrcevarr := []MachineRequestDataEntry{}

	mrcevarr = append(mrcevarr, MachineRequestDataEntry{
		Key: "provider-provisioningGroupId",
		Value: MachineRequestDataEntryValue{
			Type:  "string",
			Value: "f04f060d-73be-48a3-b82c-20cb98efd2d2",
		},
	})

	mrcevarr = append(mrcevarr, MachineRequestDataEntry{
		Key: "provider-VirtualMachine.CPU.Count",
		Value: MachineRequestDataEntryValue{
			Type:  "integer",
			Value: "1",
		},
	})

	mrcevarr = append(mrcevarr, MachineRequestDataEntry{
		Key: "provider-VirtualMachine.Memory.Size",
		Value: MachineRequestDataEntryValue{
			Type:  "integer",
			Value: "1024",
		},
	})

	mrcevarr = append(mrcevarr, MachineRequestDataEntry{
		Key: "provider-Cafe.Shim.VirtualMachine.Description",
		Value: MachineRequestDataEntryValue{
			Type:  "string",
			Value: "Description test",
		},
	})

	mrcevarr = append(mrcevarr, MachineRequestDataEntry{
		Key: "reasons",
		Value: MachineRequestDataEntryValue{
			Type:  "string",
			Value: "Reasons test",
		},
	})

	createRequest := &MachineCreateRequest{
		Type: "CatalogItemRequest",
		CatalogItemRef: MachineCreateRequestCatalogItemRef{
			ID: "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
		},
		Organization: MachineCreateRequestOrganization{
			TenantRef:    "vsphere.local",
			SubtenantRef: "f04f060d-73be-48a3-b82c-20cb98efd2d2",
		},
		State:         "SUBMITTED",
		RequestNumber: 0,
		RequestData: MachineCreateRequestRequestData{
			Entries: mrcevarr,
		},
	}

	dateCreated, _ := time.Parse(time.RFC3339, "2016-07-15T12:55:28.369Z")
	lastUpdated, _ := time.Parse(time.RFC3339, "2016-07-15T12:55:32.507Z")

	want := &Machine{
		ID: "b4ce1b01-365a-4964-963a-90263c16ebda",
		ResourceTypeRef: MachineResourceTypeRef{
			ID:    "Infrastructure.Virtual",
			Label: "Virtual Machine",
		},
		Name:   "vm001451",
		Status: "ACTIVE",
		CatalogItem: MachineCatalogItem{
			ID: "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
		},
		RequestID: "9f3f8c07-16a3-4f68-a34b-73f756ae960a",
		Organization: MachineOrganization{
			TenantRef:    "vsphere.local",
			SubtenantRef: "f04f060d-73be-48a3-b82c-20cb98efd2d2",
		},
		DateCreated: dateCreated,
		LastUpdated: lastUpdated,
		HasLease:    true,
		LeaseForDisplay: map[string]interface{}{
			"amount": 5.0,
			"type":   "timeSpan",
			"unit":   "DAYS",
		},
		Operations: []MachineOperation{{
			Name:           "Connect by Using SSH",
			Description:    "Infrastructure connect using SSH",
			IconID:         "cafe_default_icon_genericResourceOperation",
			Type:           "EXTENSION",
			ID:             "0c369513-715c-42af-a9c7-260031244dae",
			ExtensionID:    "csp.places.iaas.item.window.ConnectViaSsh",
			ProviderTypeID: nil,
			BindingID:      nil,
			HasForm:        false,
			FormScale:      nil,
		}, {
			Name:           "Connect to Remote Console",
			Description:    "Infrastructure connect using remote console",
			IconID:         "cafe_default_icon_genericResourceOperation",
			Type:           "EXTENSION",
			ID:             "fcbf9796-cabe-4d92-ad5c-ebb7843787ab",
			ExtensionID:    "csp.places.iaas.item.window.ConnectViaVmrc",
			ProviderTypeID: nil,
			BindingID:      nil,
			HasForm:        false,
			FormScale:      nil,
		}},
		ResourceData: MachineResourceData{
			Entries: []MachineDataEntry{
				{
					Key: "NETWORK_LIST",
					Value: map[string]interface{}{
						"items": []interface{}{
							map[string]interface{}{
								"typeFilter": interface{}(nil), "values": map[string]interface{}{
									"entries": []interface{}{
										map[string]interface{}{
											"key": "NETWORK_ADDRESS", "value": map[string]interface{}{
												"type": "string", "value": "10.10.10.20",
											},
										}, map[string]interface{}{
											"key": "NETWORK_MAC_ADDRESS",
											"value": map[string]interface{}{
												"type":  "string",
												"value": "00:50:50:aa:01:1c",
											},
										}, map[string]interface{}{
											"key": "NETWORK_NAME",
											"value": map[string]interface{}{
												"type":  "string",
												"value": "local.network",
											},
										},
									},
								},
								"type":            "complex",
								"componentTypeId": "com.vmware.csp.component.iaas.proxy.provider",
								"componentId":     interface{}(nil),
								"classId":         "dynamicops.api.model.NetworkViewModel",
							},
						},
						"type":          "multiple",
						"elementTypeId": "COMPLEX",
					},
				},
				{
					Key:   "MachineGroupName",
					Value: map[string]interface{}{"type": "string", "value": "Elements"},
				},
				{
					Key:   "MachineName",
					Value: map[string]interface{}{"type": "string", "value": "vm001234"},
				},
				{
					Key:   "Expire",
					Value: map[string]interface{}{"type": "boolean", "value": true},
				},
			},
		},
	}

	mux.HandleFunc("/catalog-service/api/consumer/requests", func(w http.ResponseWriter, r *http.Request) {

		v := new(MachineCreateRequest)

		err := json.NewDecoder(r.Body).Decode(v)

		if err != nil {
			t.Fatal(err)
		}

		testMethod(t, r, "POST")

		// Expect 201 response back
		w.Header().Set("Location", "/catalog-service/api/consumer/requests/9f3f8c07-16a3-4f68-a34b-73f756ae960a")

		if !reflect.DeepEqual(v, createRequest) {
			t.Errorf("Request body = %+v, expected %+v", v, createRequest)
		}
	})

	mux.HandleFunc("/catalog-service/api/consumer/requests/9f3f8c07-16a3-4f68-a34b-73f756ae960a", func(w http.ResponseWriter, r *http.Request) {

		testMethod(t, r, "GET")
		fmt.Fprint(w, jBlobRequestMachineSuccess)
	})

	mux.HandleFunc("/catalog-service/api/consumer/requests/9f3f8c07-16a3-4f68-a34b-73f756ae960a/resources", func(w http.ResponseWriter, r *http.Request) {

		testMethod(t, r, "GET")

		fmt.Fprint(w, jBlobRequestsResources)
	})

	got, _, err := client.Machine.CreateMachine(createRequest)

	if err != nil {
		t.Errorf("Machine.CreateMachine returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Machine.CreateMachine returned %+v, want %+v", got, want)
	}
}

func TestMachine_Destroy(t *testing.T) {
	setup()
	defer teardown()

	destroyRequest := &vrealizeResourceActionRequest{
		Type: "ResourceActionRequest",
		ResourceRef: vrealizeResourceActionRequestResourceRef{
			ID: "b4ce1b01-365a-4964-963a-90263c16ebda",
		},
		ResourceActionRef: vrealizeResourceActionRequestResourceActionRef{
			ID: "5883bbea-cd9a-4bf3-a2b3-f6cc20135435",
		},
		Organization: vrealizeResourceActionRequestOrganization{
			TenantRef:    "vsphere.local",
			SubtenantRef: "f04f060d-73be-48a3-b82c-20cb98efd2d2",
		},
		State:         "SUBMITTED",
		RequestNumber: 0,
	}

	mux.HandleFunc("/catalog-service/api/consumer/resources/b4ce1b01-365a-4964-963a-90263c16ebda", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, jBlobResource)
	})

	mux.HandleFunc("/catalog-service/api/consumer/requests", func(w http.ResponseWriter, r *http.Request) {

		v := new(vrealizeResourceActionRequest)

		err := json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")

		if err != nil {
			t.Fatal(err)
		}

		// Expect 201 response back
		w.Header().Set("Location", "/catalog-service/api/consumer/requests/b1c22613-7602-4df2-b098-0d2af4d40cc2")

		if !reflect.DeepEqual(v, destroyRequest) {
			t.Errorf("returned %#v, expected %#v", v, destroyRequest)
		}
	})

	mux.HandleFunc("/catalog-service/api/consumer/requests/b1c22613-7602-4df2-b098-0d2af4d40cc2", func(w http.ResponseWriter, r *http.Request) {

		testMethod(t, r, "GET")
		fmt.Fprint(w, jBlobDestroyRequestResponse)
	})

	destroyed, _, err := client.Machine.DestroyMachine("b4ce1b01-365a-4964-963a-90263c16ebda")

	if err != nil {
		t.Errorf("Machine.DestroyMachine returned error: %#v", err)
	}

	if !destroyed {
		t.Errorf("Machine.DestroyMachine returned %+v, want %+v", destroyed, true)
	}
}
