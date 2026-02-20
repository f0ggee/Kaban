package serverManagment

import (
	"log/slog"
	"os"
	"time"
)

type ServerManagement struct{}

func (s ServerManagement) GetServerKey(Num int) string {

	slog.Time("Start searching for server key", time.Now())
	switch Num {
	case 1:

		slog.Info("start getting the server key", "ServerNumber", Num)

		return os.Getenv("server1")

	}

	slog.Info("Couldn't find the server key", "ServerNumber", Num)
	return ""
}
