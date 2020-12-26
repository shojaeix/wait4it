package model

import (
	"golang.org/x/net/context"
)

// CheckInterface is an interface to handle single check
type CheckInterface interface {
	BuildContext(cx CheckContext)
	Validate() error
	Check(ctx context.Context) (bool, bool, error)
}
