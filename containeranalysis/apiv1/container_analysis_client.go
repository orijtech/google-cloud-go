// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gapic-generator. DO NOT EDIT.

package containeranalysis

import (
	"context"
	"fmt"
	"net/url"

	grafeas "cloud.google.com/go/grafeas/apiv1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	containeranalysispb "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	SetIamPolicy       []gax.CallOption
	GetIamPolicy       []gax.CallOption
	TestIamPermissions []gax.CallOption
}

func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("containeranalysis.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultCallOptions() *CallOptions {
	retry := map[[2]string][]gax.CallOption{}
	return &CallOptions{
		SetIamPolicy:       retry[[2]string{"default", "non_idempotent"}],
		GetIamPolicy:       retry[[2]string{"default", "non_idempotent"}],
		TestIamPermissions: retry[[2]string{"default", "non_idempotent"}],
	}
}

// Client is a client for interacting with Container Analysis API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type Client struct {
	// The connection to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	client containeranalysispb.ContainerAnalysisClient

	// A pre-crearted grafeas client.
	grafeasClient *grafeas.Client

	// The call options for this service.
	CallOptions *CallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new container analysis client.
//
// Retrieves analysis results of Cloud components such as Docker container
// images. The Container Analysis API is an implementation of the
// Grafeas (at https://grafeas.io) API.
//
// Analysis results are stored as a series of occurrences. An Occurrence
// contains information about a specific analysis instance on a resource. An
// occurrence refers to a Note. A note contains details describing the
// analysis and is generally stored in a separate project, called a Provider.
// Multiple occurrences can refer to the same note.
//
// For example, an SSL vulnerability could affect multiple images. In this case,
// there would be one note for the vulnerability and an occurrence for each
// image with the vulnerability referring to that note.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	connPool, err := gtransport.DialPool(ctx, append(defaultClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	gc, err := grafeas.NewClient(ctx, append(defaultClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		connPool:    connPool,
		CallOptions: defaultCallOptions(),

		client:        containeranalysispb.NewContainerAnalysisClient(connPool),
		grafeasClient: gc,
	}
	c.setGoogleClientInfo()
	return c, nil
}

// GetGrafeasClient returns a grafeas client connected to containeranalysis.
func (c *Client) GetGrafeasClient() *grafeas.Client {
	return c.grafeasClient
}

// Connection returns the client's connection to the API service.
func (c *Client) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
//
// Calling Close will also close the underlying grafeas client.
func (c *Client) Close() error {
	c.grafeasClient.Close()
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// SetIamPolicy sets the access control policy on the specified note or occurrence.
// Requires containeranalysis.notes.setIamPolicy or
// containeranalysis.occurrences.setIamPolicy permission if the resource is
// a note or an occurrence, respectively.
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *Client) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SetIamPolicy[0:len(c.CallOptions.SetIamPolicy):len(c.CallOptions.SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetIamPolicy gets the access control policy for a note or an occurrence resource.
// Requires containeranalysis.notes.setIamPolicy or
// containeranalysis.occurrences.setIamPolicy permission if the resource is
// a note or occurrence, respectively.
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *Client) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetIamPolicy[0:len(c.CallOptions.GetIamPolicy):len(c.CallOptions.GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TestIamPermissions returns the permissions that a caller has on the specified note or
// occurrence. Requires list permission on the project (for example,
// containeranalysis.notes.list).
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *Client) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.TestIamPermissions[0:len(c.CallOptions.TestIamPermissions):len(c.CallOptions.TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
