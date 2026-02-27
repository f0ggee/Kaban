package serverManagment

import (
	"log/slog"
	"time"
)

func (s *ServerManagement) GetServerName(i int) string {

	slog.Time("Start searching for server name", time.Now())

	switch i {

	case 1:
		slog.Info("Found the server name", "ServerNumber", i)
		return "server1"

	case 2:
		slog.Info("Found the server name", "ServerNumber", i)
		return "server2"

	}

	slog.Info("Couldn't find the server name", "ServerNumber", i)
	return ""
}
