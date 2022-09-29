package web

type RoleCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	DeleteBy    string `json:"delete_by"`
}

type RoleUpdateRequest struct {
	ID          uint64
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	DeleteBy    string `json:"delete_by"`
}