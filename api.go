package engagebay

import (
	"fmt"
	"net/http"
)

type API struct {
	client *restClient
}

func New(key string) *API {
	return &API{
		client: newAPIClient(key),
	}
}

func (a *API) CreateContact(params *CreateContactParams) (*Contact, error) {
	var c Contact
	return &c, a.client.do("/subscribers/subscriber", http.MethodPost, ContentTypeJson, params, &c)
}

func (a *API) GetContact(id int) (*Contact, error) {
	var c Contact
	return &c, a.client.do(fmt.Sprintf("/subscribers/%d", id), http.MethodGet, "", nil, &c)
}

func (a *API) GetContactTags(id int) (tags Tags, err error) {
	return tags, a.client.do(fmt.Sprintf("/subscribers/get-tags-by-id/%d", id), http.MethodGet, "", nil, &tags)
}

func (a *API) AddContactTags(params *AddContactTagsParams) error {
	return a.client.do(fmt.Sprintf("/subscribers/contact/tags/add2/%d", params.ID), http.MethodPost, ContentTypeJson, params.Tags, nil)
}
