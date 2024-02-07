package domain

import (
	"github.com/nafisalfiani/p3-final-project/lib/log"
)

type Domains struct {
}

func Init(logger log.Interface) *Domains {
	return &Domains{}
}
