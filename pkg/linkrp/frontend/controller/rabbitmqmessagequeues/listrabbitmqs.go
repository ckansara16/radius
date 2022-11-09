// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package rabbitmqmessagequeues

import (
	"context"
	"net/http"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	ctrl "github.com/project-radius/radius/pkg/armrpc/frontend/controller"
	"github.com/project-radius/radius/pkg/linkrp/datamodel"
	"github.com/project-radius/radius/pkg/linkrp/datamodel/converter"

	"github.com/project-radius/radius/pkg/armrpc/rest"
	"github.com/project-radius/radius/pkg/ucp/store"
)

var _ ctrl.Controller = (*ListRabbitMQMessageQueues)(nil)

// ListRabbitMQMessageQueues is the controller implementation to get the list of rabbitMQ link resources in the resource group.
type ListRabbitMQMessageQueues struct {
	ctrl.BaseController
}

// NewListRabbitMQMessageQueues creates a new instance of ListRabbitMQMessageQueues.
func NewListRabbitMQMessageQueues(opts ctrl.Options) (ctrl.Controller, error) {
	return &ListRabbitMQMessageQueues{ctrl.NewBaseController(opts)}, nil
}

func (rabbitmq *ListRabbitMQMessageQueues) Run(ctx context.Context, w http.ResponseWriter, req *http.Request) (rest.Response, error) {
	serviceCtx := v1.ARMRequestContextFromContext(ctx)

	query := store.Query{
		RootScope:    serviceCtx.ResourceID.RootScope(),
		ResourceType: serviceCtx.ResourceID.Type(),
	}

	result, err := rabbitmq.StorageClient().Query(ctx, query, store.WithPaginationToken(serviceCtx.SkipToken), store.WithMaxQueryItemCount(serviceCtx.Top))
	if err != nil {
		return nil, err
	}

	paginatedList, err := rabbitmq.createPaginatedList(ctx, req, result)

	return rest.NewOKResponse(paginatedList), err
}

func (rabbitmq *ListRabbitMQMessageQueues) createPaginatedList(ctx context.Context, req *http.Request, result *store.ObjectQueryResult) (*v1.PaginatedList, error) {
	serviceCtx := v1.ARMRequestContextFromContext(ctx)

	items := []interface{}{}
	for _, item := range result.Items {
		dm := &datamodel.RabbitMQMessageQueue{}
		if err := item.As(dm); err != nil {
			return nil, err
		}

		versioned, err := converter.RabbitMQMessageQueueDataModelToVersioned(dm, serviceCtx.APIVersion)
		if err != nil {
			return nil, err
		}

		items = append(items, versioned)
	}

	return &v1.PaginatedList{
		Value:    items,
		NextLink: ctrl.GetNextLinkURL(ctx, req, result.PaginationToken),
	}, nil
}
