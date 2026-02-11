package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tomatosAt/IT01-api/module/front-end/dto"
	"github.com/tomatosAt/IT01-api/pkg/util"
)

func (h *Handler) RegisterUserHandler(ctx *fiber.Ctx) error {
	var payload dto.UserPayload
	if err := ctx.BodyParser(&payload); err != nil {
		return util.HttpError(ctx, http.StatusBadRequest, err.Error())
	}
	res, status, err := h.svc.UserSVC(ctx.UserContext(), payload)
	if err != nil {
		return util.HttpError(ctx, http.StatusInternalServerError, err.Error())
	}
	return util.HttpSuccess(ctx, status, res)
}

func (h *Handler) UserDashBoardHandler(ctx *fiber.Ctx) error {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	res, status, err := h.svc.DashboardUser(ctx.UserContext(), limit, page)
	if err != nil {
		return util.HttpError(ctx, http.StatusInternalServerError, err.Error())
	}
	return util.HttpSuccess(ctx, status, res)
}
