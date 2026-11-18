package httpapi

import (
	"context"

	"github.com/gofiber/fiber/v3"

	"github.com/karncomsci/developer-learning-roadmap/api/internal/store"
)

// Server implements the OpenAPI contract. Domain rules live outside this
// package (ADR 0002, ADR 0004); handlers only translate HTTP <-> domain.
type Server struct {
	q *store.Queries
}

var _ StrictServerInterface = (*Server)(nil)

func NewServer(q *store.Queries) *Server { return &Server{q: q} }

func (s *Server) GetHealth(ctx context.Context, _ GetHealthRequestObject) (GetHealthResponseObject, error) {
	return GetHealth200JSONResponse{Status: Ok}, nil
}

func (s *Server) ListRoles(ctx context.Context, _ ListRolesRequestObject) (ListRolesResponseObject, error) {
	rows, err := s.q.ListRolesWithProgress(ctx)
	if err != nil {
		return nil, err
	}
	out := make(ListRoles200JSONResponse, 0, len(rows))
	for _, r := range rows {
		out = append(out, RoleSummary{
			Id:          r.ID,
			Slug:        r.Slug,
			Name:        r.Name,
			Description: r.Description,
			Icon:        r.Icon,
			Color:       r.Color,
			Progress: RoleProgress{
				Done:     int(r.DoneCount),
				Learning: int(r.LearningCount),
				Total:    int(r.TotalCount),
			},
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}
	return out, nil
}

// Endpoints below are implemented in later milestones (docs/SPEC.md §12).

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
