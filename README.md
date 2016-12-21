# govrealize
A Go library for interacting with vRealize Automation 6.2 REST API

## Introduction
govrealize is a Golang client library for interacting with the vRealize REST API detailed here: https://developercenter.vmware.com/web/sdk/6.2.0/vrealize-automation. Currently only the creation, reading and deletion of machines is implemented.

## Usage
Create a vRealize client
```golang
client := m.(*govrealize.Client)
```

Create a machine
```golang
entries := []govrealize.MachineRequestDataEntry{}

mrde := govrealize.MachineRequestDataEntry{
	Key: key,
	Value: govrealize.MachineRequestDataEntryValue{
		Type:  "string",
		Value: "value",
	},
}
entries = append(entries, mrde)

opts := &govrealize.MachineCreateRequest{
  Type: "CatalogItemRequest",
	CatalogItemRef: govrealize.MachineCreateRequestCatalogItemRef{
		ID: "xxxx-xxxxx-xxxxx-xxxx",
	},
	Organization: govrealize.MachineCreateRequestOrganization{
		TenantRef:    "tenant",
		SubtenantRef: "subtenant",
	},
	State:         "SUBMITTED",
	RequestNumber: 0,
	RequestData: govrealize.MachineCreateRequestRequestData{
		Entries: entries,
	},
}

machine, resp, err := client.Machine.CreateMachine(opts)
```

Get a machine
```golang
machine, resp, err := client.Machine.GetMachine("xxxx-xxxx-xxxx-xxxx")
```

Destroy a machine
```golang
destroyed, resp, err := client.Machine.DestroyMachine("xxxx-xxxx-xxxx-xxxx")
```
