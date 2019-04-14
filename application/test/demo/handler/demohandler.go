package handler

import (
	"net/http"

	"github.com/yakaa/cuter/application/test/demo/logic"
	"github.com/yakaa/cuter/common/baseresponse"
	"github.com/yakaa/cuter/lib/httpx"
	"github.com/yakaa/cuter/lib/logx"
)

type DemoHandler struct {
	demoLogic *logic.DemoLogic
}

func NewDemoHandler(demoLogic *logic.DemoLogic) *DemoHandler {
	return &DemoHandler{
		demoLogic: demoLogic,
	}
}

func (h *DemoHandler) Demo(w http.ResponseWriter, r *http.Request) {
	var req logic.DemoRequest
	if err := httpx.Parse(r, &req); err != nil {
		baseresponse.HttpParamError(w, err)
		return
	}
	logx.Infof("%+v", req)
	resp, err := h.demoLogic.Demo(&req)
	baseresponse.FormatResponse(resp, err, w)
}
