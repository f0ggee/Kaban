package Uttiltesss

import (
	"context"
	"time"
)

func Contexte() (ctx context.Context, cancel context.CancelFunc) {

	return context.WithTimeout(context.Background(), 5*time.Second)

}
