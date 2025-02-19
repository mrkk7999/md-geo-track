package controller

import "net/http"

func (c *Controller) HeartBeatHandler(w http.ResponseWriter, r *http.Request) {
	response := c.service.HeartBeat()
	encodeJSONResponse(w, http.StatusOK, response, nil)
}
