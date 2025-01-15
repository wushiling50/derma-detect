package aidoctor

import (
	"derma/detect/biz/model/api"
	aireq "derma/detect/pkg/ai_req"
)

func (s *AIDoctorService) AIDoctor(req api.AiDoctorRequest) (string, error) {
	return aireq.DoctorReq(req.Question)
}
