package domain

import (
	"time"
)

// Client representa la tabla client en la base de datos.
type Client struct {
	ID               uint64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	FirstName        string
	LastName         *string
	ClientTypeID     uint64
	ClientType       ClientType
	Document         string
	Email            string
	Profession       *string
	Address          string
	Precinct         string
	MunicipalityCode int
	Municipality     Municipality
	DepartmentCode   int
	Department       Department
	StatusID         uint64
	Status           ClientStatus
	Latitude         *string
	Longitude        *string
	Nap              *string
	AccessPoint      *string
	PhoneNumbers     []string
}

// ClientType representa la tabla client_type en la base de datos.
type ClientType struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

// ClientStatus representa la tabla client_status en la base de datos.
type ClientStatus struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
}

// Municipality representa la tabla municipality en la base de datos.
type Municipality struct {
	Code      int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Department representa la tabla department en la base de datos.
type Department struct {
	Code      int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
