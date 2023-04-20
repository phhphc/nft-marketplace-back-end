package models

type EnumEvent string

const (
	EventNewCollection EnumEvent = "EventNewCollection"
	EventNewErc721     EnumEvent = "EventNewErc721"
)
