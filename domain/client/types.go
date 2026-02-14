package client

type PostExampleRequest struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

type PostExampleResponse struct {
	ID int `json:"ID"`
}

type GetExampleRequest struct {
	ID int `json:"ID"`
}

type GetExampleResponse struct {
	NumberOfEntries int `json:"numberOfEntries"`
}
