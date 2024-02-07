package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
)

// @Summary Trigger Scheduler
// @Description Trigger Scheduler
// @Security BearerAuth
// @Security XDateTimes
// @Tags Scheduler
// @Param trigger_input body entity.TriggerSchedulerParams true "Parameter for triggering scheduler"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 403 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Router /api/v1/admin/scheduler/trigger [POST]
func (r *rest) TriggerScheduler(ctx *gin.Context) {
	triggerParams := entity.TriggerSchedulerParams{}
	if err := r.Bind(ctx, &triggerParams); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.scheduler.TriggerScheduler(triggerParams.Name); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}
