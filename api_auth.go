// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/vault-client-go/schema"
)

// Auth is a simple wrapper around the client for Auth requests
type Auth struct {
	client *Client
}

// AWSConfigDeleteCertificate
// certName: Name of the certificate.
func (a *Auth) AWSConfigDeleteCertificate(ctx context.Context, certName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificate/{cert_name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"cert_name"+"}", url.PathEscape(certName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigDeleteClient
func (a *Auth) AWSConfigDeleteClient(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/client"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigDeleteIdentityAccessList
func (a *Auth) AWSConfigDeleteIdentityAccessList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigDeleteIdentityWhiteList
func (a *Auth) AWSConfigDeleteIdentityWhiteList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigDeleteRoleTagBlackList
func (a *Auth) AWSConfigDeleteRoleTagBlackList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigDeleteRoleTagDenyList
func (a *Auth) AWSConfigDeleteRoleTagDenyList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigDeleteSecurityTokenServiceAccount
// accountId: AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
func (a *Auth) AWSConfigDeleteSecurityTokenServiceAccount(ctx context.Context, accountId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts/{account_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"account_id"+"}", url.PathEscape(accountId), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigListCertificates
func (a *Auth) AWSConfigListCertificates(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificates"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigListSecurityTokenService
func (a *Auth) AWSConfigListSecurityTokenService(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigReadCertificate
// certName: Name of the certificate.
func (a *Auth) AWSConfigReadCertificate(ctx context.Context, certName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificate/{cert_name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"cert_name"+"}", url.PathEscape(certName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigReadClient
func (a *Auth) AWSConfigReadClient(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/client"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigReadIdentity
func (a *Auth) AWSConfigReadIdentity(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/identity"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigReadIdentityAccessList
func (a *Auth) AWSConfigReadIdentityAccessList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigReadIdentityWhiteList
func (a *Auth) AWSConfigReadIdentityWhiteList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigReadRoleTagBlackList
func (a *Auth) AWSConfigReadRoleTagBlackList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigReadRoleTagDenyList
func (a *Auth) AWSConfigReadRoleTagDenyList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigReadSecurityTokenServiceAccount
// accountId: AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
func (a *Auth) AWSConfigReadSecurityTokenServiceAccount(ctx context.Context, accountId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts/{account_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"account_id"+"}", url.PathEscape(accountId), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigRotateRoot
func (a *Auth) AWSConfigRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigWriteCertificate
// certName: Name of the certificate.
func (a *Auth) AWSConfigWriteCertificate(ctx context.Context, certName string, request schema.AWSConfigWriteCertificateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificate/{cert_name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"cert_name"+"}", url.PathEscape(certName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigWriteClient
func (a *Auth) AWSConfigWriteClient(ctx context.Context, request schema.AWSConfigWriteClientRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/client"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigWriteIdentity
func (a *Auth) AWSConfigWriteIdentity(ctx context.Context, request schema.AWSConfigWriteIdentityRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/identity"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigWriteIdentityAccessList
func (a *Auth) AWSConfigWriteIdentityAccessList(ctx context.Context, request schema.AWSConfigWriteIdentityAccessListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigWriteIdentityWhiteList
func (a *Auth) AWSConfigWriteIdentityWhiteList(ctx context.Context, request schema.AWSConfigWriteIdentityWhiteListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigWriteRoleTagBlackList
func (a *Auth) AWSConfigWriteRoleTagBlackList(ctx context.Context, request schema.AWSConfigWriteRoleTagBlackListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigWriteRoleTagDenyList
func (a *Auth) AWSConfigWriteRoleTagDenyList(ctx context.Context, request schema.AWSConfigWriteRoleTagDenyListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigWriteSecurityTokenServiceAccount
// accountId: AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
func (a *Auth) AWSConfigWriteSecurityTokenServiceAccount(ctx context.Context, accountId string, request schema.AWSConfigWriteSecurityTokenServiceAccountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts/{account_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"account_id"+"}", url.PathEscape(accountId), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSDeleteAuthRole
// role: Name of the role.
func (a *Auth) AWSDeleteAuthRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSDeleteIdentityAccessListFor
// instanceId: EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
func (a *Auth) AWSDeleteIdentityAccessListFor(ctx context.Context, instanceId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-accesslist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSDeleteIdentityWhiteListFor
// instanceId: EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
func (a *Auth) AWSDeleteIdentityWhiteListFor(ctx context.Context, instanceId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-whitelist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSDeleteRoleTagBlackListFor
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AWSDeleteRoleTagBlackListFor(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSDeleteRoleTagDenyListFor
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AWSDeleteRoleTagDenyListFor(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSListAuthRoles
func (a *Auth) AWSListAuthRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSListAuthRoles2
func (a *Auth) AWSListAuthRoles2(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSListIdentityAccessList
func (a *Auth) AWSListIdentityAccessList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSListIdentityWhiteList
func (a *Auth) AWSListIdentityWhiteList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSListRoleTagBlackList
func (a *Auth) AWSListRoleTagBlackList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSListRoleTagDenyList
func (a *Auth) AWSListRoleTagDenyList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSLogin
func (a *Auth) AWSLogin(ctx context.Context, request schema.AWSLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSReadAuthRole
// role: Name of the role.
func (a *Auth) AWSReadAuthRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSReadIdentityAccessListFor
// instanceId: EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
func (a *Auth) AWSReadIdentityAccessListFor(ctx context.Context, instanceId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-accesslist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSReadIdentityWhiteListFor
// instanceId: EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
func (a *Auth) AWSReadIdentityWhiteListFor(ctx context.Context, instanceId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-whitelist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSReadRoleTagBlackListFor
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AWSReadRoleTagBlackListFor(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSReadRoleTagDenyListFor
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AWSReadRoleTagDenyListFor(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSWriteAuthRole
// role: Name of the role.
func (a *Auth) AWSWriteAuthRole(ctx context.Context, role string, request schema.AWSWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSWriteAuthRoleTag
// role: Name of the role.
func (a *Auth) AWSWriteAuthRoleTag(ctx context.Context, role string, request schema.AWSWriteAuthRoleTagRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}/tag"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSWriteIdentityAccessListTidySettings
func (a *Auth) AWSWriteIdentityAccessListTidySettings(ctx context.Context, request schema.AWSWriteIdentityAccessListTidySettingsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSWriteIdentityWhiteListTidySettings
func (a *Auth) AWSWriteIdentityWhiteListTidySettings(ctx context.Context, request schema.AWSWriteIdentityWhiteListTidySettingsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSWriteRoleTagBlackListFor
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AWSWriteRoleTagBlackListFor(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSWriteRoleTagBlackListTidySettings
func (a *Auth) AWSWriteRoleTagBlackListTidySettings(ctx context.Context, request schema.AWSWriteRoleTagBlackListTidySettingsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSWriteRoleTagDenyListFor
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AWSWriteRoleTagDenyListFor(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSWriteRoleTagDenyListTidySettings
func (a *Auth) AWSWriteRoleTagDenyListTidySettings(ctx context.Context, request schema.AWSWriteRoleTagDenyListTidySettingsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudDeleteAuthRole Create a role and associate policies to it.
// role: The name of the role as it should appear in Vault.
func (a *Auth) AliCloudDeleteAuthRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudListAuthRoles Lists all the roles that are registered with Vault.
func (a *Auth) AliCloudListAuthRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudListAuthRoles2 Lists all the roles that are registered with Vault.
func (a *Auth) AliCloudListAuthRoles2(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudLogin Authenticates an RAM entity with Vault.
func (a *Auth) AliCloudLogin(ctx context.Context, request schema.AliCloudLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudReadAuthRole Create a role and associate policies to it.
// role: The name of the role as it should appear in Vault.
func (a *Auth) AliCloudReadAuthRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudWriteAuthRole Create a role and associate policies to it.
// role: The name of the role as it should appear in Vault.
func (a *Auth) AliCloudWriteAuthRole(ctx context.Context, role string, request schema.AliCloudWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteBindSecretID
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteBindSecretID(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bind-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteBoundCIDRList
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteBoundCIDRList(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bound-cidr-list"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeletePeriod
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeletePeriod(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/period"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeletePolicies
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeletePolicies(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/policies"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteRole
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteRole(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteSecretIDAccessorDestroy
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteSecretIDAccessorDestroy(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteSecretIDBoundCIDRs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteSecretIDBoundCIDRs(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteSecretIDDestroy
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteSecretIDDestroy(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id/destroy"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteSecretIDNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteSecretIDNumUses(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteSecretIDTTL
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteSecretIDTTL(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteTokenBoundCIDRs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteTokenBoundCIDRs(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteTokenMaxTTL
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteTokenMaxTTL(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-max-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteTokenNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteTokenNumUses(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteTokenTTL
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteTokenTTL(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleListRoles
func (a *Auth) AppRoleListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.AppRoleListRolesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[schema.AppRoleListRolesResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleListSecretID
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleListSecretID(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleListSecretIDResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[schema.AppRoleListSecretIDResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleLogin
func (a *Auth) AppRoleLogin(ctx context.Context, request schema.AppRoleLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadBindSecretID
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadBindSecretID(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadBindSecretIDResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bind-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadBindSecretIDResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadBoundCIDRList
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadBoundCIDRList(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadBoundCIDRListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bound-cidr-list"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadBoundCIDRListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadLocalSecretIDs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadLocalSecretIDs(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadLocalSecretIDsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/local-secret-ids"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadLocalSecretIDsResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadPeriod
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadPeriod(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadPeriodResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/period"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadPeriodResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadPolicies
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadPolicies(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadPoliciesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/policies"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadPoliciesResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadRole
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadRole(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadRoleResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadRoleID
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadRoleID(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadRoleIDResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/role-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadRoleIDResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadSecretIDBoundCIDRs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadSecretIDBoundCIDRs(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadSecretIDBoundCIDRsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadSecretIDBoundCIDRsResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadSecretIDNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadSecretIDNumUses(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadSecretIDNumUsesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadSecretIDNumUsesResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadSecretIDTTL
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadSecretIDTTL(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadSecretIDTTLResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadSecretIDTTLResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadTokenBoundCIDRs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadTokenBoundCIDRs(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadTokenBoundCIDRsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadTokenBoundCIDRsResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadTokenMaxTTL
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadTokenMaxTTL(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadTokenMaxTTLResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-max-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadTokenMaxTTLResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadTokenNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadTokenNumUses(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadTokenNumUsesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadTokenNumUsesResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadTokenTTL
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadTokenTTL(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadTokenTTLResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.AppRoleReadTokenTTLResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleTidySecretID Trigger the clean-up of expired SecretID entries.
func (a *Auth) AppRoleTidySecretID(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/tidy/secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteBindSecretID
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteBindSecretID(ctx context.Context, roleName string, request schema.AppRoleWriteBindSecretIDRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bind-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteBoundCIDRList
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteBoundCIDRList(ctx context.Context, roleName string, request schema.AppRoleWriteBoundCIDRListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bound-cidr-list"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteCustomSecretID
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteCustomSecretID(ctx context.Context, roleName string, request schema.AppRoleWriteCustomSecretIDRequest, options ...RequestOption) (*Response[schema.AppRoleWriteCustomSecretIDResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/custom-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.AppRoleWriteCustomSecretIDResponse](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWritePeriod
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWritePeriod(ctx context.Context, roleName string, request schema.AppRoleWritePeriodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/period"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWritePolicies
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWritePolicies(ctx context.Context, roleName string, request schema.AppRoleWritePoliciesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/policies"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteRole
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteRole(ctx context.Context, roleName string, request schema.AppRoleWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteRoleID
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteRoleID(ctx context.Context, roleName string, request schema.AppRoleWriteRoleIDRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/role-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretID
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretID(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIDRequest, options ...RequestOption) (*Response[schema.AppRoleWriteSecretIDResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.AppRoleWriteSecretIDResponse](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretIDAccessorDestroy
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretIDAccessorDestroy(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIDAccessorDestroyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretIDAccessorLookup
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretIDAccessorLookup(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIDAccessorLookupRequest, options ...RequestOption) (*Response[schema.AppRoleWriteSecretIDAccessorLookupResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/lookup"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.AppRoleWriteSecretIDAccessorLookupResponse](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretIDBoundCIDRs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretIDBoundCIDRs(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIDBoundCIDRsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretIDDestroy
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretIDDestroy(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIDDestroyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id/destroy"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretIDLookup
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretIDLookup(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIDLookupRequest, options ...RequestOption) (*Response[schema.AppRoleWriteSecretIDLookupResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id/lookup"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.AppRoleWriteSecretIDLookupResponse](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretIDNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretIDNumUses(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIDNumUsesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretIDTTL
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretIDTTL(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIDTTLRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteTokenBoundCIDRs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteTokenBoundCIDRs(ctx context.Context, roleName string, request schema.AppRoleWriteTokenBoundCIDRsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteTokenMaxTTL
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteTokenMaxTTL(ctx context.Context, roleName string, request schema.AppRoleWriteTokenMaxTTLRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-max-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteTokenNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteTokenNumUses(ctx context.Context, roleName string, request schema.AppRoleWriteTokenNumUsesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteTokenTTL
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteTokenTTL(ctx context.Context, roleName string, request schema.AppRoleWriteTokenTTLRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureDeleteAuthConfig
func (a *Auth) AzureDeleteAuthConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureDeleteAuthRole
// name: Name of the role.
func (a *Auth) AzureDeleteAuthRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureListAuthRoles
func (a *Auth) AzureListAuthRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureLogin
func (a *Auth) AzureLogin(ctx context.Context, request schema.AzureLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureReadAuthConfig
func (a *Auth) AzureReadAuthConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureReadAuthRole
// name: Name of the role.
func (a *Auth) AzureReadAuthRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureWriteAuthConfig
func (a *Auth) AzureWriteAuthConfig(ctx context.Context, request schema.AzureWriteAuthConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureWriteAuthRole
// name: Name of the role.
func (a *Auth) AzureWriteAuthRole(ctx context.Context, name string, request schema.AzureWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CentrifyLogin Log in with a username and password.
func (a *Auth) CentrifyLogin(ctx context.Context, request schema.CentrifyLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{centrify_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"centrify_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("centrify")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CentrifyReadConfig This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
func (a *Auth) CentrifyReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{centrify_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"centrify_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("centrify")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CentrifyWriteConfig This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
func (a *Auth) CentrifyWriteConfig(ctx context.Context, request schema.CentrifyWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{centrify_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"centrify_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("centrify")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CertificatesDelete Manage trusted certificates used for authentication.
// name: The name of the certificate
func (a *Auth) CertificatesDelete(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertificatesDeleteCRL Manage Certificate Revocation Lists checked during authentication.
// name: The name of the certificate
func (a *Auth) CertificatesDeleteCRL(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertificatesList Manage trusted certificates used for authentication.
func (a *Auth) CertificatesList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertificatesListCRLs
func (a *Auth) CertificatesListCRLs(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertificatesLogin
func (a *Auth) CertificatesLogin(ctx context.Context, request schema.CertificatesLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CertificatesRead Manage trusted certificates used for authentication.
// name: The name of the certificate
func (a *Auth) CertificatesRead(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertificatesReadCRL Manage Certificate Revocation Lists checked during authentication.
// name: The name of the certificate
func (a *Auth) CertificatesReadCRL(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertificatesReadConfig
func (a *Auth) CertificatesReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertificatesWrite Manage trusted certificates used for authentication.
// name: The name of the certificate
func (a *Auth) CertificatesWrite(ctx context.Context, name string, request schema.CertificatesWriteRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CertificatesWriteCRL Manage Certificate Revocation Lists checked during authentication.
// name: The name of the certificate
func (a *Auth) CertificatesWriteCRL(ctx context.Context, name string, request schema.CertificatesWriteCRLRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CertificatesWriteConfig
func (a *Auth) CertificatesWriteConfig(ctx context.Context, request schema.CertificatesWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryDeleteConfig
func (a *Auth) CloudFoundryDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryDeleteRole
// role: The name of the role.
func (a *Auth) CloudFoundryDeleteRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryListRoles
func (a *Auth) CloudFoundryListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryLogin
func (a *Auth) CloudFoundryLogin(ctx context.Context, request schema.CloudFoundryLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryReadConfig
func (a *Auth) CloudFoundryReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryReadRole
// role: The name of the role.
func (a *Auth) CloudFoundryReadRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryWriteConfig
func (a *Auth) CloudFoundryWriteConfig(ctx context.Context, request schema.CloudFoundryWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryWriteRole
// role: The name of the role.
func (a *Auth) CloudFoundryWriteRole(ctx context.Context, role string, request schema.CloudFoundryWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GitHubDeleteMapTeam Read/write/delete a single teams mapping
// key: Key for the teams mapping
func (a *Auth) GitHubDeleteMapTeam(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GitHubDeleteMapUser Read/write/delete a single users mapping
// key: Key for the users mapping
func (a *Auth) GitHubDeleteMapUser(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GitHubLogin
func (a *Auth) GitHubLogin(ctx context.Context, request schema.GitHubLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GitHubReadConfig
func (a *Auth) GitHubReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GitHubReadMapTeam Read/write/delete a single teams mapping
// key: Key for the teams mapping
func (a *Auth) GitHubReadMapTeam(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GitHubReadMapTeams Read mappings for teams
func (a *Auth) GitHubReadMapTeams(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GitHubReadMapUser Read/write/delete a single users mapping
// key: Key for the users mapping
func (a *Auth) GitHubReadMapUser(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GitHubReadMapUsers Read mappings for users
func (a *Auth) GitHubReadMapUsers(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GitHubWriteConfig
func (a *Auth) GitHubWriteConfig(ctx context.Context, request schema.GitHubWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GitHubWriteMapTeam Read/write/delete a single teams mapping
// key: Key for the teams mapping
func (a *Auth) GitHubWriteMapTeam(ctx context.Context, key string, request schema.GitHubWriteMapTeamRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GitHubWriteMapUser Read/write/delete a single users mapping
// key: Key for the users mapping
func (a *Auth) GitHubWriteMapUser(ctx context.Context, key string, request schema.GitHubWriteMapUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudDeleteRole Create a GCP role with associated policies and required attributes.
// name: Name of the role.
func (a *Auth) GoogleCloudDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudListRoles Lists all the roles that are registered with Vault.
func (a *Auth) GoogleCloudListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudListRoles2 Lists all the roles that are registered with Vault.
func (a *Auth) GoogleCloudListRoles2(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudLogin
func (a *Auth) GoogleCloudLogin(ctx context.Context, request schema.GoogleCloudLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadAuthConfig Configure credentials used to query the GCP IAM API to verify authenticating service accounts
func (a *Auth) GoogleCloudReadAuthConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadRole Create a GCP role with associated policies and required attributes.
// name: Name of the role.
func (a *Auth) GoogleCloudReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteAuthConfig Configure credentials used to query the GCP IAM API to verify authenticating service accounts
func (a *Auth) GoogleCloudWriteAuthConfig(ctx context.Context, request schema.GoogleCloudWriteAuthConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteRole Create a GCP role with associated policies and required attributes.
// name: Name of the role.
func (a *Auth) GoogleCloudWriteRole(ctx context.Context, name string, request schema.GoogleCloudWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteRoleLabels Add or remove labels for an existing 'gce' role
// name: Name of the role.
func (a *Auth) GoogleCloudWriteRoleLabels(ctx context.Context, name string, request schema.GoogleCloudWriteRoleLabelsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}/labels"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteRoleServiceAccounts Add or remove service accounts for an existing `iam` role
// name: Name of the role.
func (a *Auth) GoogleCloudWriteRoleServiceAccounts(ctx context.Context, name string, request schema.GoogleCloudWriteRoleServiceAccountsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}/service-accounts"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// JWTDeleteRole Delete an existing role.
// name: Name of the role.
func (a *Auth) JWTDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// JWTListRoles Lists all the roles registered with the backend.
// The list will contain the names of the roles.
func (a *Auth) JWTListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// JWTLogin Authenticates to Vault using a JWT (or OIDC) token.
func (a *Auth) JWTLogin(ctx context.Context, request schema.JWTLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// JWTReadConfig Read the current JWT authentication backend configuration.
func (a *Auth) JWTReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// JWTReadOIDCCallback Callback endpoint to complete an OIDC login.
func (a *Auth) JWTReadOIDCCallback(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/oidc/callback"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// JWTReadRole Read an existing role.
// name: Name of the role.
func (a *Auth) JWTReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// JWTWriteConfig Configure the JWT authentication backend.
// The JWT authentication backend validates JWTs (or OIDC) using the configured credentials. If using OIDC Discovery, the URL must be provided, along with (optionally) the CA cert to use for the connection. If performing JWT validation locally, a set of public keys must be provided.
func (a *Auth) JWTWriteConfig(ctx context.Context, request schema.JWTWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// JWTWriteOIDCAuthURL Request an authorization URL to start an OIDC login flow.
func (a *Auth) JWTWriteOIDCAuthURL(ctx context.Context, request schema.JWTWriteOIDCAuthURLRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/oidc/auth_url"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// JWTWriteOIDCCallback Callback endpoint to handle form_posts.
func (a *Auth) JWTWriteOIDCCallback(ctx context.Context, request schema.JWTWriteOIDCCallbackRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/oidc/callback"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// JWTWriteRole Register an role with the backend.
// A role is required to authenticate with this backend. The role binds   JWT token information with token policies and settings.   The bindings, token polices and token settings can all be configured   using this endpoint
// name: Name of the role.
func (a *Auth) JWTWriteRole(ctx context.Context, name string, request schema.JWTWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosDeleteGroup
// name: Name of the LDAP group.
func (a *Auth) KerberosDeleteGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosListGroups
func (a *Auth) KerberosListGroups(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosLogin
func (a *Auth) KerberosLogin(ctx context.Context, request schema.KerberosLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosReadConfig
func (a *Auth) KerberosReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosReadGroup
// name: Name of the LDAP group.
func (a *Auth) KerberosReadGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosReadLDAPConfig
func (a *Auth) KerberosReadLDAPConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config/ldap"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosWriteConfig
func (a *Auth) KerberosWriteConfig(ctx context.Context, request schema.KerberosWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosWriteGroup
// name: Name of the LDAP group.
func (a *Auth) KerberosWriteGroup(ctx context.Context, name string, request schema.KerberosWriteGroupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosWriteLDAPConfig
func (a *Auth) KerberosWriteLDAPConfig(ctx context.Context, request schema.KerberosWriteLDAPConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config/ldap"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesDeleteAuthRole Register an role with the backend.
// name: Name of the role.
func (a *Auth) KubernetesDeleteAuthRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesListAuthRoles Lists all the roles registered with the backend.
func (a *Auth) KubernetesListAuthRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesLogin Authenticates Kubernetes service accounts with Vault.
func (a *Auth) KubernetesLogin(ctx context.Context, request schema.KubernetesLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesReadAuthConfig Configures the JWT Public Key and Kubernetes API information.
func (a *Auth) KubernetesReadAuthConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesReadAuthRole Register an role with the backend.
// name: Name of the role.
func (a *Auth) KubernetesReadAuthRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesWriteAuthConfig Configures the JWT Public Key and Kubernetes API information.
func (a *Auth) KubernetesWriteAuthConfig(ctx context.Context, request schema.KubernetesWriteAuthConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesWriteAuthRole Register an role with the backend.
// name: Name of the role.
func (a *Auth) KubernetesWriteAuthRole(ctx context.Context, name string, request schema.KubernetesWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPDeleteGroup Manage additional groups for users allowed to authenticate.
// name: Name of the LDAP group.
func (a *Auth) LDAPDeleteGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPDeleteUser Manage users allowed to authenticate.
// name: Name of the LDAP user.
func (a *Auth) LDAPDeleteUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPListGroups Manage additional groups for users allowed to authenticate.
func (a *Auth) LDAPListGroups(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPListUsers Manage users allowed to authenticate.
func (a *Auth) LDAPListUsers(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPLogin Log in with a username and password.
// username: DN (distinguished name) to be used for login.
func (a *Auth) LDAPLogin(ctx context.Context, username string, request schema.LDAPLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/login/{username}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPReadAuthConfig Configure the LDAP server to connect to, along with its options.
func (a *Auth) LDAPReadAuthConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPReadGroup Manage additional groups for users allowed to authenticate.
// name: Name of the LDAP group.
func (a *Auth) LDAPReadGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPReadUser Manage users allowed to authenticate.
// name: Name of the LDAP user.
func (a *Auth) LDAPReadUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPWriteAuthConfig Configure the LDAP server to connect to, along with its options.
func (a *Auth) LDAPWriteAuthConfig(ctx context.Context, request schema.LDAPWriteAuthConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPWriteGroup Manage additional groups for users allowed to authenticate.
// name: Name of the LDAP group.
func (a *Auth) LDAPWriteGroup(ctx context.Context, name string, request schema.LDAPWriteGroupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPWriteUser Manage users allowed to authenticate.
// name: Name of the LDAP user.
func (a *Auth) LDAPWriteUser(ctx context.Context, name string, request schema.LDAPWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OCIDeleteConfig Manages the configuration for the Vault Auth Plugin.
func (a *Auth) OCIDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OCIDeleteRole Create a role and associate policies to it.
// role: Name of the role.
func (a *Auth) OCIDeleteRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OCIListRoles Lists all the roles that are registered with Vault.
func (a *Auth) OCIListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OCILoginWithRole Authenticates to Vault using OCI credentials
// role: Name of the role.
func (a *Auth) OCILoginWithRole(ctx context.Context, role string, request schema.OCILoginWithRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/login/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OCIReadConfig Manages the configuration for the Vault Auth Plugin.
func (a *Auth) OCIReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OCIReadRole Create a role and associate policies to it.
// role: Name of the role.
func (a *Auth) OCIReadRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OCIWriteConfig Manages the configuration for the Vault Auth Plugin.
func (a *Auth) OCIWriteConfig(ctx context.Context, request schema.OCIWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OCIWriteRole Create a role and associate policies to it.
// role: Name of the role.
func (a *Auth) OCIWriteRole(ctx context.Context, role string, request schema.OCIWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OIDCDeleteAuthRole Delete an existing role.
// name: Name of the role.
func (a *Auth) OIDCDeleteAuthRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oidc_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"oidc_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oidc")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OIDCListAuthRoles Lists all the roles registered with the backend.
// The list will contain the names of the roles.
func (a *Auth) OIDCListAuthRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oidc_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"oidc_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oidc")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OIDCLogin Authenticates to Vault using a JWT (or OIDC) token.
func (a *Auth) OIDCLogin(ctx context.Context, request schema.OIDCLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oidc_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"oidc_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oidc")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OIDCReadAuthConfig Read the current JWT authentication backend configuration.
func (a *Auth) OIDCReadAuthConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oidc_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"oidc_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oidc")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OIDCReadAuthRole Read an existing role.
// name: Name of the role.
func (a *Auth) OIDCReadAuthRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oidc_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"oidc_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oidc")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OIDCReadCallback Callback endpoint to complete an OIDC login.
func (a *Auth) OIDCReadCallback(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oidc_mount_path}/oidc/callback"
	requestPath = strings.Replace(requestPath, "{"+"oidc_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oidc")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OIDCWriteAuthConfig Configure the JWT authentication backend.
// The JWT authentication backend validates JWTs (or OIDC) using the configured credentials. If using OIDC Discovery, the URL must be provided, along with (optionally) the CA cert to use for the connection. If performing JWT validation locally, a set of public keys must be provided.
func (a *Auth) OIDCWriteAuthConfig(ctx context.Context, request schema.OIDCWriteAuthConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oidc_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"oidc_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oidc")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OIDCWriteAuthRole Register an role with the backend.
// A role is required to authenticate with this backend. The role binds   JWT token information with token policies and settings.   The bindings, token polices and token settings can all be configured   using this endpoint
// name: Name of the role.
func (a *Auth) OIDCWriteAuthRole(ctx context.Context, name string, request schema.OIDCWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oidc_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"oidc_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oidc")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OIDCWriteAuthURL Request an authorization URL to start an OIDC login flow.
func (a *Auth) OIDCWriteAuthURL(ctx context.Context, request schema.OIDCWriteAuthURLRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oidc_mount_path}/oidc/auth_url"
	requestPath = strings.Replace(requestPath, "{"+"oidc_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oidc")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OIDCWriteCallback Callback endpoint to handle form_posts.
func (a *Auth) OIDCWriteCallback(ctx context.Context, request schema.OIDCWriteCallbackRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oidc_mount_path}/oidc/callback"
	requestPath = strings.Replace(requestPath, "{"+"oidc_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oidc")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaDeleteGroup Manage users allowed to authenticate.
// name: Name of the Okta group.
func (a *Auth) OktaDeleteGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaDeleteUser Manage additional groups for users allowed to authenticate.
// name: Name of the user.
func (a *Auth) OktaDeleteUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaListGroups Manage users allowed to authenticate.
func (a *Auth) OktaListGroups(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaListUsers Manage additional groups for users allowed to authenticate.
func (a *Auth) OktaListUsers(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaLogin Log in with a username and password.
// username: Username to be used for login.
func (a *Auth) OktaLogin(ctx context.Context, username string, request schema.OktaLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/login/{username}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaReadConfig This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
func (a *Auth) OktaReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaReadGroup Manage users allowed to authenticate.
// name: Name of the Okta group.
func (a *Auth) OktaReadGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaReadUser Manage additional groups for users allowed to authenticate.
// name: Name of the user.
func (a *Auth) OktaReadUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaVerify
// nonce: Nonce provided during a login request to retrieve the number verification challenge for the matching request.
func (a *Auth) OktaVerify(ctx context.Context, nonce string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/verify/{nonce}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"nonce"+"}", url.PathEscape(nonce), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaWriteConfig This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
func (a *Auth) OktaWriteConfig(ctx context.Context, request schema.OktaWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaWriteGroup Manage users allowed to authenticate.
// name: Name of the Okta group.
func (a *Auth) OktaWriteGroup(ctx context.Context, name string, request schema.OktaWriteGroupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaWriteUser Manage additional groups for users allowed to authenticate.
// name: Name of the user.
func (a *Auth) OktaWriteUser(ctx context.Context, name string, request schema.OktaWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusDeleteUser Manage users allowed to authenticate.
// name: Name of the RADIUS user.
func (a *Auth) RadiusDeleteUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusListUsers Manage users allowed to authenticate.
func (a *Auth) RadiusListUsers(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusLogin Log in with a username and password.
func (a *Auth) RadiusLogin(ctx context.Context, request schema.RadiusLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusLoginWithUsername Log in with a username and password.
// urlusername: Username to be used for login. (URL parameter)
func (a *Auth) RadiusLoginWithUsername(ctx context.Context, urlusername string, request schema.RadiusLoginWithUsernameRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/login/{urlusername}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlusername"+"}", url.PathEscape(urlusername), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusReadConfig Configure the RADIUS server to connect to, along with its options.
func (a *Auth) RadiusReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusReadUser Manage users allowed to authenticate.
// name: Name of the RADIUS user.
func (a *Auth) RadiusReadUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusWriteConfig Configure the RADIUS server to connect to, along with its options.
func (a *Auth) RadiusWriteConfig(ctx context.Context, request schema.RadiusWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusWriteUser Manage users allowed to authenticate.
// name: Name of the RADIUS user.
func (a *Auth) RadiusWriteUser(ctx context.Context, name string, request schema.RadiusWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenDeleteRole
// roleName: Name of the role
func (a *Auth) TokenDeleteRole(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/roles/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenListAccessors List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires 'sudo' capability in addition to 'list'.
func (a *Auth) TokenListAccessors(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/accessors/"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenListRoles This endpoint lists configured roles.
func (a *Auth) TokenListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenReadLookup This endpoint will lookup a token and its properties.
func (a *Auth) TokenReadLookup(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/lookup"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenReadLookupSelf This endpoint will lookup a token and its properties.
func (a *Auth) TokenReadLookupSelf(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/lookup-self"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenReadRole
// roleName: Name of the role
func (a *Auth) TokenReadRole(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/roles/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRenew This endpoint will renew the given token and prevent expiration.
func (a *Auth) TokenRenew(ctx context.Context, request schema.TokenRenewRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/renew"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRenewAccessor This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.
func (a *Auth) TokenRenewAccessor(ctx context.Context, request schema.TokenRenewAccessorRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/renew-accessor"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRenewSelf This endpoint will renew the token used to call it and prevent expiration.
func (a *Auth) TokenRenewSelf(ctx context.Context, request schema.TokenRenewSelfRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/renew-self"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRevoke This endpoint will delete the given token and all of its child tokens.
func (a *Auth) TokenRevoke(ctx context.Context, request schema.TokenRevokeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/revoke"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRevokeAccessor This endpoint will delete the token associated with the accessor and all of its child tokens.
func (a *Auth) TokenRevokeAccessor(ctx context.Context, request schema.TokenRevokeAccessorRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/revoke-accessor"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRevokeOrphan This endpoint will delete the token and orphan its child tokens.
func (a *Auth) TokenRevokeOrphan(ctx context.Context, request schema.TokenRevokeOrphanRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/revoke-orphan"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRevokeSelf This endpoint will delete the token used to call it and all of its child tokens.
func (a *Auth) TokenRevokeSelf(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/revoke-self"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenTidy This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
func (a *Auth) TokenTidy(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/tidy"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenWriteCreate The token create path is used to create new tokens.
// format: Return json formatted output
func (a *Auth) TokenWriteCreate(ctx context.Context, request schema.TokenWriteCreateRequest, format string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/create"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("format", url.QueryEscape(format))

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenWriteCreateOrphan The token create path is used to create new orphan tokens.
// format: Return json formatted output
func (a *Auth) TokenWriteCreateOrphan(ctx context.Context, request schema.TokenWriteCreateOrphanRequest, format string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/create-orphan"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("format", url.QueryEscape(format))

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenWriteCreateWithRole This token create path is used to create new tokens adhering to the given role.
// roleName: Name of the role
// format: Return json formatted output
func (a *Auth) TokenWriteCreateWithRole(ctx context.Context, roleName string, request schema.TokenWriteCreateWithRoleRequest, format string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/create/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("format", url.QueryEscape(format))

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenWriteLookup This endpoint will lookup a token and its properties.
func (a *Auth) TokenWriteLookup(ctx context.Context, request schema.TokenWriteLookupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/lookup"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenWriteLookupAccessor This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.
func (a *Auth) TokenWriteLookupAccessor(ctx context.Context, request schema.TokenWriteLookupAccessorRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/lookup-accessor"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenWriteLookupSelf This endpoint will lookup a token and its properties.
func (a *Auth) TokenWriteLookupSelf(ctx context.Context, request schema.TokenWriteLookupSelfRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/lookup-self"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenWriteRole
// roleName: Name of the role
func (a *Auth) TokenWriteRole(ctx context.Context, roleName string, request schema.TokenWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{token_mount_path}/roles/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"token_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("token")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassDeleteUser Manage users allowed to authenticate.
// username: Username for this user.
func (a *Auth) UserpassDeleteUser(ctx context.Context, username string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassListUsers Manage users allowed to authenticate.
func (a *Auth) UserpassListUsers(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassLogin Log in with a username and password.
// username: Username of the user.
func (a *Auth) UserpassLogin(ctx context.Context, username string, request schema.UserpassLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/login/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassReadUser Manage users allowed to authenticate.
// username: Username for this user.
func (a *Auth) UserpassReadUser(ctx context.Context, username string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassWriteUser Manage users allowed to authenticate.
// username: Username for this user.
func (a *Auth) UserpassWriteUser(ctx context.Context, username string, request schema.UserpassWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassWriteUserPassword Reset user's password.
// username: Username for this user.
func (a *Auth) UserpassWriteUserPassword(ctx context.Context, username string, request schema.UserpassWriteUserPasswordRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}/password"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassWriteUserPolicies Update the policies associated with the username.
// username: Username for this user.
func (a *Auth) UserpassWriteUserPolicies(ctx context.Context, username string, request schema.UserpassWriteUserPoliciesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}/policies"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}
