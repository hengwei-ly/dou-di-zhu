package controllers

import (
	"doudizhu/app"

	"github.com/revel/revel"
)

type Rooms struct {
	App
}

func (c Rooms) Index(roomId int) revel.Result {
	c.ViewArgs["tables"] = app.Tables
	return c.Render()
}
