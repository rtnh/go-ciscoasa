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

// PhysicalInterfaceCollection represents a collection of physical interfaces.
type PhysicalInterfaceCollection struct {
	RangeInfo RangeInfo            `json:"rangeInfo"`
	Items     []*PhysicalInterface `json:"items"`
	Kind      string               `json:"kind"`
	SelfLink  string               `json:"selfLink"`
}

// Interface represents an interface.
type PhysicalInterface struct {
	HardwareID        string     `json:"hardwareID"`
	InterfaceDesc     string     `json:"interfaceDesc"`
	ChannelGroupID    string     `json:"channelGroupID"`
	ChannelGroupMode  string     `json:"channelGroupMode"`
	Duplex            string     `json:"duplex"`
	FlowcontrolOn     bool       `json:"flowcontrolOn"`
	FlowcontrolHigh   int        `json:"flowcontrolHigh"`
	FlowcontrolLow    int        `json:"flowcontrolLow"`
	FlowcontrolPeriod int        `json:"flowcontrolPeriod"`
	ForwardTrafficCX  bool       `json:"forwardTrafficCX"`
	ForwardTrafficSFR bool       `json:"forwardTrafficSFR"`
	LacpPriority      int        `json:"lacpPriority"`
	ActiveMacAddress  string     `json:"activeMacAddress"`
	StandByMacAddress string     `json:"standByMacAddress"`
	ManagementOnly    bool       `json:"managementOnly"`
	Mtu               int        `json:"mtu"`
	Name              string     `json:"name"`
	SecurityLevel     int        `json:"securityLevel"`
	Shutdown          bool       `json:"shutdown"`
	Speed             string     `json:"speed"`
	IPAddress         *IPAddress `json:"ipAddress"`
	Ipv6Info          *IPv6Info  `json:"ipv6Info"`
	Kind              string     `json:"kind"`
	ObjectID          string     `json:"objectId"`
	SelfLink          string     `json:"selfLink"`
}

// ListPhysicalInterfaces returns a collection of interfaces.
func (s *interfaceService) ListPhysicalInterfaces() (*PhysicalInterfaceCollection, error) {
	u := "/api/interfaces/physical"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	e := &PhysicalInterfaceCollection{}
	_, err = s.do(req, e)

	return e, err
}
