package azblob

import (
	"context"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"

	chk "gopkg.in/check.v1" // go get gopkg.in/check.v1
)

func (s *aztestsSuite) TestGetAccountInfo(c *chk.C) {
	sa := getBSU()

	// Ensure the call succeeded. Don't test for specific account properties because we can't/don't want to set account properties.
	sAccInfo, err := sa.GetAccountInfo(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*sAccInfo, chk.Not(chk.DeepEquals), ServiceGetAccountInfoResponse{})

	//Test on a container
	containerClient := sa.NewContainerClient(generateContainerName())
	_, err = containerClient.Create(ctx, nil)
	defer containerClient.Delete(ctx, nil)
	c.Assert(err, chk.IsNil)
	cAccInfo, err := containerClient.GetAccountInfo(ctx)
	c.Assert(err, chk.IsNil)
	c.Assert(*cAccInfo, chk.Not(chk.DeepEquals), ContainerGetAccountInfoResponse{})

	// test on a block blob URL. They all call the same thing on the base URL, so only one test is needed for that.
	blobClient := containerClient.NewBlockBlobClient(generateBlobName())
	_, err = blobClient.Upload(ctx, azcore.NopCloser(strings.NewReader("blah")), nil)
	c.Assert(err, chk.IsNil)
	bAccInfo, err := blobClient.GetAccountInfo(ctx)
	c.Assert(err, chk.IsNil)
	c.Assert(*bAccInfo, chk.Not(chk.DeepEquals), BlobGetAccountInfoResponse{})
}

func (s *aztestsSuite) TestListContainersBasic(c *chk.C) {
	sa := getBSU()
	prefix := containerPrefix
	listOptions := ListContainersSegmentOptions{Prefix: &prefix}
	pager, err := sa.ListContainersSegment(&listOptions)
	c.Check(err, chk.IsNil)

	for pager.NextPage(context.Background()) {
		page := pager.PageResponse()
		c.Assert(page.RawResponse.StatusCode, chk.Equals, 200)
		c.Assert(page.RequestId, chk.Not(chk.Equals), "")
		c.Assert(page.Version, chk.Not(chk.Equals), "")
		c.Assert(len(*page.EnumerationResults.ContainerItems) >= 0, chk.Equals, true)
		c.Assert(page.EnumerationResults.ServiceEndpoint, chk.NotNil)
	}

	// ensure no error occurred
	c.Assert(pager.Err(), chk.IsNil)

	container, name := getContainerClient(c, sa)
	defer deleteContainer(c, container)

	md := map[string]string{
		"foo": "foovalue",
		"bar": "barvalue",
	}

	options := CreateContainerOptions{Metadata: &md}
	_, err = container.Create(ctx, &options)
	c.Assert(err, chk.IsNil)

	listOptions = ListContainersSegmentOptions{Include: ListContainersDetail{Metadata: true}, Prefix: &name}
	pager, err = sa.ListContainersSegment(&listOptions)
	c.Assert(err, chk.IsNil)

	pager.NextPage(ctx)
	page := pager.PageResponse()
	items := *page.EnumerationResults.ContainerItems
	c.Assert(items, chk.HasLen, 1)
	c.Assert(items[0].Name, chk.NotNil)
	c.Assert(items[0].Properties, chk.NotNil)
	c.Assert(items[0].Properties.LastModified, chk.NotNil)
	c.Assert(items[0].Properties.Etag, chk.NotNil)
	c.Assert(*items[0].Properties.LeaseStatus, chk.Equals, LeaseStatusTypeUnlocked)
	c.Assert(*items[0].Properties.LeaseState, chk.Equals, LeaseStateTypeAvailable)
	c.Assert(items[0].Properties.LeaseDuration, chk.IsNil)
	c.Assert(items[0].Properties.PublicAccess, chk.IsNil)
	c.Assert(*items[0].Metadata, chk.DeepEquals, md)

	// ensure no error occurred
	c.Assert(pager.Err(), chk.IsNil)
}

func (s *aztestsSuite) TestListContainersPaged(c *chk.C) {
	sa := getBSU()

	const numContainers = 6
	maxResults := int32(2)
	const pagedContainersPrefix = "azcontainerpaged"

	containers := make([]ContainerClient, numContainers)
	expectedResults := make(map[string]bool)
	for i := 0; i < numContainers; i++ {
		containerClient, containerName := createNewContainerWithSuffix(c, sa, pagedContainersPrefix)
		containers[i] = containerClient
		expectedResults[containerName] = false
	}

	defer func() {
		for i := range containers {
			deleteContainer(c, containers[i])
		}
	}()

	// list for a first time
	prefix := containerPrefix + pagedContainersPrefix
	listOptions := ListContainersSegmentOptions{MaxResults: &maxResults, Prefix: &prefix}
	pager, err := sa.ListContainersSegment(&listOptions)
	c.Assert(err, chk.IsNil)
	count := 0
	results := make([]ContainerItem, 0)

	for pager.NextPage(context.Background()) {
		page := pager.PageResponse()
		c.Assert(page.RawResponse.StatusCode, chk.Equals, 200)
		c.Assert(page.RequestId, chk.Not(chk.Equals), "")
		c.Assert(page.Version, chk.Not(chk.Equals), "")
		c.Assert(len(*page.EnumerationResults.ContainerItems) <= int(maxResults), chk.Equals, true)
		c.Assert(page.EnumerationResults.ServiceEndpoint, chk.NotNil)

		if count == 0 {
			c.Assert(page.EnumerationResults.Marker, chk.IsNil)
		} else {
			c.Assert(page.EnumerationResults.Marker, chk.NotNil)
		}

		// record the results
		results = append(results, *page.EnumerationResults.ContainerItems...)
		count += 1
	}

	c.Assert(count, chk.Equals, 3)
	c.Assert(len(results), chk.Equals, numContainers)
	c.Assert(pager.Err(), chk.IsNil)

	// make sure each container we see is expected
	for _, container := range results {
		_, ok := expectedResults[*container.Name]
		c.Assert(ok, chk.Equals, true)

		expectedResults[*container.Name] = true
	}

	// make sure every expected container was seen
	for _, seen := range expectedResults {
		c.Assert(seen, chk.Equals, true)
	}
}

func (s *aztestsSuite) TestAccountListContainersEmptyPrefix(c *chk.C) {
	bsu := getBSU()
	containerURL1, _ := createNewContainer(c, bsu)
	defer deleteContainer(c, containerURL1)
	containerURL2, _ := createNewContainer(c, bsu)
	defer deleteContainer(c, containerURL2)

	pager, err := bsu.ListContainersSegment(nil)
	c.Assert(err, chk.IsNil)
	hasPage := pager.NextPage(context.Background())
	c.Assert(hasPage, chk.Equals, true)

	c.Assert(len(*pager.PageResponse().EnumerationResults.ContainerItems) >= 2, chk.Equals, true) // The response should contain at least the two created containers. Probably many more
	c.Assert(pager.Err(), chk.IsNil)
}

func (s *aztestsSuite) TestAccountListContainersMaxResultsNegative(c *chk.C) {
	bsu := getBSU()
	containerURL, _ := createNewContainer(c, bsu)
	defer deleteContainer(c, containerURL)

	illegalMaxResults := []int32{-2, 0}
	for _, num := range illegalMaxResults {
		options := ListContainersSegmentOptions{MaxResults: &num}

		// getting the pager should still work
		pager, err := bsu.ListContainersSegment(&options)
		c.Assert(err, chk.IsNil)

		// getting the next page should fail
		hasPage := pager.NextPage(context.Background())
		c.Assert(hasPage, chk.Equals, false)

		// the error (illegal parameter should have been reported)
		err = pager.Err()
		c.Assert(err, chk.NotNil)
	}
}

func (s *aztestsSuite) TestAccountListContainersMaxResultsExact(c *chk.C) {
	// If this test fails, ensure there are no extra containers prefixed with go in the account. These may be left over if a test is interrupted.
	bsu := getBSU()
	containerURL1, containerName1 := createNewContainerWithSuffix(c, bsu, "abc")
	defer deleteContainer(c, containerURL1)
	containerURL2, containerName2 := createNewContainerWithSuffix(c, bsu, "abcde")
	defer deleteContainer(c, containerURL2)

	prefix := containerPrefix + "abc"
	maxResults := int32(2)
	options := ListContainersSegmentOptions{Prefix: &prefix, MaxResults: &maxResults}
	pager, err := bsu.ListContainersSegment(&options)
	c.Assert(err, chk.IsNil)

	// getting the next page should work
	hasPage := pager.NextPage(context.Background())
	c.Assert(hasPage, chk.Equals, true)

	page := pager.PageResponse()
	c.Assert(err, chk.IsNil)
	c.Assert(*page.EnumerationResults.ContainerItems, chk.HasLen, 2)
	c.Assert(*(*page.EnumerationResults.ContainerItems)[0].Name, chk.DeepEquals, containerName1)
	c.Assert(*(*page.EnumerationResults.ContainerItems)[1].Name, chk.DeepEquals, containerName2)
}

func (s *aztestsSuite) TestAccountDeleteRetentionPolicy(c *chk.C) {
	bsu := getBSU()

	days := int32(5)
	enabled := true
	_, err := bsu.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &enabled, Days: &days}})
	c.Assert(err, chk.IsNil)

	// From FE, 30 seconds is guaranteed to be enough.
	time.Sleep(time.Second * 30)

	resp, err := bsu.GetProperties(ctx)
	c.Assert(err, chk.IsNil)
	c.Assert(*resp.StorageServiceProperties.DeleteRetentionPolicy.Enabled, chk.DeepEquals, true)
	c.Assert(*resp.StorageServiceProperties.DeleteRetentionPolicy.Days, chk.DeepEquals, int32(5))

	disabled := false
	_, err = bsu.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &disabled}})
	c.Assert(err, chk.IsNil)

	// From FE, 30 seconds is guaranteed to be enough.
	time.Sleep(time.Second * 30)

	resp, err = bsu.GetProperties(ctx)
	c.Assert(err, chk.IsNil)
	c.Assert(*resp.StorageServiceProperties.DeleteRetentionPolicy.Enabled, chk.DeepEquals, false)
	c.Assert(resp.StorageServiceProperties.DeleteRetentionPolicy.Days, chk.IsNil)
}

func (s *aztestsSuite) TestAccountDeleteRetentionPolicyEmpty(c *chk.C) {
	bsu := getBSU()

	days := int32(5)
	enabled := true
	_, err := bsu.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &enabled, Days: &days}})
	c.Assert(err, chk.IsNil)

	// From FE, 30 seconds is guaranteed to be enough.
	time.Sleep(time.Second * 30)

	resp, err := bsu.GetProperties(ctx)
	c.Assert(err, chk.IsNil)
	c.Assert(*resp.StorageServiceProperties.DeleteRetentionPolicy.Enabled, chk.DeepEquals, true)
	c.Assert(*resp.StorageServiceProperties.DeleteRetentionPolicy.Days, chk.DeepEquals, int32(5))

	// Empty retention policy causes an error, this is different from track 1.5
	_, err = bsu.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{}})
	c.Assert(err, chk.NotNil)
}

func (s *aztestsSuite) TestAccountDeleteRetentionPolicyNil(c *chk.C) {
	bsu := getBSU()

	days := int32(5)
	enabled := true
	_, err := bsu.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &enabled, Days: &days}})
	c.Assert(err, chk.IsNil)

	// From FE, 30 seconds is guaranteed to be enough.
	time.Sleep(time.Second * 30)

	resp, err := bsu.GetProperties(ctx)
	c.Assert(err, chk.IsNil)
	c.Assert(*resp.StorageServiceProperties.DeleteRetentionPolicy.Enabled, chk.DeepEquals, true)
	c.Assert(*resp.StorageServiceProperties.DeleteRetentionPolicy.Days, chk.DeepEquals, int32(5))

	_, err = bsu.SetProperties(ctx, StorageServiceProperties{})
	c.Assert(err, chk.IsNil)

	// From FE, 30 seconds is guaranteed to be enough.
	time.Sleep(time.Second * 30)

	// If an element of service properties is not passed, the service keeps the current settings.
	resp, err = bsu.GetProperties(ctx)
	c.Assert(err, chk.IsNil)
	c.Assert(*resp.StorageServiceProperties.DeleteRetentionPolicy.Enabled, chk.DeepEquals, true)
	c.Assert(*resp.StorageServiceProperties.DeleteRetentionPolicy.Days, chk.DeepEquals, int32(5))

	// Disable for other tests
	enabled = false
	bsu.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &enabled}})
}

func (s *aztestsSuite) TestAccountDeleteRetentionPolicyDaysTooSmall(c *chk.C) {
	bsu := getBSU()

	days := int32(0) // Minimum days is 1. Validated on the client.
	enabled := true
	_, err := bsu.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &enabled, Days: &days}})
	c.Assert(err, chk.NotNil)
}

func (s *aztestsSuite) TestAccountDeleteRetentionPolicyDaysTooLarge(c *chk.C) {
	bsu := getBSU()

	days := int32(366) // Max days is 365. Left to the service for validation.
	enabled := true
	_, err := bsu.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &enabled, Days: &days}})
	c.Assert(err, chk.NotNil)

	// TODO the error should have more details, follow up with Joel
	//validateStorageError(c, err, ServiceCodeInvalidXMLDocument)
}

func (s *aztestsSuite) TestAccountDeleteRetentionPolicyDaysOmitted(c *chk.C) {
	bsu := getBSU()

	// Days is required if enabled is true.
	enabled := true
	_, err := bsu.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &enabled}})
	c.Assert(err, chk.NotNil)

	// TODO the error should have more details, follow up with Joel
	//validateStorageError(c, err, ServiceCodeInvalidXMLDocument)
}
