package router

import (
	"project/ent"
)

func RegisterRouter(c *ent.Client) {
	registerUserRouter(c)
}
