package apis

import (
	"context"
	"fmt"

	apisv1 "github.com/unkeyed/unkey/apps/agent/gen/proto/apis/v1"
	"github.com/unkeyed/unkey/apps/agent/pkg/errors"
)

func (s *service) DeleteApi(ctx context.Context, req *apisv1.DeleteApiRequest) (*apisv1.DeleteApiResponse, error) {
	api, found, err := s.database.FindApi(ctx, req.ApiId)
	if err != nil {
		return nil, errors.New(errors.ErrInternalServerError, err)
	}
	if !found {
		return nil, errors.New(errors.ErrNotFound, fmt.Errorf("api %s does not exist", req.ApiId))
	}
	if api.WorkspaceId != req.AuthorizedWorkspaceId {
		return nil, errors.New(errors.ErrUnauthorized, fmt.Errorf("access to workspace denied"))
	}

	err = s.database.DeleteApi(ctx, req.ApiId)
	if err != nil {
		return nil, errors.New(errors.ErrInternalServerError, err)
	}
	return &apisv1.DeleteApiResponse{}, nil
}
