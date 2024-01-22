package server

import (
	"context"

	audit "github.com/AngelicaNice/AuditLog/pkg/domain"
)

type AuditService interface {
	Insert(ctx context.Context, req *audit.LogRequest) error
}

type AuditServer struct {
	auditService AuditService
}

func NewAuditServer(auditService AuditService) *AuditServer {
	return &AuditServer{
		auditService: auditService,
	}
}

func (s *AuditServer) Log(ctx context.Context, r *audit.LogRequest) (*audit.Empty, error) {
	err := s.auditService.Insert(ctx, r)

	return &audit.Empty{}, err
}
