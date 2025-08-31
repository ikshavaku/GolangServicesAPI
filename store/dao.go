package store

type ServiceDAO struct {
	ID          string
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
	IsDeleted   bool
}

type ServiceVersionDAO struct {
	ServiceID     string
	VersionNumber string
	CreatedAt     string
	UpdatedAt     string
	IsDeleted     bool
}
