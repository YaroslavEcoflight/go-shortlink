package restapi

import "shortlink-service/internal/domain/service"

type LinkHandler struct {
	svc service.LinkService
}

func NewLinkHandler(svc service.LinkService) LinkHandler {
	return LinkHandler{svc: svc}
}
