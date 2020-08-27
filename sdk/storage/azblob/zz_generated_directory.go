// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// DirectoryOperations contains the methods for the Directory group.
type DirectoryOperations interface {
	// Create - Create a directory. By default, the destination is overwritten and if the destination already exists and has a lease the lease is broken. This operation supports conditional HTTP requests.  For more information, see [Specifying Conditional Headers for Blob Service Operations](https://docs.microsoft.com/en-us/rest/api/storageservices/specifying-conditional-headers-for-blob-service-operations).  To fail if the destination already exists, use a conditional request with If-None-Match: "*".
	Create(ctx context.Context, directoryCreateOptions *DirectoryCreateOptions, directoryHttpHeaders *DirectoryHttpHeaders, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectoryCreateResponse, error)
	// Delete - Deletes the directory
	Delete(ctx context.Context, recursiveDirectoryDelete bool, directoryDeleteOptions *DirectoryDeleteOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectoryDeleteResponse, error)
	// GetAccessControl - Get the owner, group, permissions, or access control list for a directory.
	GetAccessControl(ctx context.Context, directoryGetAccessControlOptions *DirectoryGetAccessControlOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectoryGetAccessControlResponse, error)
	// Rename - Rename a directory. By default, the destination is overwritten and if the destination already exists and has a lease the lease is broken. This operation supports conditional HTTP requests. For more information, see [Specifying Conditional Headers for Blob Service Operations](https://docs.microsoft.com/en-us/rest/api/storageservices/specifying-conditional-headers-for-blob-service-operations). To fail if the destination already exists, use a conditional request with If-None-Match: "*".
	Rename(ctx context.Context, renameSource string, directoryRenameOptions *DirectoryRenameOptions, directoryHttpHeaders *DirectoryHttpHeaders, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (*DirectoryRenameResponse, error)
	// SetAccessControl - Set the owner, group, permissions, or access control list for a directory.
	SetAccessControl(ctx context.Context, directorySetAccessControlOptions *DirectorySetAccessControlOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectorySetAccessControlResponse, error)
}

// directoryOperations implements the DirectoryOperations interface.
type directoryOperations struct {
	*client
	pathRenameMode *PathRenameMode
}

// Create - Create a directory. By default, the destination is overwritten and if the destination already exists and has a lease the lease is broken. This operation supports conditional HTTP requests.  For more information, see [Specifying Conditional Headers for Blob Service Operations](https://docs.microsoft.com/en-us/rest/api/storageservices/specifying-conditional-headers-for-blob-service-operations).  To fail if the destination already exists, use a conditional request with If-None-Match: "*".
func (client *directoryOperations) Create(ctx context.Context, directoryCreateOptions *DirectoryCreateOptions, directoryHttpHeaders *DirectoryHttpHeaders, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectoryCreateResponse, error) {
	req, err := client.createCreateRequest(directoryCreateOptions, directoryHttpHeaders, leaseAccessConditions, modifiedAccessConditions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createCreateRequest creates the Create request.
func (client *directoryOperations) createCreateRequest(directoryCreateOptions *DirectoryCreateOptions, directoryHttpHeaders *DirectoryHttpHeaders, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	copy := *client.u
	u := &copy
	query := u.Query()
	query.Set("resource", "directory")
	if directoryCreateOptions != nil && directoryCreateOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*directoryCreateOptions.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	if directoryCreateOptions != nil && directoryCreateOptions.DirectoryProperties != nil {
		req.Header.Set("x-ms-properties", *directoryCreateOptions.DirectoryProperties)
	}
	if directoryCreateOptions != nil && directoryCreateOptions.PosixPermissions != nil {
		req.Header.Set("x-ms-permissions", *directoryCreateOptions.PosixPermissions)
	}
	if directoryCreateOptions != nil && directoryCreateOptions.PosixUmask != nil {
		req.Header.Set("x-ms-umask", *directoryCreateOptions.PosixUmask)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.CacheControl != nil {
		req.Header.Set("x-ms-cache-control", *directoryHttpHeaders.CacheControl)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentType != nil {
		req.Header.Set("x-ms-content-type", *directoryHttpHeaders.ContentType)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentEncoding != nil {
		req.Header.Set("x-ms-content-encoding", *directoryHttpHeaders.ContentEncoding)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentLanguage != nil {
		req.Header.Set("x-ms-content-language", *directoryHttpHeaders.ContentLanguage)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentDisposition != nil {
		req.Header.Set("x-ms-content-disposition", *directoryHttpHeaders.ContentDisposition)
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-12-12")
	if directoryCreateOptions != nil && directoryCreateOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *directoryCreateOptions.RequestId)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *directoryOperations) createHandleResponse(resp *azcore.Response) (*DirectoryCreateResponse, error) {
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	result := DirectoryCreateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestId = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Content-Length"); val != "" {
		contentLength, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		result.ContentLength = &contentLength
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	return &result, nil
}

// createHandleError handles the Create error response.
func (client *directoryOperations) createHandleError(resp *azcore.Response) error {
	var err DataLakeStorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the directory
func (client *directoryOperations) Delete(ctx context.Context, recursiveDirectoryDelete bool, directoryDeleteOptions *DirectoryDeleteOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectoryDeleteResponse, error) {
	req, err := client.deleteCreateRequest(recursiveDirectoryDelete, directoryDeleteOptions, leaseAccessConditions, modifiedAccessConditions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// deleteCreateRequest creates the Delete request.
func (client *directoryOperations) deleteCreateRequest(recursiveDirectoryDelete bool, directoryDeleteOptions *DirectoryDeleteOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	copy := *client.u
	u := &copy
	query := u.Query()
	if directoryDeleteOptions != nil && directoryDeleteOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*directoryDeleteOptions.Timeout), 10))
	}
	query.Set("recursive", strconv.FormatBool(recursiveDirectoryDelete))
	if directoryDeleteOptions != nil && directoryDeleteOptions.Marker != nil {
		query.Set("continuation", *directoryDeleteOptions.Marker)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-12-12")
	if directoryDeleteOptions != nil && directoryDeleteOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *directoryDeleteOptions.RequestId)
	}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *directoryOperations) deleteHandleResponse(resp *azcore.Response) (*DirectoryDeleteResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.deleteHandleError(resp)
	}
	result := DirectoryDeleteResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-continuation"); val != "" {
		result.Continuation = &val
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestId = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	return &result, nil
}

// deleteHandleError handles the Delete error response.
func (client *directoryOperations) deleteHandleError(resp *azcore.Response) error {
	var err DataLakeStorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}

// GetAccessControl - Get the owner, group, permissions, or access control list for a directory.
func (client *directoryOperations) GetAccessControl(ctx context.Context, directoryGetAccessControlOptions *DirectoryGetAccessControlOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectoryGetAccessControlResponse, error) {
	req, err := client.getAccessControlCreateRequest(directoryGetAccessControlOptions, leaseAccessConditions, modifiedAccessConditions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getAccessControlHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getAccessControlCreateRequest creates the GetAccessControl request.
func (client *directoryOperations) getAccessControlCreateRequest(directoryGetAccessControlOptions *DirectoryGetAccessControlOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	copy := *client.u
	u := &copy
	query := u.Query()
	query.Set("action", "getAccessControl")
	if directoryGetAccessControlOptions != nil && directoryGetAccessControlOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*directoryGetAccessControlOptions.Timeout), 10))
	}
	if directoryGetAccessControlOptions != nil && directoryGetAccessControlOptions.Upn != nil {
		query.Set("upn", strconv.FormatBool(*directoryGetAccessControlOptions.Upn))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodHead, *u)
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if directoryGetAccessControlOptions != nil && directoryGetAccessControlOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *directoryGetAccessControlOptions.RequestId)
	}
	req.Header.Set("x-ms-version", "2019-12-12")
	return req, nil
}

// getAccessControlHandleResponse handles the GetAccessControl response.
func (client *directoryOperations) getAccessControlHandleResponse(resp *azcore.Response) (*DirectoryGetAccessControlResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getAccessControlHandleError(resp)
	}
	result := DirectoryGetAccessControlResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-owner"); val != "" {
		result.Owner = &val
	}
	if val := resp.Header.Get("x-ms-group"); val != "" {
		result.Group = &val
	}
	if val := resp.Header.Get("x-ms-permissions"); val != "" {
		result.Permissions = &val
	}
	if val := resp.Header.Get("x-ms-acl"); val != "" {
		result.Acl = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, nil
}

// getAccessControlHandleError handles the GetAccessControl error response.
func (client *directoryOperations) getAccessControlHandleError(resp *azcore.Response) error {
	var err DataLakeStorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}

// Rename - Rename a directory. By default, the destination is overwritten and if the destination already exists and has a lease the lease is broken. This operation supports conditional HTTP requests. For more information, see [Specifying Conditional Headers for Blob Service Operations](https://docs.microsoft.com/en-us/rest/api/storageservices/specifying-conditional-headers-for-blob-service-operations). To fail if the destination already exists, use a conditional request with If-None-Match: "*".
func (client *directoryOperations) Rename(ctx context.Context, renameSource string, directoryRenameOptions *DirectoryRenameOptions, directoryHttpHeaders *DirectoryHttpHeaders, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (*DirectoryRenameResponse, error) {
	req, err := client.renameCreateRequest(renameSource, directoryRenameOptions, directoryHttpHeaders, leaseAccessConditions, modifiedAccessConditions, sourceModifiedAccessConditions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.renameHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// renameCreateRequest creates the Rename request.
func (client *directoryOperations) renameCreateRequest(renameSource string, directoryRenameOptions *DirectoryRenameOptions, directoryHttpHeaders *DirectoryHttpHeaders, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (*azcore.Request, error) {
	copy := *client.u
	u := &copy
	query := u.Query()
	if directoryRenameOptions != nil && directoryRenameOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*directoryRenameOptions.Timeout), 10))
	}
	if directoryRenameOptions != nil && directoryRenameOptions.Marker != nil {
		query.Set("continuation", *directoryRenameOptions.Marker)
	}
	if client.pathRenameMode != nil {
		query.Set("mode", string(*client.pathRenameMode))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	req.Header.Set("x-ms-rename-source", renameSource)
	if directoryRenameOptions != nil && directoryRenameOptions.DirectoryProperties != nil {
		req.Header.Set("x-ms-properties", *directoryRenameOptions.DirectoryProperties)
	}
	if directoryRenameOptions != nil && directoryRenameOptions.PosixPermissions != nil {
		req.Header.Set("x-ms-permissions", *directoryRenameOptions.PosixPermissions)
	}
	if directoryRenameOptions != nil && directoryRenameOptions.PosixUmask != nil {
		req.Header.Set("x-ms-umask", *directoryRenameOptions.PosixUmask)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.CacheControl != nil {
		req.Header.Set("x-ms-cache-control", *directoryHttpHeaders.CacheControl)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentType != nil {
		req.Header.Set("x-ms-content-type", *directoryHttpHeaders.ContentType)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentEncoding != nil {
		req.Header.Set("x-ms-content-encoding", *directoryHttpHeaders.ContentEncoding)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentLanguage != nil {
		req.Header.Set("x-ms-content-language", *directoryHttpHeaders.ContentLanguage)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentDisposition != nil {
		req.Header.Set("x-ms-content-disposition", *directoryHttpHeaders.ContentDisposition)
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if directoryRenameOptions != nil && directoryRenameOptions.SourceLeaseId != nil {
		req.Header.Set("x-ms-source-lease-id", *directoryRenameOptions.SourceLeaseId)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfModifiedSince != nil {
		req.Header.Set("x-ms-source-if-modified-since", sourceModifiedAccessConditions.SourceIfModifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfUnmodifiedSince != nil {
		req.Header.Set("x-ms-source-if-unmodified-since", sourceModifiedAccessConditions.SourceIfUnmodifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfMatch != nil {
		req.Header.Set("x-ms-source-if-match", *sourceModifiedAccessConditions.SourceIfMatch)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfNoneMatch != nil {
		req.Header.Set("x-ms-source-if-none-match", *sourceModifiedAccessConditions.SourceIfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-12-12")
	if directoryRenameOptions != nil && directoryRenameOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *directoryRenameOptions.RequestId)
	}
	return req, nil
}

// renameHandleResponse handles the Rename response.
func (client *directoryOperations) renameHandleResponse(resp *azcore.Response) (*DirectoryRenameResponse, error) {
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, client.renameHandleError(resp)
	}
	result := DirectoryRenameResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-continuation"); val != "" {
		result.Continuation = &val
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestId = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Content-Length"); val != "" {
		contentLength, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		result.ContentLength = &contentLength
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	return &result, nil
}

// renameHandleError handles the Rename error response.
func (client *directoryOperations) renameHandleError(resp *azcore.Response) error {
	var err DataLakeStorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}

// SetAccessControl - Set the owner, group, permissions, or access control list for a directory.
func (client *directoryOperations) SetAccessControl(ctx context.Context, directorySetAccessControlOptions *DirectorySetAccessControlOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectorySetAccessControlResponse, error) {
	req, err := client.setAccessControlCreateRequest(directorySetAccessControlOptions, leaseAccessConditions, modifiedAccessConditions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.setAccessControlHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// setAccessControlCreateRequest creates the SetAccessControl request.
func (client *directoryOperations) setAccessControlCreateRequest(directorySetAccessControlOptions *DirectorySetAccessControlOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	copy := *client.u
	u := &copy
	query := u.Query()
	query.Set("action", "setAccessControl")
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*directorySetAccessControlOptions.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.Owner != nil {
		req.Header.Set("x-ms-owner", *directorySetAccessControlOptions.Owner)
	}
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.Group != nil {
		req.Header.Set("x-ms-group", *directorySetAccessControlOptions.Group)
	}
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.PosixPermissions != nil {
		req.Header.Set("x-ms-permissions", *directorySetAccessControlOptions.PosixPermissions)
	}
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.PosixAcl != nil {
		req.Header.Set("x-ms-acl", *directorySetAccessControlOptions.PosixAcl)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *directorySetAccessControlOptions.RequestId)
	}
	req.Header.Set("x-ms-version", "2019-12-12")
	return req, nil
}

// setAccessControlHandleResponse handles the SetAccessControl response.
func (client *directoryOperations) setAccessControlHandleResponse(resp *azcore.Response) (*DirectorySetAccessControlResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.setAccessControlHandleError(resp)
	}
	result := DirectorySetAccessControlResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, nil
}

// setAccessControlHandleError handles the SetAccessControl error response.
func (client *directoryOperations) setAccessControlHandleError(resp *azcore.Response) error {
	var err DataLakeStorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}