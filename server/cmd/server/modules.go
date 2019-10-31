//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/witchcraft1/Architecture-course-2/server/dormitories"
)

func ComposeApiServer(port HttpPortNumber) (*DormitoryApiServer, error) {
	wire.Build(
		NewDbConnection,
		dormitories.Providers,
		wire.Struct(new(DormitoryApiServer), "Port", "ChannelsHandler"),
	)
	return nil, nil
}
