// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cloudidentity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Membership) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"DisplayName"}, r.DisplayName); err != nil {
		return err
	}
	if err := dcl.Required(r, "preferredMemberKey"); err != nil {
		return err
	}
	if err := dcl.Required(r, "roles"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Group, "Group"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.PreferredMemberKey) {
		if err := r.PreferredMemberKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DisplayName) {
		if err := r.DisplayName.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipPreferredMemberKey) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	return nil
}
func (r *MembershipRoles) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ExpiryDetail) {
		if err := r.ExpiryDetail.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RestrictionEvaluations) {
		if err := r.RestrictionEvaluations.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipRolesExpiryDetail) validate() error {
	return nil
}
func (r *MembershipRolesRestrictionEvaluations) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MemberRestrictionEvaluation) {
		if err := r.MemberRestrictionEvaluation.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) validate() error {
	return nil
}
func (r *MembershipDisplayName) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"GivenName"}, r.GivenName); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"FamilyName"}, r.FamilyName); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"FullName"}, r.FullName); err != nil {
		return err
	}
	return nil
}
func (r *Membership) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://cloudidentity.googleapis.com/v1/", params)
}

func (r *Membership) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"group": dcl.ValueOrEmptyString(nr.Group),
		"name":  dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("groups/{{group}}/memberships/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Membership) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"group": dcl.ValueOrEmptyString(nr.Group),
	}
	return dcl.URL("groups/{{group}}/memberships", nr.basePath(), userBasePath, params), nil

}

func (r *Membership) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"group": dcl.ValueOrEmptyString(nr.Group),
	}
	return dcl.URL("groups/{{group}}/memberships/", nr.basePath(), userBasePath, params), nil

}

func (r *Membership) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"group": dcl.ValueOrEmptyString(nr.Group),
		"name":  dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("groups/{{group}}/memberships/{{name}}", nr.basePath(), userBasePath, params), nil
}

// membershipApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type membershipApiOperation interface {
	do(context.Context, *Membership, *Client) error
}

// newUpdateMembershipUpdateMembershipRequest creates a request for an
// Membership resource's UpdateMembership update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateMembershipUpdateMembershipRequest(ctx context.Context, f *Membership, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("groups/%s/memberships/%s", f.Name, dcl.SelfLinkToName(f.Group), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v, err := expandMembershipRolesSlice(c, f.Roles); err != nil {
		return nil, fmt.Errorf("error expanding Roles into roles: %w", err)
	} else if v != nil {
		req["roles"] = v
	}
	return req, nil
}

// marshalUpdateMembershipUpdateMembershipRequest converts the update into
// the final JSON request body.
func marshalUpdateMembershipUpdateMembershipRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateMembershipUpdateMembershipOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listMembershipRaw(ctx context.Context, r *Membership, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != MembershipMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listMembershipOperation struct {
	Memberships []map[string]interface{} `json:"memberships"`
	Token       string                   `json:"nextPageToken"`
}

func (c *Client) listMembership(ctx context.Context, r *Membership, pageToken string, pageSize int32) ([]*Membership, string, error) {
	b, err := c.listMembershipRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listMembershipOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Membership
	for _, v := range m.Memberships {
		res, err := unmarshalMapMembership(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Group = r.Group
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllMembership(ctx context.Context, f func(*Membership) bool, resources []*Membership) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteMembership(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteMembershipOperation struct{}

func (op *deleteMembershipOperation) do(ctx context.Context, r *Membership, c *Client) error {
	r, err := c.GetMembership(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Membership not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetMembership checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetMembership(ctx, r)
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createMembershipOperation struct {
	response map[string]interface{}
}

func (op *createMembershipOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createMembershipOperation) do(ctx context.Context, r *Membership, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string in %v, was %T", op.response, op.response["name"])
	}
	r.Name = &name

	if _, err := c.GetMembership(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getMembershipRaw(ctx context.Context, r *Membership) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) membershipDiffsForRawDesired(ctx context.Context, rawDesired *Membership, opts ...dcl.ApplyOption) (initial, desired *Membership, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Membership
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Membership); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Membership, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeMembershipDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetMembership(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Membership resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Membership resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Membership resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeMembershipDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Membership: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Membership: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeMembershipInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Membership: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeMembershipDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Membership: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffMembership(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeMembershipInitialState(rawInitial, rawDesired *Membership) (*Membership, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.DisplayName) {
		// Check if anything else is set.
		if dcl.AnySet() {
			rawInitial.DisplayName = EmptyMembershipDisplayName
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeMembershipDesiredState(rawDesired, rawInitial *Membership, opts ...dcl.ApplyOption) (*Membership, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.PreferredMemberKey = canonicalizeMembershipPreferredMemberKey(rawDesired.PreferredMemberKey, nil, opts...)
		rawDesired.DisplayName = canonicalizeMembershipDisplayName(rawDesired.DisplayName, nil, opts...)

		return rawDesired, nil
	}

	if rawDesired.DisplayName != nil || rawInitial.DisplayName != nil {
		// Check if anything else is set.
		if dcl.AnySet() {
			rawDesired.DisplayName = nil
			rawInitial.DisplayName = nil
		}
	}

	canonicalDesired := &Membership{}
	if dcl.IsZeroValue(rawDesired.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.PreferredMemberKey = canonicalizeMembershipPreferredMemberKey(rawDesired.PreferredMemberKey, rawInitial.PreferredMemberKey, opts...)
	canonicalDesired.Roles = canonicalizeMembershipRolesSlice(rawDesired.Roles, rawInitial.Roles, opts...)
	if dcl.NameToSelfLink(rawDesired.Group, rawInitial.Group) {
		canonicalDesired.Group = rawInitial.Group
	} else {
		canonicalDesired.Group = rawDesired.Group
	}

	return canonicalDesired, nil
}

func canonicalizeMembershipNewState(c *Client, rawNew, rawDesired *Membership) (*Membership, error) {

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.PreferredMemberKey) && dcl.IsNotReturnedByServer(rawDesired.PreferredMemberKey) {
		rawNew.PreferredMemberKey = rawDesired.PreferredMemberKey
	} else {
		rawNew.PreferredMemberKey = canonicalizeNewMembershipPreferredMemberKey(c, rawDesired.PreferredMemberKey, rawNew.PreferredMemberKey)
	}

	if dcl.IsNotReturnedByServer(rawNew.CreateTime) && dcl.IsNotReturnedByServer(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.UpdateTime) && dcl.IsNotReturnedByServer(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Roles) && dcl.IsNotReturnedByServer(rawDesired.Roles) {
		rawNew.Roles = rawDesired.Roles
	} else {
		rawNew.Roles = canonicalizeNewMembershipRolesSet(c, rawDesired.Roles, rawNew.Roles)
	}

	if dcl.IsNotReturnedByServer(rawNew.Type) && dcl.IsNotReturnedByServer(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.DeliverySetting) && dcl.IsNotReturnedByServer(rawDesired.DeliverySetting) {
		rawNew.DeliverySetting = rawDesired.DeliverySetting
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.DisplayName) && dcl.IsNotReturnedByServer(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		rawNew.DisplayName = canonicalizeNewMembershipDisplayName(c, rawDesired.DisplayName, rawNew.DisplayName)
	}

	rawNew.Group = rawDesired.Group

	return rawNew, nil
}

func canonicalizeMembershipPreferredMemberKey(des, initial *MembershipPreferredMemberKey, opts ...dcl.ApplyOption) *MembershipPreferredMemberKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipPreferredMemberKey{}

	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		cDes.Id = initial.Id
	} else {
		cDes.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Namespace, initial.Namespace) || dcl.IsZeroValue(des.Namespace) {
		cDes.Namespace = initial.Namespace
	} else {
		cDes.Namespace = des.Namespace
	}

	return cDes
}

func canonicalizeMembershipPreferredMemberKeySlice(des, initial []MembershipPreferredMemberKey, opts ...dcl.ApplyOption) []MembershipPreferredMemberKey {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipPreferredMemberKey, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipPreferredMemberKey(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipPreferredMemberKey, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipPreferredMemberKey(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipPreferredMemberKey(c *Client, des, nw *MembershipPreferredMemberKey) *MembershipPreferredMemberKey {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipPreferredMemberKey while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Namespace, nw.Namespace) {
		nw.Namespace = des.Namespace
	}

	return nw
}

func canonicalizeNewMembershipPreferredMemberKeySet(c *Client, des, nw []MembershipPreferredMemberKey) []MembershipPreferredMemberKey {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipPreferredMemberKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipPreferredMemberKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewMembershipPreferredMemberKeySlice(c *Client, des, nw []MembershipPreferredMemberKey) []MembershipPreferredMemberKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipPreferredMemberKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipPreferredMemberKey(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipRoles(des, initial *MembershipRoles, opts ...dcl.ApplyOption) *MembershipRoles {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipRoles{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	cDes.ExpiryDetail = canonicalizeMembershipRolesExpiryDetail(des.ExpiryDetail, initial.ExpiryDetail, opts...)
	cDes.RestrictionEvaluations = canonicalizeMembershipRolesRestrictionEvaluations(des.RestrictionEvaluations, initial.RestrictionEvaluations, opts...)

	return cDes
}

func canonicalizeMembershipRolesSlice(des, initial []MembershipRoles, opts ...dcl.ApplyOption) []MembershipRoles {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipRoles, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipRoles(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipRoles, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipRoles(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipRoles(c *Client, des, nw *MembershipRoles) *MembershipRoles {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipRoles while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.ExpiryDetail = canonicalizeNewMembershipRolesExpiryDetail(c, des.ExpiryDetail, nw.ExpiryDetail)
	nw.RestrictionEvaluations = canonicalizeNewMembershipRolesRestrictionEvaluations(c, des.RestrictionEvaluations, nw.RestrictionEvaluations)

	return nw
}

func canonicalizeNewMembershipRolesSet(c *Client, des, nw []MembershipRoles) []MembershipRoles {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipRoles
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipRolesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewMembershipRolesSlice(c *Client, des, nw []MembershipRoles) []MembershipRoles {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipRoles
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipRoles(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipRolesExpiryDetail(des, initial *MembershipRolesExpiryDetail, opts ...dcl.ApplyOption) *MembershipRolesExpiryDetail {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipRolesExpiryDetail{}

	if dcl.IsZeroValue(des.ExpireTime) {
		des.ExpireTime = initial.ExpireTime
	} else {
		cDes.ExpireTime = des.ExpireTime
	}

	return cDes
}

func canonicalizeMembershipRolesExpiryDetailSlice(des, initial []MembershipRolesExpiryDetail, opts ...dcl.ApplyOption) []MembershipRolesExpiryDetail {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipRolesExpiryDetail, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipRolesExpiryDetail(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipRolesExpiryDetail, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipRolesExpiryDetail(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipRolesExpiryDetail(c *Client, des, nw *MembershipRolesExpiryDetail) *MembershipRolesExpiryDetail {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipRolesExpiryDetail while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewMembershipRolesExpiryDetailSet(c *Client, des, nw []MembershipRolesExpiryDetail) []MembershipRolesExpiryDetail {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipRolesExpiryDetail
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipRolesExpiryDetailNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewMembershipRolesExpiryDetailSlice(c *Client, des, nw []MembershipRolesExpiryDetail) []MembershipRolesExpiryDetail {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipRolesExpiryDetail
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipRolesExpiryDetail(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipRolesRestrictionEvaluations(des, initial *MembershipRolesRestrictionEvaluations, opts ...dcl.ApplyOption) *MembershipRolesRestrictionEvaluations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipRolesRestrictionEvaluations{}

	cDes.MemberRestrictionEvaluation = canonicalizeMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(des.MemberRestrictionEvaluation, initial.MemberRestrictionEvaluation, opts...)

	return cDes
}

func canonicalizeMembershipRolesRestrictionEvaluationsSlice(des, initial []MembershipRolesRestrictionEvaluations, opts ...dcl.ApplyOption) []MembershipRolesRestrictionEvaluations {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipRolesRestrictionEvaluations, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipRolesRestrictionEvaluations(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipRolesRestrictionEvaluations, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipRolesRestrictionEvaluations(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipRolesRestrictionEvaluations(c *Client, des, nw *MembershipRolesRestrictionEvaluations) *MembershipRolesRestrictionEvaluations {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipRolesRestrictionEvaluations while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.MemberRestrictionEvaluation = canonicalizeNewMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, des.MemberRestrictionEvaluation, nw.MemberRestrictionEvaluation)

	return nw
}

func canonicalizeNewMembershipRolesRestrictionEvaluationsSet(c *Client, des, nw []MembershipRolesRestrictionEvaluations) []MembershipRolesRestrictionEvaluations {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipRolesRestrictionEvaluations
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipRolesRestrictionEvaluationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewMembershipRolesRestrictionEvaluationsSlice(c *Client, des, nw []MembershipRolesRestrictionEvaluations) []MembershipRolesRestrictionEvaluations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipRolesRestrictionEvaluations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipRolesRestrictionEvaluations(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(des, initial *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, opts ...dcl.ApplyOption) *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}

	return cDes
}

func canonicalizeMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice(des, initial []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, opts ...dcl.ApplyOption) []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c *Client, des, nw *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSet(c *Client, des, nw []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice(c *Client, des, nw []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipDisplayName(des, initial *MembershipDisplayName, opts ...dcl.ApplyOption) *MembershipDisplayName {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.GivenName != nil || (initial != nil && initial.GivenName != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.GivenName = nil
			if initial != nil {
				initial.GivenName = nil
			}
		}
	}

	if des.FamilyName != nil || (initial != nil && initial.FamilyName != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.FamilyName = nil
			if initial != nil {
				initial.FamilyName = nil
			}
		}
	}

	if des.FullName != nil || (initial != nil && initial.FullName != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.FullName = nil
			if initial != nil {
				initial.FullName = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipDisplayName{}

	return cDes
}

func canonicalizeMembershipDisplayNameSlice(des, initial []MembershipDisplayName, opts ...dcl.ApplyOption) []MembershipDisplayName {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipDisplayName, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipDisplayName(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipDisplayName, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipDisplayName(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipDisplayName(c *Client, des, nw *MembershipDisplayName) *MembershipDisplayName {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipDisplayName while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.GivenName, nw.GivenName) {
		nw.GivenName = des.GivenName
	}
	if dcl.StringCanonicalize(des.FamilyName, nw.FamilyName) {
		nw.FamilyName = des.FamilyName
	}
	if dcl.StringCanonicalize(des.FullName, nw.FullName) {
		nw.FullName = des.FullName
	}

	return nw
}

func canonicalizeNewMembershipDisplayNameSet(c *Client, des, nw []MembershipDisplayName) []MembershipDisplayName {
	if des == nil {
		return nw
	}
	var reorderedNew []MembershipDisplayName
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMembershipDisplayNameNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewMembershipDisplayNameSlice(c *Client, des, nw []MembershipDisplayName) []MembershipDisplayName {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipDisplayName
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipDisplayName(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffMembership(c *Client, desired, actual *Membership, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PreferredMemberKey, actual.PreferredMemberKey, dcl.Info{ObjectFunction: compareMembershipPreferredMemberKeyNewStyle, EmptyObject: EmptyMembershipPreferredMemberKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PreferredMemberKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Roles, actual.Roles, dcl.Info{Type: "Set", ObjectFunction: compareMembershipRolesNewStyle, EmptyObject: EmptyMembershipRoles, OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Roles")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeliverySetting, actual.DeliverySetting, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeliverySetting")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OutputOnly: true, ObjectFunction: compareMembershipDisplayNameNewStyle, EmptyObject: EmptyMembershipDisplayName, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Group, actual.Group, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Group")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareMembershipPreferredMemberKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipPreferredMemberKey)
	if !ok {
		desiredNotPointer, ok := d.(MembershipPreferredMemberKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipPreferredMemberKey or *MembershipPreferredMemberKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipPreferredMemberKey)
	if !ok {
		actualNotPointer, ok := a.(MembershipPreferredMemberKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipPreferredMemberKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipRolesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipRoles)
	if !ok {
		desiredNotPointer, ok := d.(MembershipRoles)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRoles or *MembershipRoles", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipRoles)
	if !ok {
		actualNotPointer, ok := a.(MembershipRoles)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRoles", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpiryDetail, actual.ExpiryDetail, dcl.Info{ObjectFunction: compareMembershipRolesExpiryDetailNewStyle, EmptyObject: EmptyMembershipRolesExpiryDetail, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExpiryDetail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RestrictionEvaluations, actual.RestrictionEvaluations, dcl.Info{ObjectFunction: compareMembershipRolesRestrictionEvaluationsNewStyle, EmptyObject: EmptyMembershipRolesRestrictionEvaluations, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RestrictionEvaluations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipRolesExpiryDetailNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipRolesExpiryDetail)
	if !ok {
		desiredNotPointer, ok := d.(MembershipRolesExpiryDetail)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesExpiryDetail or *MembershipRolesExpiryDetail", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipRolesExpiryDetail)
	if !ok {
		actualNotPointer, ok := a.(MembershipRolesExpiryDetail)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesExpiryDetail", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ExpireTime, actual.ExpireTime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExpireTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipRolesRestrictionEvaluationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipRolesRestrictionEvaluations)
	if !ok {
		desiredNotPointer, ok := d.(MembershipRolesRestrictionEvaluations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesRestrictionEvaluations or *MembershipRolesRestrictionEvaluations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipRolesRestrictionEvaluations)
	if !ok {
		actualNotPointer, ok := a.(MembershipRolesRestrictionEvaluations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesRestrictionEvaluations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MemberRestrictionEvaluation, actual.MemberRestrictionEvaluation, dcl.Info{ObjectFunction: compareMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationNewStyle, EmptyObject: EmptyMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MemberRestrictionEvaluation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation)
	if !ok {
		desiredNotPointer, ok := d.(MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation or *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation)
	if !ok {
		actualNotPointer, ok := a.(MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipDisplayNameNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipDisplayName)
	if !ok {
		desiredNotPointer, ok := d.(MembershipDisplayName)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipDisplayName or *MembershipDisplayName", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipDisplayName)
	if !ok {
		actualNotPointer, ok := a.(MembershipDisplayName)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipDisplayName", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GivenName, actual.GivenName, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GivenName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FamilyName, actual.FamilyName, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FamilyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FullName, actual.FullName, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FullName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Membership) urlNormalized() *Membership {
	normalized := dcl.Copy(*r).(Membership)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Group = dcl.SelfLinkToName(r.Group)
	return &normalized
}

func (r *Membership) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateMembership" {
		fields := map[string]interface{}{
			"group": dcl.ValueOrEmptyString(nr.Group),
			"name":  dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("groups/{{group}}/memberships/{{name}}:modifyMembershipRoles", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Membership resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Membership) marshal(c *Client) ([]byte, error) {
	m, err := expandMembership(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Membership: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalMembership decodes JSON responses into the Membership resource schema.
func unmarshalMembership(b []byte, c *Client) (*Membership, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapMembership(m, c)
}

func unmarshalMapMembership(m map[string]interface{}, c *Client) (*Membership, error) {

	flattened := flattenMembership(c, m)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandMembership expands Membership into a JSON request object.
func expandMembership(c *Client, f *Membership) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("groups/%s/memberships/%s", f.Name, dcl.SelfLinkToName(f.Group), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v, err := expandMembershipPreferredMemberKey(c, f.PreferredMemberKey); err != nil {
		return nil, fmt.Errorf("error expanding PreferredMemberKey into preferredMemberKey: %w", err)
	} else if v != nil {
		m["preferredMemberKey"] = v
	}
	if v, err := expandMembershipRolesSlice(c, f.Roles); err != nil {
		return nil, fmt.Errorf("error expanding Roles into roles: %w", err)
	} else {
		m["roles"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Group into group: %w", err)
	} else if v != nil {
		m["group"] = v
	}

	return m, nil
}

// flattenMembership flattens Membership from a JSON request object into the
// Membership type.
func flattenMembership(c *Client, i interface{}) *Membership {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Membership{}
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	res.PreferredMemberKey = flattenMembershipPreferredMemberKey(c, m["preferredMemberKey"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Roles = flattenMembershipRolesSlice(c, m["roles"])
	res.Type = flattenMembershipTypeEnum(m["type"])
	res.DeliverySetting = flattenMembershipDeliverySettingEnum(m["deliverySetting"])
	res.DisplayName = flattenMembershipDisplayName(c, m["displayName"])
	res.Group = dcl.FlattenString(m["group"])

	return res
}

// expandMembershipPreferredMemberKeyMap expands the contents of MembershipPreferredMemberKey into a JSON
// request object.
func expandMembershipPreferredMemberKeyMap(c *Client, f map[string]MembershipPreferredMemberKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipPreferredMemberKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipPreferredMemberKeySlice expands the contents of MembershipPreferredMemberKey into a JSON
// request object.
func expandMembershipPreferredMemberKeySlice(c *Client, f []MembershipPreferredMemberKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipPreferredMemberKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipPreferredMemberKeyMap flattens the contents of MembershipPreferredMemberKey from a JSON
// response object.
func flattenMembershipPreferredMemberKeyMap(c *Client, i interface{}) map[string]MembershipPreferredMemberKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipPreferredMemberKey{}
	}

	if len(a) == 0 {
		return map[string]MembershipPreferredMemberKey{}
	}

	items := make(map[string]MembershipPreferredMemberKey)
	for k, item := range a {
		items[k] = *flattenMembershipPreferredMemberKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipPreferredMemberKeySlice flattens the contents of MembershipPreferredMemberKey from a JSON
// response object.
func flattenMembershipPreferredMemberKeySlice(c *Client, i interface{}) []MembershipPreferredMemberKey {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipPreferredMemberKey{}
	}

	if len(a) == 0 {
		return []MembershipPreferredMemberKey{}
	}

	items := make([]MembershipPreferredMemberKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipPreferredMemberKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipPreferredMemberKey expands an instance of MembershipPreferredMemberKey into a JSON
// request object.
func expandMembershipPreferredMemberKey(c *Client, f *MembershipPreferredMemberKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Namespace; !dcl.IsEmptyValueIndirect(v) {
		m["namespace"] = v
	}

	return m, nil
}

// flattenMembershipPreferredMemberKey flattens an instance of MembershipPreferredMemberKey from a JSON
// response object.
func flattenMembershipPreferredMemberKey(c *Client, i interface{}) *MembershipPreferredMemberKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipPreferredMemberKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipPreferredMemberKey
	}
	r.Id = dcl.FlattenString(m["id"])
	r.Namespace = dcl.FlattenString(m["namespace"])

	return r
}

// expandMembershipRolesMap expands the contents of MembershipRoles into a JSON
// request object.
func expandMembershipRolesMap(c *Client, f map[string]MembershipRoles) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipRoles(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipRolesSlice expands the contents of MembershipRoles into a JSON
// request object.
func expandMembershipRolesSlice(c *Client, f []MembershipRoles) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipRoles(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipRolesMap flattens the contents of MembershipRoles from a JSON
// response object.
func flattenMembershipRolesMap(c *Client, i interface{}) map[string]MembershipRoles {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipRoles{}
	}

	if len(a) == 0 {
		return map[string]MembershipRoles{}
	}

	items := make(map[string]MembershipRoles)
	for k, item := range a {
		items[k] = *flattenMembershipRoles(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipRolesSlice flattens the contents of MembershipRoles from a JSON
// response object.
func flattenMembershipRolesSlice(c *Client, i interface{}) []MembershipRoles {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipRoles{}
	}

	if len(a) == 0 {
		return []MembershipRoles{}
	}

	items := make([]MembershipRoles, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipRoles(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipRoles expands an instance of MembershipRoles into a JSON
// request object.
func expandMembershipRoles(c *Client, f *MembershipRoles) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandMembershipRolesExpiryDetail(c, f.ExpiryDetail); err != nil {
		return nil, fmt.Errorf("error expanding ExpiryDetail into expiryDetail: %w", err)
	} else if v != nil {
		m["expiryDetail"] = v
	}
	if v, err := expandMembershipRolesRestrictionEvaluations(c, f.RestrictionEvaluations); err != nil {
		return nil, fmt.Errorf("error expanding RestrictionEvaluations into restrictionEvaluations: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["restrictionEvaluations"] = v
	}

	return m, nil
}

// flattenMembershipRoles flattens an instance of MembershipRoles from a JSON
// response object.
func flattenMembershipRoles(c *Client, i interface{}) *MembershipRoles {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipRoles{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipRoles
	}
	r.Name = dcl.FlattenString(m["name"])
	r.ExpiryDetail = flattenMembershipRolesExpiryDetail(c, m["expiryDetail"])
	r.RestrictionEvaluations = flattenMembershipRolesRestrictionEvaluations(c, m["restrictionEvaluations"])

	return r
}

// expandMembershipRolesExpiryDetailMap expands the contents of MembershipRolesExpiryDetail into a JSON
// request object.
func expandMembershipRolesExpiryDetailMap(c *Client, f map[string]MembershipRolesExpiryDetail) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipRolesExpiryDetail(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipRolesExpiryDetailSlice expands the contents of MembershipRolesExpiryDetail into a JSON
// request object.
func expandMembershipRolesExpiryDetailSlice(c *Client, f []MembershipRolesExpiryDetail) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipRolesExpiryDetail(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipRolesExpiryDetailMap flattens the contents of MembershipRolesExpiryDetail from a JSON
// response object.
func flattenMembershipRolesExpiryDetailMap(c *Client, i interface{}) map[string]MembershipRolesExpiryDetail {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipRolesExpiryDetail{}
	}

	if len(a) == 0 {
		return map[string]MembershipRolesExpiryDetail{}
	}

	items := make(map[string]MembershipRolesExpiryDetail)
	for k, item := range a {
		items[k] = *flattenMembershipRolesExpiryDetail(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipRolesExpiryDetailSlice flattens the contents of MembershipRolesExpiryDetail from a JSON
// response object.
func flattenMembershipRolesExpiryDetailSlice(c *Client, i interface{}) []MembershipRolesExpiryDetail {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipRolesExpiryDetail{}
	}

	if len(a) == 0 {
		return []MembershipRolesExpiryDetail{}
	}

	items := make([]MembershipRolesExpiryDetail, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipRolesExpiryDetail(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipRolesExpiryDetail expands an instance of MembershipRolesExpiryDetail into a JSON
// request object.
func expandMembershipRolesExpiryDetail(c *Client, f *MembershipRolesExpiryDetail) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ExpireTime; !dcl.IsEmptyValueIndirect(v) {
		m["expireTime"] = v
	}

	return m, nil
}

// flattenMembershipRolesExpiryDetail flattens an instance of MembershipRolesExpiryDetail from a JSON
// response object.
func flattenMembershipRolesExpiryDetail(c *Client, i interface{}) *MembershipRolesExpiryDetail {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipRolesExpiryDetail{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipRolesExpiryDetail
	}
	r.ExpireTime = dcl.FlattenString(m["expireTime"])

	return r
}

// expandMembershipRolesRestrictionEvaluationsMap expands the contents of MembershipRolesRestrictionEvaluations into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluationsMap(c *Client, f map[string]MembershipRolesRestrictionEvaluations) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipRolesRestrictionEvaluations(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipRolesRestrictionEvaluationsSlice expands the contents of MembershipRolesRestrictionEvaluations into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluationsSlice(c *Client, f []MembershipRolesRestrictionEvaluations) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipRolesRestrictionEvaluations(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipRolesRestrictionEvaluationsMap flattens the contents of MembershipRolesRestrictionEvaluations from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMap(c *Client, i interface{}) map[string]MembershipRolesRestrictionEvaluations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipRolesRestrictionEvaluations{}
	}

	if len(a) == 0 {
		return map[string]MembershipRolesRestrictionEvaluations{}
	}

	items := make(map[string]MembershipRolesRestrictionEvaluations)
	for k, item := range a {
		items[k] = *flattenMembershipRolesRestrictionEvaluations(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipRolesRestrictionEvaluationsSlice flattens the contents of MembershipRolesRestrictionEvaluations from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsSlice(c *Client, i interface{}) []MembershipRolesRestrictionEvaluations {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipRolesRestrictionEvaluations{}
	}

	if len(a) == 0 {
		return []MembershipRolesRestrictionEvaluations{}
	}

	items := make([]MembershipRolesRestrictionEvaluations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipRolesRestrictionEvaluations(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipRolesRestrictionEvaluations expands an instance of MembershipRolesRestrictionEvaluations into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluations(c *Client, f *MembershipRolesRestrictionEvaluations) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, f.MemberRestrictionEvaluation); err != nil {
		return nil, fmt.Errorf("error expanding MemberRestrictionEvaluation into memberRestrictionEvaluation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["memberRestrictionEvaluation"] = v
	}

	return m, nil
}

// flattenMembershipRolesRestrictionEvaluations flattens an instance of MembershipRolesRestrictionEvaluations from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluations(c *Client, i interface{}) *MembershipRolesRestrictionEvaluations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipRolesRestrictionEvaluations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipRolesRestrictionEvaluations
	}
	r.MemberRestrictionEvaluation = flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, m["memberRestrictionEvaluation"])

	return r
}

// expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationMap expands the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationMap(c *Client, f map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice expands the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice(c *Client, f []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationMap flattens the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationMap(c *Client, i interface{}) map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}

	if len(a) == 0 {
		return map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}

	items := make(map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation)
	for k, item := range a {
		items[k] = *flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice flattens the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice(c *Client, i interface{}) []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}

	if len(a) == 0 {
		return []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}

	items := make([]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation expands an instance of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c *Client, f *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation flattens an instance of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c *Client, i interface{}) *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation
	}
	r.State = flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(m["state"])

	return r
}

// expandMembershipDisplayNameMap expands the contents of MembershipDisplayName into a JSON
// request object.
func expandMembershipDisplayNameMap(c *Client, f map[string]MembershipDisplayName) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipDisplayName(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipDisplayNameSlice expands the contents of MembershipDisplayName into a JSON
// request object.
func expandMembershipDisplayNameSlice(c *Client, f []MembershipDisplayName) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipDisplayName(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipDisplayNameMap flattens the contents of MembershipDisplayName from a JSON
// response object.
func flattenMembershipDisplayNameMap(c *Client, i interface{}) map[string]MembershipDisplayName {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipDisplayName{}
	}

	if len(a) == 0 {
		return map[string]MembershipDisplayName{}
	}

	items := make(map[string]MembershipDisplayName)
	for k, item := range a {
		items[k] = *flattenMembershipDisplayName(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMembershipDisplayNameSlice flattens the contents of MembershipDisplayName from a JSON
// response object.
func flattenMembershipDisplayNameSlice(c *Client, i interface{}) []MembershipDisplayName {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipDisplayName{}
	}

	if len(a) == 0 {
		return []MembershipDisplayName{}
	}

	items := make([]MembershipDisplayName, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipDisplayName(c, item.(map[string]interface{})))
	}

	return items
}

// expandMembershipDisplayName expands an instance of MembershipDisplayName into a JSON
// request object.
func expandMembershipDisplayName(c *Client, f *MembershipDisplayName) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenMembershipDisplayName flattens an instance of MembershipDisplayName from a JSON
// response object.
func flattenMembershipDisplayName(c *Client, i interface{}) *MembershipDisplayName {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipDisplayName{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipDisplayName
	}
	r.GivenName = dcl.FlattenString(m["givenName"])
	r.FamilyName = dcl.FlattenString(m["familyName"])
	r.FullName = dcl.FlattenString(m["fullName"])

	return r
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumMap flattens the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumMap(c *Client, i interface{}) map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum{}
	}

	if len(a) == 0 {
		return map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum{}
	}

	items := make(map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum)
	for k, item := range a {
		items[k] = *flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(item.(interface{}))
	}

	return items
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumSlice flattens the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumSlice(c *Client, i interface{}) []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum{}
	}

	if len(a) == 0 {
		return []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum{}
	}

	items := make([]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(item.(interface{})))
	}

	return items
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum asserts that an interface is a string, and returns a
// pointer to a *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum with the same value as that string.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(i interface{}) *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	s, ok := i.(string)
	if !ok {
		return MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumRef("")
	}

	return MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumRef(s)
}

// flattenMembershipTypeEnumMap flattens the contents of MembershipTypeEnum from a JSON
// response object.
func flattenMembershipTypeEnumMap(c *Client, i interface{}) map[string]MembershipTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]MembershipTypeEnum{}
	}

	items := make(map[string]MembershipTypeEnum)
	for k, item := range a {
		items[k] = *flattenMembershipTypeEnum(item.(interface{}))
	}

	return items
}

// flattenMembershipTypeEnumSlice flattens the contents of MembershipTypeEnum from a JSON
// response object.
func flattenMembershipTypeEnumSlice(c *Client, i interface{}) []MembershipTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipTypeEnum{}
	}

	if len(a) == 0 {
		return []MembershipTypeEnum{}
	}

	items := make([]MembershipTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipTypeEnum(item.(interface{})))
	}

	return items
}

// flattenMembershipTypeEnum asserts that an interface is a string, and returns a
// pointer to a *MembershipTypeEnum with the same value as that string.
func flattenMembershipTypeEnum(i interface{}) *MembershipTypeEnum {
	s, ok := i.(string)
	if !ok {
		return MembershipTypeEnumRef("")
	}

	return MembershipTypeEnumRef(s)
}

// flattenMembershipDeliverySettingEnumMap flattens the contents of MembershipDeliverySettingEnum from a JSON
// response object.
func flattenMembershipDeliverySettingEnumMap(c *Client, i interface{}) map[string]MembershipDeliverySettingEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipDeliverySettingEnum{}
	}

	if len(a) == 0 {
		return map[string]MembershipDeliverySettingEnum{}
	}

	items := make(map[string]MembershipDeliverySettingEnum)
	for k, item := range a {
		items[k] = *flattenMembershipDeliverySettingEnum(item.(interface{}))
	}

	return items
}

// flattenMembershipDeliverySettingEnumSlice flattens the contents of MembershipDeliverySettingEnum from a JSON
// response object.
func flattenMembershipDeliverySettingEnumSlice(c *Client, i interface{}) []MembershipDeliverySettingEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipDeliverySettingEnum{}
	}

	if len(a) == 0 {
		return []MembershipDeliverySettingEnum{}
	}

	items := make([]MembershipDeliverySettingEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipDeliverySettingEnum(item.(interface{})))
	}

	return items
}

// flattenMembershipDeliverySettingEnum asserts that an interface is a string, and returns a
// pointer to a *MembershipDeliverySettingEnum with the same value as that string.
func flattenMembershipDeliverySettingEnum(i interface{}) *MembershipDeliverySettingEnum {
	s, ok := i.(string)
	if !ok {
		return MembershipDeliverySettingEnumRef("")
	}

	return MembershipDeliverySettingEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Membership) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalMembership(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Group == nil && ncr.Group == nil {
			c.Config.Logger.Info("Both Group fields null - considering equal.")
		} else if nr.Group == nil || ncr.Group == nil {
			c.Config.Logger.Info("Only one Group field is null - considering unequal.")
			return false
		} else if *nr.Group != *ncr.Group {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type membershipDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         membershipApiOperation
}

func convertFieldDiffsToMembershipDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]membershipDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []membershipDiff
	// For each operation name, create a membershipDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := membershipDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToMembershipApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToMembershipApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (membershipApiOperation, error) {
	switch opName {

	case "updateMembershipUpdateMembershipOperation":
		return &updateMembershipUpdateMembershipOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractMembershipFields(r *Membership) error {
	vPreferredMemberKey := r.PreferredMemberKey
	if vPreferredMemberKey == nil {
		// note: explicitly not the empty object.
		vPreferredMemberKey = &MembershipPreferredMemberKey{}
	}
	if err := extractMembershipPreferredMemberKeyFields(r, vPreferredMemberKey); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPreferredMemberKey) {
		r.PreferredMemberKey = vPreferredMemberKey
	}
	vDisplayName := r.DisplayName
	if vDisplayName == nil {
		// note: explicitly not the empty object.
		vDisplayName = &MembershipDisplayName{}
	}
	if err := extractMembershipDisplayNameFields(r, vDisplayName); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDisplayName) {
		r.DisplayName = vDisplayName
	}
	return nil
}
func extractMembershipPreferredMemberKeyFields(r *Membership, o *MembershipPreferredMemberKey) error {
	return nil
}
func extractMembershipRolesFields(r *Membership, o *MembershipRoles) error {
	vExpiryDetail := o.ExpiryDetail
	if vExpiryDetail == nil {
		// note: explicitly not the empty object.
		vExpiryDetail = &MembershipRolesExpiryDetail{}
	}
	if err := extractMembershipRolesExpiryDetailFields(r, vExpiryDetail); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vExpiryDetail) {
		o.ExpiryDetail = vExpiryDetail
	}
	vRestrictionEvaluations := o.RestrictionEvaluations
	if vRestrictionEvaluations == nil {
		// note: explicitly not the empty object.
		vRestrictionEvaluations = &MembershipRolesRestrictionEvaluations{}
	}
	if err := extractMembershipRolesRestrictionEvaluationsFields(r, vRestrictionEvaluations); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vRestrictionEvaluations) {
		o.RestrictionEvaluations = vRestrictionEvaluations
	}
	return nil
}
func extractMembershipRolesExpiryDetailFields(r *Membership, o *MembershipRolesExpiryDetail) error {
	return nil
}
func extractMembershipRolesRestrictionEvaluationsFields(r *Membership, o *MembershipRolesRestrictionEvaluations) error {
	vMemberRestrictionEvaluation := o.MemberRestrictionEvaluation
	if vMemberRestrictionEvaluation == nil {
		// note: explicitly not the empty object.
		vMemberRestrictionEvaluation = &MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}
	if err := extractMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationFields(r, vMemberRestrictionEvaluation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMemberRestrictionEvaluation) {
		o.MemberRestrictionEvaluation = vMemberRestrictionEvaluation
	}
	return nil
}
func extractMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationFields(r *Membership, o *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) error {
	return nil
}
func extractMembershipDisplayNameFields(r *Membership, o *MembershipDisplayName) error {
	return nil
}

func postReadExtractMembershipFields(r *Membership) error {
	vPreferredMemberKey := r.PreferredMemberKey
	if vPreferredMemberKey == nil {
		// note: explicitly not the empty object.
		vPreferredMemberKey = &MembershipPreferredMemberKey{}
	}
	if err := postReadExtractMembershipPreferredMemberKeyFields(r, vPreferredMemberKey); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPreferredMemberKey) {
		r.PreferredMemberKey = vPreferredMemberKey
	}
	vDisplayName := r.DisplayName
	if vDisplayName == nil {
		// note: explicitly not the empty object.
		vDisplayName = &MembershipDisplayName{}
	}
	if err := postReadExtractMembershipDisplayNameFields(r, vDisplayName); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDisplayName) {
		r.DisplayName = vDisplayName
	}
	return nil
}
func postReadExtractMembershipPreferredMemberKeyFields(r *Membership, o *MembershipPreferredMemberKey) error {
	return nil
}
func postReadExtractMembershipRolesFields(r *Membership, o *MembershipRoles) error {
	vExpiryDetail := o.ExpiryDetail
	if vExpiryDetail == nil {
		// note: explicitly not the empty object.
		vExpiryDetail = &MembershipRolesExpiryDetail{}
	}
	if err := extractMembershipRolesExpiryDetailFields(r, vExpiryDetail); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vExpiryDetail) {
		o.ExpiryDetail = vExpiryDetail
	}
	vRestrictionEvaluations := o.RestrictionEvaluations
	if vRestrictionEvaluations == nil {
		// note: explicitly not the empty object.
		vRestrictionEvaluations = &MembershipRolesRestrictionEvaluations{}
	}
	if err := extractMembershipRolesRestrictionEvaluationsFields(r, vRestrictionEvaluations); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vRestrictionEvaluations) {
		o.RestrictionEvaluations = vRestrictionEvaluations
	}
	return nil
}
func postReadExtractMembershipRolesExpiryDetailFields(r *Membership, o *MembershipRolesExpiryDetail) error {
	return nil
}
func postReadExtractMembershipRolesRestrictionEvaluationsFields(r *Membership, o *MembershipRolesRestrictionEvaluations) error {
	vMemberRestrictionEvaluation := o.MemberRestrictionEvaluation
	if vMemberRestrictionEvaluation == nil {
		// note: explicitly not the empty object.
		vMemberRestrictionEvaluation = &MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}
	if err := extractMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationFields(r, vMemberRestrictionEvaluation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMemberRestrictionEvaluation) {
		o.MemberRestrictionEvaluation = vMemberRestrictionEvaluation
	}
	return nil
}
func postReadExtractMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationFields(r *Membership, o *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) error {
	return nil
}
func postReadExtractMembershipDisplayNameFields(r *Membership, o *MembershipDisplayName) error {
	return nil
}