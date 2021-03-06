//
// Copyright 2017, Rutger te Nijenhuis & Sander van Harmelen
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
//

package ciscoasa

import "fmt"

// ExtendedACLObjectCollection represents a collection of access control list objects.
type ExtendedACLObjectCollection struct {
	RangeInfo RangeInfo            `json:"rangeInfo"`
	Items     []*ExtendedACLObject `json:"items"`
	Kind      string               `json:"kind"`
	SelfLink  string               `json:"selfLink"`
}

// ExtendedACLObject represents an access control list object.
type ExtendedACLObject struct {
	Name     string `json:"name,omitempty"`
	Kind     string `json:"kind,omitempty"`
	ObjectID string `json:"objectId,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// ExtendedACEObjectCollection represents a collection of access control element objects.
type ExtendedACEObjectCollection struct {
	RangeInfo RangeInfo            `json:"rangeInfo"`
	Items     []*ExtendedACEObject `json:"items"`
	Kind      string               `json:"kind"`
	SelfLink  string               `json:"selfLink"`
}

// ExtendedACEObject represents an access control element object
type ExtendedACEObject struct {
	SrcAddress   *AddressObject `json:"sourceAddress,omitempty"`
	SrcService   *ServiceObject `json:"sourceService,omitempty"`
	DstAddress   *AddressObject `json:"destinationAddress,omitempty"`
	DstService   *ServiceObject `json:"destinationService,omitempty"`
	RuleLogging  *RuleLogging   `json:"ruleLogging,omitempty"`
	Position     int            `json:"position,omitempty"`
	Permit       bool           `json:"permit,omitempty"`
	Active       bool           `json:"active"`
	IsAccessRule bool           `json:"isAccessRule"`
	Remarks      []string       `json:"remarks,omitempty"`
	Kind         string         `json:"kind,omitempty"`
	ObjectID     string         `json:"objectId,omitempty"`
	SelfLink     string         `json:"selfLink,omitempty"`
}

// RuleLogging represents the rule logging settings
type RuleLogging struct {
	LogStatus   string `json:"logStatus,omitempty"`
	LogInterval int    `json:"logInterval,omitempty"`
}

// ListExtendedACLs returns a collection of access control list objects.
func (s *objectsService) ListExtendedACLs() (*ExtendedACLObjectCollection, error) {
	u := "/api/objects/extendedacls"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	e := &ExtendedACLObjectCollection{}
	_, err = s.do(req, e)

	return e, err
}

// ListExtendedACLACEs returns a collection of access control element objects.
func (s *objectsService) ListExtendedACLACEs(aclName string) (*ExtendedACEObjectCollection, error) {
	u := fmt.Sprintf("/api/objects/extendedacls/%s/aces", aclName)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	e := &ExtendedACEObjectCollection{}
	_, err = s.do(req, e)

	return e, err
}

// CreateExtendedACLACE creates an access control element.
func (s *objectsService) CreateExtendedACLACE(aclName, src, srcService, dst, dstService string, active, permit bool) (string, error) {
	u := fmt.Sprintf("/api/objects/extendedacls/%s/aces", aclName)

	e := &ExtendedACEObject{
		Active: active,
		Permit: permit,
		Kind:   "object#ExtendedACE",
	}

	var err error
	if e.SrcAddress, err = s.Objects.objectFromAddress(src); err != nil {
		return "", err
	}
	if e.SrcService, err = s.Objects.objectFromService(srcService); err != nil {
		return "", err
	}
	if e.DstAddress, err = s.Objects.objectFromAddress(dst); err != nil {
		return "", err
	}
	if e.DstService, err = s.Objects.objectFromService(dstService); err != nil {
		return "", err
	}

	req, err := s.newRequest("POST", u, e)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// GetExtendedACLACE retrieves an access control element.
func (s *objectsService) GetExtendedACLACE(aclName string, aceID string) (*ExtendedACEObject, error) {
	u := fmt.Sprintf("/api/objects/extendedacls/%s/aces/%s", aclName, aceID)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	e := &ExtendedACEObject{}
	_, err = s.do(req, e)

	return e, err
}

// UpdateExtendedACLACE updates an access control element.
func (s *objectsService) UpdateExtendedACLACE(aclName, aceID, src, srcService, dst, dstService string, active, permit bool) (string, error) {
	u := fmt.Sprintf("/api/objects/extendedacls/%s/aces/%s", aclName, aceID)

	e := &ExtendedACEObject{
		Active: active,
		Permit: permit,
		Kind:   "object#ExtendedACE",
	}

	var err error
	if e.SrcAddress, err = s.Objects.objectFromAddress(src); err != nil {
		return "", err
	}
	if e.SrcService, err = s.Objects.objectFromService(src); err != nil {
		return "", err
	}
	if e.DstAddress, err = s.Objects.objectFromAddress(dst); err != nil {
		return "", err
	}
	if e.DstService, err = s.Objects.objectFromService(dstService); err != nil {
		return "", err
	}

	req, err := s.newRequest("POST", u, e)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// DeleteExtendedACLACE deletes an access control element.
func (s *objectsService) DeleteExtendedACLACE(aclName string, aceID string) error {
	u := fmt.Sprintf("/api/objects/extendedacls/%s/aces/%s", aclName, aceID)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
