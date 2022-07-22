package router

import (
	"project/ent"
)

func RegisterRouter(c *ent.Client) {
	registerHealthCheckRouter()
	registerUserRouter(c)
}
