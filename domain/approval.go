package domain

import "time"

const (
	ApprovalStatusPending  = "pending"
	ApprovalStatusBlocked  = "blocked"
	ApprovalStatusSkipped  = "skipped"
	ApprovalStatusApproved = "approved"
	ApprovalStatusRejected = "rejected"
)

type Approval struct {
	ID            uint    `json:"id" yaml:"id"`
	Name          string  `json:"name" yaml:"name"`
	Index         int     `json:"-" yaml:"-"`
	AppealID      uint    `json:"appeal_id" yaml:"appeal_id"`
	Status        string  `json:"status" yaml:"status"`
	Actor         *string `json:"actor" yaml:"actor"`
	Reason        string  `json:"reason,omitempty" yaml:"reason,omitempty"`
	PolicyID      string  `json:"policy_id" yaml:"policy_id"`
	PolicyVersion uint    `json:"policy_version" yaml:"policy_version"`

	Approvers []string `json:"approvers,omitempty" yaml:"approvers,omitempty"`
	Appeal    *Appeal  `json:"appeal,omitempty" yaml:"appeal,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}

func (a *Approval) IsManualApproval() bool {
	return len(a.Approvers) > 0
}

type ListApprovalsFilter struct {
	AccountID string   `mapstructure:"account_id" validate:"omitempty,required"`
	Statuses  []string `mapstructure:"statuses" validate:"omitempty,min=1"`
}

type ApprovalRepository interface {
	BulkInsert([]*Approval) error
	ListApprovals(*ListApprovalsFilter) ([]*Approval, error)
}

type ApprovalService interface {
	BulkInsert([]*Approval) error
	ListApprovals(*ListApprovalsFilter) ([]*Approval, error)
	AdvanceApproval(appeal *Appeal) error
}