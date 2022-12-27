package engagebay

type API struct {
	client *restClient
}

func New(key string) *API {
	return &API{
		client: newAPIClient(key),
	}
}

func (a *API) CreateContact(params *CreateContactParams) (c *Contact, err error) {
	return c, a.client.do("/subscribers/subscriber", "POST", params, c)
}
