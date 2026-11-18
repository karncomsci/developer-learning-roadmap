package httpapi

import (
	"context"

	"github.com/gofiber/fiber/v3"
)

// Server implements the OpenAPI contract. Domain rules live outside this
// package (ADR 0002, ADR 0004); handlers only translate HTTP <-> domain.
type Server struct{}

var _ StrictServerInterface = (*Server)(nil)

func NewServer() *Server { return &Server{} }

func (s *Server) GetHealth(ctx context.Context, _ GetHealthRequestObject) (GetHealthResponseObject, error) {
	return GetHealth200JSONResponse{Status: Ok}, nil
}

// Endpoints below are implemented in later milestones (docs/SPEC.md §12).

func (s *Server) ListRoles(context.Context, ListRolesRequestObject) (ListRolesResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) CreateRole(context.Context, CreateRoleRequestObject) (CreateRoleResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) GetRoleBySlug(context.Context, GetRoleBySlugRequestObject) (GetRoleBySlugResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) UpdateRole(context.Context, UpdateRoleRequestObject) (UpdateRoleResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) DeleteRole(context.Context, DeleteRoleRequestObject) (DeleteRoleResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) CreateTopic(context.Context, CreateTopicRequestObject) (CreateTopicResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) UpdateTopic(context.Context, UpdateTopicRequestObject) (UpdateTopicResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) DeleteTopic(context.Context, DeleteTopicRequestObject) (DeleteTopicResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) SetTopicStatus(context.Context, SetTopicStatusRequestObject) (SetTopicStatusResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) CreateResource(context.Context, CreateResourceRequestObject) (CreateResourceResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) UpdateResource(context.Context, UpdateResourceRequestObject) (UpdateResourceResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) DeleteResource(context.Context, DeleteResourceRequestObject) (DeleteResourceResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) UploadTopicImage(context.Context, UploadTopicImageRequestObject) (UploadTopicImageResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}

func (s *Server) DeleteTopicImage(context.Context, DeleteTopicImageRequestObject) (DeleteTopicImageResponseObject, error) {
	return nil, fiber.ErrNotImplemented
}
