package ipa

import (
	"encoding/json"
)

type Group struct {
	CN string `json:"cn"`
	ID string `json:"id"`
}

type GroupResponse struct {
	Objectclass        []string `json:"objectclass"`
	Ipauniqueid        []string `json:"ipauniqueid"`
	Cn                 []string `json:"cn"`
	Gidnumber          []string `json:"gidnumber"`
	MemberUser         []string `json:"member_user,omitempty"`
	Dn                 string   `json:"dn"`
	Description        []string `json:"description,omitempty"`
	MemberGroup        []string `json:"member_group,omitempty"`
	MemberindirectUser []string `json:"memberindirect_user,omitempty"`
}

type GroupListResponse struct {
	Result struct {
		Result    []GroupResponse `json:"result"`
		Count     int             `json:"count"`
		Truncated bool            `json:"truncated"`
		Summary   string          `json:"summary"`
	} `json:"result"`
	Error     interface{} `json:"error"`
	Id        interface{} `json:"id"`
	Principal string      `json:"principal"`
	Version   string      `json:"version"`
}

func (c *Client) GroupFind(criteria string, options Options) ([]*Group, error) {
	if options == nil {
		options = Options{}
	}

	options["all"] = true

	body, err := c.RequestBody("group_find", []string{criteria}, options)
	if err != nil {
		return nil, err
	}

	res := new(GroupListResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	groups := make([]*Group, 0)
	for _, t := range res.Result.Result {
		if len(t.Cn) < 1 || len(t.Ipauniqueid) < 1 {
			continue
		}

		group := Group{
			CN: t.Cn[0],
			ID: t.Ipauniqueid[0],
		}
		groups = append(groups, &group)
	}

	return groups, nil
}
