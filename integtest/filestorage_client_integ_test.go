// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package integtest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oracle/oci-go-sdk/common"

	"github.com/oracle/oci-go-sdk/filestorage"
)

// sanity test for file storage service as the current plan is rely on Java test for other APIs
func TestIdentityClient_ListFileSystems(t *testing.T) {
	c, clerr := filestorage.NewFileStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	req := filestorage.ListFileSystemsRequest{
		CompartmentId:      common.String(getCompartmentID()),
		AvailabilityDomain: common.String(validAD()),
	}

	resp, err := c.ListFileSystems(context.Background(), req)
	failIfError(t, err)

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.OpcRequestId)
}
