package aidoctor

import (
	"context"
)

type AIDoctorService struct {
	ctx context.Context
}

// NewAIDoctorService new AIDoctorService
func NewAIDoctorService(ctx context.Context) *AIDoctorService {
	return &AIDoctorService{ctx: ctx}
}
