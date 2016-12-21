package govrealize

var jBlobListMachines = `
{
"links": [],
    "content": [{
        "@type": "CatalogResource",
        "id": "536ee25b-94b1-48b1-b5e3-b15c52690c4c",
        "iconId": "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
        "resourceTypeRef": {
            "id": "Infrastructure.Virtual",
            "label": "Virtual Machine"
        },
        "name": "vm001448",
        "description": null,
        "status": "ACTIVE",
        "catalogItem": {
            "id": "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
            "label": "Bitnami nginxstack"
        },
        "requestId": "161c66a0-e300-4703-aaa4-aac1029c9f34",
        "providerBinding": {
            "bindingId": "2cac4e37-ec19-481a-a1ea-7615aef7f8df",
            "providerRef": {
                "id": "f8d54a24-ed6c-46f6-bf21-c82d503c7af8",
                "label": "iaas-service"
            }
        },
        "owners": [{
            "tenantName": "vsphere.local",
            "ref": "user1@example.com",
            "type": "USER",
            "value": "user1"
        }],
        "organization": {
            "tenantRef": "vsphere.local",
            "tenantLabel": "vsphere.local",
            "subtenantRef": "f04f060d-73be-48a3-b82c-20cb98efd2d2",
            "subtenantLabel": "Elements"
        },
        "dateCreated": "2016-07-15T11:14:06.542Z",
        "lastUpdated": "2016-07-15T11:14:09.468Z",
        "hasLease": true,
        "lease": {
            "start": "2016-07-15T11:06:49.000Z",
            "end": "2016-07-20T11:06:49.000Z"
        },
        "leaseForDisplay": {
            "amount": 5,
            "type": "timeSpan",
            "unit": "DAYS"
        },
        "hasCosts": true,
        "costs": {
            "leaseRate": {
                "type": "moneyTimeRate",
                "cost": {
                    "type": "money",
                    "currencyCode": "GBP",
                    "amount": 0.0
                },
                "basis": {
                    "type": "timeSpan",
                    "unit": "DAYS",
                    "amount": 1
                }
            }
        },
        "costToDate": {
            "type": "money",
            "currencyCode": "GBP",
            "amount": 0.0
        },
        "totalCost": {
            "type": "money",
            "currencyCode": "GBP",
            "amount": 0.0
        },
        "childResources": [],
        "operations": [{
            "name": "Connect by Using SSH",
            "description": "Infrastructure connect using SSH",
            "iconId": "cafe_default_icon_genericResourceOperation",
            "type": "EXTENSION",
            "id": "0c369513-715c-42af-a9c7-260031244dae",
            "extensionId": "csp.places.iaas.item.window.ConnectViaSsh",
            "providerTypeId": null,
            "bindingId": null,
            "hasForm": false,
            "formScale": null
        }, {
            "name": "Connect to Remote Console",
            "description": "Infrastructure connect using remote console",
            "iconId": "cafe_default_icon_genericResourceOperation",
            "type": "EXTENSION",
            "id": "fcbf9796-cabe-4d92-ad5c-ebb7843787ab",
            "extensionId": "csp.places.iaas.item.window.ConnectViaVmrc",
            "providerTypeId": null,
            "bindingId": null,
            "hasForm": false,
            "formScale": null
        }, {
      		"name": "Destroy",
      		"description": "Destroy a virtual machine.",
      		"iconId": "virtualDestroy.png",
      		"type": "ACTION",
      		"id": "5883bbea-cd9a-4bf3-a2b3-f6cc20135435",
      		"extensionId": null,
      		"providerTypeId": "com.vmware.csp.iaas.blueprint.service",
      		"bindingId": "Infrastructure.Virtual.Action.Destroy",
      		"hasForm": false,
      		"formScale": null
      	}, {
      		"name": "Reconfigure",
      		"description": "Reconfigure a machine.",
      		"iconId": "machineReconfigure.png",
      		"type": "ACTION",
      		"id": "4ae8147e-ce45-4b25-a8a7-3c0f02281363",
      		"extensionId": null,
      		"providerTypeId": "com.vmware.csp.iaas.blueprint.service",
      		"bindingId": "Infrastructure.Machine.Action.Reconfigure",
      		"hasForm": true,
      		"formScale": "BIG"
      	}],
        "forms": {
            "catalogResourceInfoHidden": true,
            "details": {
                "type": "extension",
                "extensionId": "csp.places.iaas.item.details",
                "extensionPointId": null
            }
        },
        "resourceData": {
            "entries": [{
				"key": "NETWORK_LIST",
				"value": {
					"type": "multiple",
					"elementTypeId": "COMPLEX",
					"items": [{
						"type": "complex",
						"componentTypeId": "com.vmware.csp.component.iaas.proxy.provider",
						"componentId": null,
						"classId": "dynamicops.api.model.NetworkViewModel",
						"typeFilter": null,
						"values": {
							"entries": [{
								"key": "NETWORK_ADDRESS",
								"value": {
									"type": "string",
									"value": "10.10.10.20"
								}
							}, {
								"key": "NETWORK_MAC_ADDRESS",
								"value": {
									"type": "string",
									"value": "00:50:50:aa:01:1c"
								}
							}, {
								"key": "NETWORK_NAME",
								"value": {
									"type": "string",
									"value": "local.network"
								}
							}]
						}
					}]
				}
			}, {
                "key": "MachineGroupName",
                "value": {
                    "value": "Elements",
                    "type": "string"
                }
            }, {
                "key": "MachineName",
                "value": {
                    "value": "vm001234",
                    "type": "string"
                }
            }, {
                "key": "Expire",
                "value": {
                    "type": "boolean",
                    "value": true
                }
            }]
        }
    },
    {
        "@type": "CatalogResource",
        "id": "96d414c6-295e-4e3a-ac59-eb9456c1e1d1",
        "iconId": "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
        "resourceTypeRef": {
            "id": "Infrastructure.Virtual",
            "label": "Virtual Machine"
        },
        "name": "vm001448",
        "description": null,
        "status": "ACTIVE",
        "catalogItem": {
            "id": "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
            "label": "Bitnami nginxstack"
        },
        "requestId": "161c66a0-e300-4703-aaa4-aac1029c9f34",
        "providerBinding": {
            "bindingId": "2cac4e37-ec19-481a-a1ea-7615aef7f8df",
            "providerRef": {
                "id": "f8d54a24-ed6c-46f6-bf21-c82d503c7af8",
                "label": "iaas-service"
            }
        },
        "owners": [{
            "tenantName": "vsphere.local",
            "ref": "user1@example.com",
            "type": "USER",
            "value": "user1"
        }],
        "organization": {
            "tenantRef": "vsphere.local",
            "tenantLabel": "vsphere.local",
            "subtenantRef": "f04f060d-73be-48a3-b82c-20cb98efd2d2",
            "subtenantLabel": "Elements"
        },
        "dateCreated": "2016-07-15T11:14:06.542Z",
        "lastUpdated": "2016-07-15T11:14:09.468Z",
        "hasLease": true,
        "lease": {
            "start": "2016-07-15T11:06:49.000Z",
            "end": "2016-07-20T11:06:49.000Z"
        },
        "leaseForDisplay": {
        "amount": 5,
        "type": "timeSpan",
        "unit": "DAYS"
        },
        "hasCosts": true,
        "costs": {
            "leaseRate": {
                "type": "moneyTimeRate",
                "cost": {
                    "type": "money",
                    "currencyCode": "GBP",
                    "amount": 0.0
                },
                "basis": {
                    "type": "timeSpan",
                    "unit": "DAYS",
                    "amount": 1
                }
            }
        },
        "costToDate": {
            "type": "money",
            "currencyCode": "GBP",
            "amount": 0.0
        },
        "totalCost": {
            "type": "money",
            "currencyCode": "GBP",
            "amount": 0.0
        },
        "childResources": [],
        "operations": [{
            "name": "Connect by Using SSH",
            "description": "Infrastructure connect using SSH",
            "iconId": "cafe_default_icon_genericResourceOperation",
            "type": "EXTENSION",
            "id": "0c369513-715c-42af-a9c7-260031244dae",
            "extensionId": "csp.places.iaas.item.window.ConnectViaSsh",
            "providerTypeId": null,
            "bindingId": null,
            "hasForm": false,
            "formScale": null
        }, {
            "name": "Connect to Remote Console",
            "description": "Infrastructure connect using remote console",
            "iconId": "cafe_default_icon_genericResourceOperation",
            "type": "EXTENSION",
            "id": "fcbf9796-cabe-4d92-ad5c-ebb7843787ab",
            "extensionId": "csp.places.iaas.item.window.ConnectViaVmrc",
            "providerTypeId": null,
            "bindingId": null,
            "hasForm": false,
            "formScale": null
        }, {
      		"name": "Destroy",
      		"description": "Destroy a virtual machine.",
      		"iconId": "virtualDestroy.png",
      		"type": "ACTION",
      		"id": "5883bbea-cd9a-4bf3-a2b3-f6cc20135436",
      		"extensionId": null,
      		"providerTypeId": "com.vmware.csp.iaas.blueprint.service",
      		"bindingId": "Infrastructure.Virtual.Action.Destroy",
      		"hasForm": false,
      		"formScale": null
      	}, {
      		"name": "Reconfigure",
      		"description": "Reconfigure a machine.",
      		"iconId": "machineReconfigure.png",
      		"type": "ACTION",
      		"id": "4ae8147e-ce45-4b25-a8a7-3c0f02281364",
      		"extensionId": null,
      		"providerTypeId": "com.vmware.csp.iaas.blueprint.service",
      		"bindingId": "Infrastructure.Machine.Action.Reconfigure",
      		"hasForm": true,
      		"formScale": "BIG"
      	}],
        "forms": {
            "catalogResourceInfoHidden": true,
            "details": {
                "type": "extension",
                "extensionId": "csp.places.iaas.item.details",
                "extensionPointId": null
            }
        },
        "resourceData": {
            "entries": [
            {
				"key": "NETWORK_LIST",
				"value": {
					"type": "multiple",
					"elementTypeId": "COMPLEX",
					"items": [{
						"type": "complex",
						"componentTypeId": "com.vmware.csp.component.iaas.proxy.provider",
						"componentId": null,
						"classId": "dynamicops.api.model.NetworkViewModel",
						"typeFilter": null,
						"values": {
							"entries": [{
								"key": "NETWORK_ADDRESS",
								"value": {
									"type": "string",
									"value": "10.70.10.10"
								}
							}, {
								"key": "NETWORK_MAC_ADDRESS",
								"value": {
									"type": "string",
									"value": "00:50:50:aa:01:1c"
								}
							}, {
								"key": "NETWORK_NAME",
								"value": {
									"type": "string",
									"value": "local.network"
								}
							}]
						}
					}]
				}
			}, {
                "key": "MachineGroupName",
                "value": {
                    "value": "Elements",
                    "type": "string"
                }
            }, {
                "key": "MachineName",
                "value": {
                    "value": "vm005678",
                    "type": "string"
                }
            }, {
                "key": "Expire",
                "value": {
                    "value": true,
                    "type": "boolean"
                }
            }]
        }
    }],
    "metadata": {
        "size": 20,
        "totalElements": 2,
        "totalPages": 1,
        "number": 1,
        "offset": 0
    }
}`

var jBlobRequestMachineSuccess = `{
    "@type": "CatalogItemRequest",
    "id": "9f3f8c07-16a3-4f68-a34b-73f756ae960a",
    "iconId": "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
    "version": 2,
    "requestNumber": 58950,
    "state": "IN_PROGRESS",
    "description": null,
    "reasons": "showcasing API",
    "requestedFor": "user1@example.com",
    "requestedBy": "user1@example.com",
    "organization": {
        "tenantRef": "vsphere.local",
        "tenantLabel": "vsphere.local",
        "subtenantRef": "f04f060d-73be-48a3-b82c-20cb98efd2d2",
        "subtenantLabel": "Elements"
    },
    "requestorEntitlementId": "1dd804b4-4aa6-4d9c-8e17-64724ca9c837",
    "preApprovalId": null,
    "postApprovalId": null,
    "dateCreated": "2016-07-15T12:48:52.150Z",
    "lastUpdated": "2016-07-15T12:48:54.394Z",
    "dateSubmitted": "2016-07-15T12:48:52.150Z",
    "dateApproved": null,
    "dateCompleted": null,
    "quote": {
        "leasePeriod": {
            "type": "timeSpan",
            "unit": "DAYS",
            "amount": 5
        }
    },
    "requestCompletion": {
        "requestCompletionState": "SUCCESSFUL",
        "completionDetails": "The request was successfully completed"
    },
    "requestData": {
        "entries": [{
            "key": "provider-Cafe.Shim.VirtualMachine.Description",
            "value": {
                "type": "string",
                "value": "Test API request"
            }
        }, {
            "key": "provider-provisioningGroupId",
            "value": {
                "type": "string",
                "value": "f04f060d-73be-48a3-b82c-20cb98efd2d2"
            }
        }, {
            "key": "provider-VirtualMachine.LeaseDays",
            "value": {
                "type": "integer",
                "value": 5
            }
        }, {
            "key": "provider-VirtualMachine.Memory.Size",
            "value": {
                "type": "integer",
                "value": 1024
            }
        }, {
            "key": "provider-Cafe.Shim.VirtualMachine.NumberOfInstances",
            "value": {
                "type": "integer",
                "value": 1
            }
        }, {
            "key": "provider-VirtualMachine.CPU.Count",
            "value": {
                "type": "integer",
                "value": 1
            }
        }]
    },
    "retriesRemaining": 3,
    "requestedItemName": "Bitnami nginxstack",
    "requestedItemDescription": "A free, time-limited VM for development and testing purposes that you have full administrative access to and that you must administer yourself. It is provided without support and is not backed up or monitored.  It will be permanently deleted after 30 days.",
    "stateName": "In Progress",
    "phase": "IN_PROGRESS",
    "approvalStatus": "PRE_APPROVED",
    "waitingStatus": "WAITING_FOR_PROVIDER",
    "executionStatus": "STARTED",
    "catalogItemRef": {
        "id": "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
        "label": "Bitnami nginxstack"
    }
}`

var jBlobResource = `{
    "id": "b4ce1b01-365a-4964-963a-90263c16ebda",
    "iconId": "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
    "resourceTypeRef": {
        "id": "Infrastructure.Virtual",
        "label": "Virtual Machine"
    },
    "name": "vm001451",
    "description": null,
    "status": "ACTIVE",
    "catalogItem": {
        "id": "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
        "label": "Bitnami nginxstack"
    },
    "requestId": "9f3f8c07-16a3-4f68-a34b-73f756ae960a",
    "providerBinding": {
        "bindingId": "9cb3d68b-df2d-424d-a82f-881ddb4d0701",
        "providerRef": {
            "id": "f8d54a24-ed6c-46f6-bf21-c82d503c7af8",
            "label": "iaas-service"
        }
    },
    "owners": [{
        "tenantName": "vsphere.local",
        "ref": "user1@example.com",
        "type": "USER",
        "value": "user1"
    }],
    "organization": {
        "tenantRef": "vsphere.local",
        "tenantLabel": "vsphere.local",
        "subtenantRef": "f04f060d-73be-48a3-b82c-20cb98efd2d2",
        "subtenantLabel": "Elements"
    },
    "dateCreated": "2016-07-15T12:55:28.369Z",
    "lastUpdated": "2016-07-15T12:55:32.507Z",
    "hasLease": true,
    "lease": {
        "start": "2016-07-15T12:49:01.000Z",
        "end": "2016-07-20T12:49:01.000Z"
    },
    "leaseForDisplay": {
        "type": "timeSpan",
        "unit": "DAYS",
        "amount": 5
    },
    "hasCosts": true,
    "costs": {
        "leaseRate": {
            "type": "moneyTimeRate",
            "cost": {
                "type": "money",
                "currencyCode": "GBP",
                "amount": 0.0
            },
            "basis": {
                "type": "timeSpan",
                "unit": "DAYS",
                "amount": 1
            }
        }
    },
    "costToDate": {
        "type": "money",
        "currencyCode": "GBP",
        "amount": 0.0
    },
    "totalCost": {
        "type": "money",
        "currencyCode": "GBP",
        "amount": 0.0
    },
    "childResources": [],
    "operations": [{
        "name": "Connect by Using SSH",
        "description": "Infrastructure connect using SSH",
        "iconId": "cafe_default_icon_genericResourceOperation",
        "type": "EXTENSION",
        "id": "0c369513-715c-42af-a9c7-260031244dae",
        "extensionId": "csp.places.iaas.item.window.ConnectViaSsh",
        "providerTypeId": null,
        "bindingId": null,
        "hasForm": false,
        "formScale": null
    }, {
        "name": "Connect to Remote Console",
        "description": "Infrastructure connect using remote console",
        "iconId": "cafe_default_icon_genericResourceOperation",
        "type": "EXTENSION",
        "id": "fcbf9796-cabe-4d92-ad5c-ebb7843787ab",
        "extensionId": "csp.places.iaas.item.window.ConnectViaVmrc",
        "providerTypeId": null,
        "bindingId": null,
        "hasForm": false,
        "formScale": null
    }, {
  		"name": "Destroy",
  		"description": "Destroy a virtual machine.",
  		"iconId": "virtualDestroy.png",
  		"type": "ACTION",
  		"id": "5883bbea-cd9a-4bf3-a2b3-f6cc20135435",
  		"extensionId": null,
  		"providerTypeId": "com.vmware.csp.iaas.blueprint.service",
  		"bindingId": "Infrastructure.Virtual.Action.Destroy",
  		"hasForm": false,
  		"formScale": null
	}, {
    "name": "Reconfigure",
    "description": "Reconfigure a machine.",
    "iconId": "machineReconfigure.png",
    "type": "ACTION",
    "id": "4ae8147e-ce45-4b25-a8a7-3c0f02281363",
    "extensionId": null,
    "providerTypeId": "com.vmware.csp.iaas.blueprint.service",
    "bindingId": "Infrastructure.Machine.Action.Reconfigure",
    "hasForm": true,
    "formScale": "BIG"
  }],
    "forms": {
        "catalogResourceInfoHidden": true,
        "details": {
            "type": "extension",
            "extensionId": "csp.places.iaas.item.details",
            "extensionPointId": null
        }
    },
    "resourceData": {
        "entries": [{
				"key": "NETWORK_LIST",
				"value": {
					"type": "multiple",
					"elementTypeId": "COMPLEX",
					"items": [{
						"type": "complex",
						"componentTypeId": "com.vmware.csp.component.iaas.proxy.provider",
						"componentId": null,
						"classId": "dynamicops.api.model.NetworkViewModel",
						"typeFilter": null,
						"values": {
							"entries": [{
								"key": "NETWORK_ADDRESS",
								"value": {
									"type": "string",
									"value": "10.10.10.20"
								}
							}, {
								"key": "NETWORK_MAC_ADDRESS",
								"value": {
									"type": "string",
									"value": "00:50:50:aa:01:1c"
								}
							}, {
								"key": "NETWORK_NAME",
								"value": {
									"type": "string",
									"value": "local.network"
								}
							}]
						}
					}]
				}
			},
        {
            "key": "MachineGroupName",
            "value": {
                "type": "string",
                "value": "Elements"
            }
        },
        {
            "key": "MachineName",
            "value": {
                "type": "string",
                "value": "vm001234"
            }
        }, {
            "key": "Expire",
            "value": {
                "type": "boolean",
                "value": true
            }
        }]
    }
}`

var jBlobRequestsResources = `{
    "links": [],
    "content": [{
        "@type": "CatalogResource",
        "id": "b4ce1b01-365a-4964-963a-90263c16ebda",
        "iconId": "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
        "resourceTypeRef": {
            "id": "Infrastructure.Virtual",
            "label": "Virtual Machine"
        },
        "name": "vm001451",
        "description": null,
        "status": "ACTIVE",
        "catalogItem": {
            "id": "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
            "label": "Bitnami nginxstack"
        },
        "requestId": "9f3f8c07-16a3-4f68-a34b-73f756ae960a",
        "providerBinding": {
            "bindingId": "9cb3d68b-df2d-424d-a82f-881ddb4d0701",
            "providerRef": {
                "id": "f8d54a24-ed6c-46f6-bf21-c82d503c7af8",
                "label": "iaas-service"
            }
        },
        "owners": [{
            "tenantName": "vsphere.local",
            "ref": "user1@example.com",
            "type": "USER",
            "value": "user1"
        }],
        "organization": {
            "tenantRef": "vsphere.local",
            "tenantLabel": "vsphere.local",
            "subtenantRef": "f04f060d-73be-48a3-b82c-20cb98efd2d2",
            "subtenantLabel": "Elements"
        },
        "dateCreated": "2016-07-15T12:55:28.369Z",
        "lastUpdated": "2016-07-15T12:55:32.507Z",
        "hasLease": true,
        "lease": {
            "start": "2016-07-15T12:49:01.000Z",
            "end": "2016-07-20T12:49:01.000Z"
        },
        "leaseForDisplay": {
            "type": "timeSpan",
            "unit": "DAYS",
            "amount": 5
        },
        "hasCosts": true,
        "costs": {
            "leaseRate": {
                "type": "moneyTimeRate",
                "cost": {
                    "type": "money",
                    "currencyCode": "GBP",
                    "amount": 0.0
                },
                "basis": {
                    "type": "timeSpan",
                    "unit": "DAYS",
                    "amount": 1
                }
            }
        },
        "costToDate": {
            "type": "money",
            "currencyCode": "GBP",
            "amount": 0.0
        },
        "totalCost": {
            "type": "money",
            "currencyCode": "GBP",
            "amount": 0.0
        },
        "childResources": [],
        "operations": [{
            "name": "Connect by Using SSH",
            "description": "Infrastructure connect using SSH",
            "iconId": "cafe_default_icon_genericResourceOperation",
            "type": "EXTENSION",
            "id": "0c369513-715c-42af-a9c7-260031244dae",
            "extensionId": "csp.places.iaas.item.window.ConnectViaSsh",
            "providerTypeId": null,
            "bindingId": null,
            "hasForm": false,
            "formScale": null
        }, {
            "name": "Connect to Remote Console",
            "description": "Infrastructure connect using remote console",
            "iconId": "cafe_default_icon_genericResourceOperation",
            "type": "EXTENSION",
            "id": "fcbf9796-cabe-4d92-ad5c-ebb7843787ab",
            "extensionId": "csp.places.iaas.item.window.ConnectViaVmrc",
            "providerTypeId": null,
            "bindingId": null,
            "hasForm": false,
            "formScale": null
        }],
        "forms": {
            "catalogResourceInfoHidden": true,
            "details": {
                "type": "extension",
                "extensionId": "csp.places.iaas.item.details",
                "extensionPointId": null
            }
        },
        "resourceData": {
            "entries": [{
				"key": "NETWORK_LIST",
				"value": {
					"type": "multiple",
					"elementTypeId": "COMPLEX",
					"items": [{
						"type": "complex",
						"componentTypeId": "com.vmware.csp.component.iaas.proxy.provider",
						"componentId": null,
						"classId": "dynamicops.api.model.NetworkViewModel",
						"typeFilter": null,
						"values": {
							"entries": [{
								"key": "NETWORK_ADDRESS",
								"value": {
									"type": "string",
									"value": "10.10.10.20"
								}
							}, {
								"key": "NETWORK_MAC_ADDRESS",
								"value": {
									"type": "string",
									"value": "00:50:50:aa:01:1c"
								}
							}, {
								"key": "NETWORK_NAME",
								"value": {
									"type": "string",
									"value": "local.network"
								}
							}]
						}
					}]
				}
			},{
                "key": "MachineGroupName",
                "value": {
                    "type": "string",
                    "value": "Elements"
                }
            },{
                "key": "MachineName",
                "value": {
                    "type": "string",
                    "value": "vm001234"
                }
            },{
                "key": "Expire",
                "value": {
                    "type": "boolean",
                    "value": true
                }
            }]
        }
    }],
    "metadata": {
        "size": 20,
        "totalElements": 1,
        "totalPages": 1,
        "number": 1,
        "offset": 0
    }
}`

var jBlobDestroyRequestResponse = `{
    "@type": "ResourceActionRequest",
    "id": "b1c22613-7602-4df2-b098-0d2af4d40cc2",
    "iconId": "virtualDestroy.png",
    "version": 5,
    "requestNumber": 3137,
    "state": "SUCCESSFUL",
    "description": null,
    "reasons": null,
    "requestedFor": "user1@example.com",
    "requestedBy": "user1@example.com",
    "organization": {
        "tenantRef": "vsphere.local",
        "tenantLabel": "vsphere.local",
        "subtenantRef": "f04f060d-73be-48a3-b82c-20cb98efd2d2",
        "subtenantLabel": "Elements"
    },
    "requestorEntitlementId": "1dd804b4-4aa6-4d9c-8e17-64724ca9c837",
    "preApprovalId": null,
    "postApprovalId": null,
    "dateCreated": "2015-04-17T16:18:28.461Z",
    "lastUpdated": "2015-04-17T16:18:36.958Z",
    "dateSubmitted": "2015-04-17T16:18:28.461Z",
    "dateApproved": null,
    "dateCompleted": "2015-04-17T16:18:36.958Z",
    "quote": {
        "leasePeriod": {
            "type": "timeSpan",
            "unit": "MILLISECONDS",
            "amount": 1296000000
        },
        "leaseRate": {
            "type": "moneyTimeRate",
            "cost": {
                "type": "money",
                "currencyCode": "USD",
                "amount": 0
            },
            "basis": {
                "type": "timeSpan",
                "unit": "DAYS",
                "amount": 1
            }
        },
        "totalLeaseCost": {
            "type": "money",
            "currencyCode": "USD",
            "amount": 0
        }
    },
    "requestCompletion": {
        "requestCompletionState": "SUCCESSFUL",
        "completionDetails": "The request was successfully completed"
    },
    "requestData": {
        "entries": []
    },
    "retriesRemaining": 3,
    "requestedItemName": "Destroy - ut001165",
    "requestedItemDescription": "{com.vmware.csp.component.iaas.proxy.provider@resource.action.name.desc.virtual.Destroy}",
    "stateName": "Successful",
    "phase": "SUCCESSFUL",
    "approvalStatus": "POST_APPROVED",
    "waitingStatus": "NOT_WAITING",
    "executionStatus": "STOPPED",
    "resourceRef": {
        "id": "b4ce1b01-365a-4964-963a-90263c16ebda",
        "label": "ut001165"
    },
    "resourceActionRef": {
        "id": "5883bbea-cd9a-4bf3-a2b3-f6cc20135435",
        "label": "Destroy"
    }
}`
