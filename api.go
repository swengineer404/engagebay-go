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

func (a *API) GetContactByEmail(email string) (*Contact, error) {
	var c Contact
	return &c, a.client.do(fmt.Sprintf("/subscribers/contact-by-email/%s", email), http.MethodGet, "", nil, &c)
}

func (a *API) GetContactTags(id int) (tags Tags, err error) {
	return tags, a.client.do(fmt.Sprintf("/subscribers/get-tags-by-id/%d", id), http.MethodGet, "", nil, &tags)
}

func (a *API) AddContactTags(id int, tags ...string) error {
	var t Tags
	t.Add(tags...)
	return a.client.do(fmt.Sprintf("/subscribers/contact/tags/add2/%d", id), http.MethodPost, ContentTypeJson, t, nil)
}

func (a *API) DeleteContactTags(id int, tags ...string) error {
	var t Tags
	t.Add(tags...)
	return a.client.do(fmt.Sprintf("/subscribers/contact/tags/delete/%d", id), http.MethodPost, ContentTypeJson, t, nil)
}
