package service

type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ListServicesResponse struct {
	Total int       `json:"total"`
	Size  int       `json:"size"`
	Page  int       `json:"page"`
	Data  []Service `json:"data"`
}

type GetServiceResponse struct {
	Service
}

type ServiceVersion struct {
	ServiceID string `json:"service_id"`
	VersionID string `json:"version_id"`
	CreatedAt string `json:"created_at"`
	Active    bool   `json:"is_active"`
}

type ListServiceVersionsByServiceIDResponse struct {
	Data  []ServiceVersion `json:"data"`
	Total int              `json:"total"`
}
