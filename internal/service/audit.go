package service

import (
	"context"

	audit "github.com/AngelicaNice/AuditLog/pkg/domain"
)

type Repository interface {
	Insert(ctx context.Context, item audit.LogItem) error
}

type AuditService struct {
	repo Repository
}

func NewAudit(repo Repository) *AuditService {
	return &AuditService{
		repo: repo,
	}
}

func (s *AuditService) Insert(ctx context.Context, req *audit.LogRequest) error {
	item := audit.LogItem{
		Entity:    req.Action.String(),
		Action:    req.Entity.String(),
		EntityID:  req.EntityId,
		Timestamp: req.Timestamp.AsTime(),
	}
	return s.repo.Insert(ctx, item)
}
