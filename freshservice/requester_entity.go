package freshservice

import (
	"fmt"
	"strings"
	"time"
)

// Requesters holds a list of Freshservice Requesters
type Requesters struct {
	List        []RequesterDetails `json:"requesters"`
	Description string             `json:"description"`
	Errors      []Error            `json:"errors"`
}

// Requester holds the details of a specific Freshservice Requester
type Requester struct {
	Details     RequesterDetails `json:"requester"`
	Description string           `json:"description"`
	Errors      []Error          `json:"errors"`
}

// RequesterDetails contains the details of a specific Freshservice Requester

type RequesterDetails struct {
	ID                                        int      `json:"id"`
	FirstName                                 string   `json:"first_name"`
	LastName                                  string   `json:"last_name"`
	JobTitle                                  string   `json:"job_title"`
	PrimaryEmail                              string   `json:"primary_email"`
	SecondaryEmails                           []string `json:"secondary_emails"`
	WorkPhoneNumber                           string   `json:"work_phone_number"`
	MobilePhoneNumber                         string   `json:"mobile_phone_number"`
	DepartmentIDs                             []int    `json:"department_ids"`
	CanSeeAllTicketsFromAssociatedDepartments bool     `json:"can_see_all_tickets_from_associated_departments"`
	ReportingManagerID                        int      `json:"reporting_manager_id"`
	Address                                   string   `json:"address"`
	TimeZone                                  string   `json:"time_zone"`
	TimeFormat                                string   `json:"time_format"`
	Language                                  string   `json:"language"`
	LocationID                                int      `json:"location_id"`
	BackgroundInformation                     string   `json:"background_information"`
	CustomFields                              struct {
		House string `json:"house"`
	} `json:"custom_fields"`
	Active           bool      `json:"active"`
	HasLoggedIn      bool      `json:"has_logged_in"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	IsRequesterGroup bool      `json:"is_Requester"`
}

func (r *RequesterDetails) Validate() error {
	validTimeFormats := []string{
		"24h",
		"12h",
	}

	if !StringInSlice(r.TimeFormat, validTimeFormats) {
		return fmt.Errorf("Time format is invalid; choose from %s", strings.Join(validTimeFormats, ","))
	}

	return nil
}

// RequesterListFilter holds the filters available when listing Freservice Requesters
type RequesterListFilter struct {
	PageQuery     string
	Email         *string
	MobilePhone   *int
	WorkPhone     *int
	Active        bool
	IncludeAgents bool
}

// QueryString allows the available filter items to meet the QueryFilter interface
func (rf *RequesterListFilter) QueryString() string {
	var qs []string
	if rf.PageQuery != "" {
		qs = append(qs, rf.PageQuery)
	}

	switch {
	case rf.Email != nil:
		qs = append(qs, fmt.Sprintf("email=%s", *rf.Email))
	case rf.MobilePhone != nil:
		qs = append(qs, fmt.Sprintf("mobile_phone_number=%d", *rf.MobilePhone))
	case rf.WorkPhone != nil:
		qs = append(qs, fmt.Sprintf("work_phone_number=%d", *rf.WorkPhone))
	case rf.Active:
		qs = append(qs, fmt.Sprintf("active=%v", rf.Active))
	case rf.IncludeAgents:
		qs = append(qs, fmt.Sprintf("include_agents=%v", rf.IncludeAgents))
	}
	return strings.Join(qs, "&")
}
