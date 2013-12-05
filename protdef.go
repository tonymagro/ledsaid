package main

import (
	"github.com/tonymagro/asign"
)

type protocolDef struct {
	Escape            map[string]byte
	TypeCode          map[string]byte
	CommandCode       map[string]byte
	DisplayPosition   map[string]byte
	ModeCode          map[string]byte
	SpecialMode       map[string]byte
	SpecialGraphics   map[string]byte
	Color             map[string]byte
	ValidLabel        map[string]byte
	ExtendedCharacter map[string]byte
}

var protocol = protocolDef{
	asign.Escape,
	asign.TypeCode,
	asign.CommandCode,
	asign.DisplayPosition,
	asign.ModeCode,
	asign.SpecialMode,
	asign.SpecialGraphics,
	asign.Color,
	asign.ValidLabel,
	asign.ExtendedCharacter,
}
