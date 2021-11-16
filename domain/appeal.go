package domain

import (
	"time"
)

const (
	AppealActionNameApprove = "approve"
	AppealActionNameReject  = "reject"

	AppealStatusPending    = "pending"
	AppealStatusCanceled   = "canceled"
	AppealStatusActive     = "active"
	AppealStatusRejected   = "rejected"
	AppealStatusTerminated = "terminated"

	SystemActorName = "system"

	DefaultAppealAccountType = "user"
)

// AppealOptions
type AppealOptions struct {
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	Duration       string     `json:"duration"`
}

// Appeal struct
type Appeal struct {
	ID            uint                   `json:"id"`
	ResourceID    uint                   `json:"resource_id"`
	PolicyID      string                 `json:"policy_id"`
	PolicyVersion uint                   `json:"policy_version"`
	Status        string                 `json:"status"`
	AccountID     string                 `json:"account_id"`
	AccountType   string                 `json:"account_type" default:"user"`
	CreatedBy     string                 `json:"created_by"`
	Creator       interface{}            `json:"creator"`
	Role          string                 `json:"role"`
	Options       *AppealOptions         `json:"options"`
	Details       map[string]interface{} `json:"details"`
	Labels        map[string]string      `json:"labels"`

	RevokedBy    string    `json:"revoked_by"`
	RevokedAt    time.Time `json:"revoked_at"`
	RevokeReason string    `json:"revoke_reason"`

	Policy    *Policy     `json:"-"`
	Resource  *Resource   `json:"resource,omitempty"`
	Approvals []*Approval `json:"approvals,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Appeal) GetNextPendingApproval() *Approval {
	for _, approval := range a.Approvals {
		if approval.Status == ApprovalStatusPending && approval.IsManualApproval() {
			return approval
		}
	}
	return nil
}

func (a *Appeal) Terminate() {
	a.Status = AppealStatusTerminated
}

func (a *Appeal) Activate() error {
	a.Status = AppealStatusActive

	if a.Options != nil && a.Options.Duration != "" {
		duration, err := time.ParseDuration(a.Options.Duration)
		if err != nil {
			return err
		}

		expirationDate := time.Now().Add(duration)
		a.Options.ExpirationDate = &expirationDate
	}

	return nil
}

type ApprovalAction struct {
	AppealID     uint   `validate:"required"`
	ApprovalName string `validate:"required"`
	Actor        string `validate:"email"`
	Action       string `validate:"required,oneof=approve reject"`
}

// AppealRepository interface
type AppealRepository interface {
	BulkUpsert([]*Appeal) error
	Find(map[string]interface{}) ([]*Appeal, error) // TODO: create ListAppealsFilter as the filter param type
	GetByID(uint) (*Appeal, error)
	Update(*Appeal) error
}

// AppealService interface
type AppealService interface {
	Create([]*Appeal) error
	Find(map[string]interface{}) ([]*Appeal, error)
	GetByID(uint) (*Appeal, error)
	MakeAction(ApprovalAction) (*Appeal, error)
	Cancel(uint) (*Appeal, error)
	Revoke(id uint, actor, reason string) (*Appeal, error)
}