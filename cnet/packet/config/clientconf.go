package config

import (
	"bytes"
	"io"

	"github.com/sleiderr/gomc/ctypes"
)

type ClientChatMode byte
type ClientMainHand byte

const (
	ChatModeEnabled      ClientChatMode = 0
	ChatModeCommandsOnly ClientChatMode = 1
	ChatModeHidden       ClientChatMode = 2
)

const (
	MainHandLeft  ClientMainHand = 0
	MainHandRight ClientMainHand = 1
)

const (
	CapeEnabled = 1 << iota
	JacketEnabled
	LeftSleeveEnabled
	RightSleeveEnabled
	LeftPantsEnabled
	RightPantsEnabled
	HatEnabled
)

type ClientInformation struct {
	Locale            string
	ViewDistance      byte
	ChatMode          ClientChatMode
	ChatColors        bool
	DisplayedSkinPart byte
	MainHand          ClientMainHand
	TextFiltering     bool
	ServerListings    bool
}

type PluginMessage struct {
	Identifier string
	Data       []byte
}

func (p *ClientInformation) Raw() []byte {
	rawData := new(bytes.Buffer)

	ctypes.WriteString(rawData, p.Locale)
	rawData.WriteByte(p.ViewDistance)
	rawData.WriteByte(byte(p.ChatMode))
	if p.ChatColors {
		rawData.WriteByte(1)
	} else {
		rawData.WriteByte(0)
	}
	rawData.WriteByte(p.DisplayedSkinPart)
	rawData.WriteByte(byte(p.MainHand))
	if p.TextFiltering {
		rawData.WriteByte(1)
	} else {
		rawData.WriteByte(0)
	}
	if p.ServerListings {
		rawData.WriteByte(1)
	} else {
		rawData.WriteByte(0)
	}

	return rawData.Bytes()
}

func (p *ClientInformation) FillFromRaw(raw []byte) error {
	r := bytes.NewReader(raw)

	locale, err := ctypes.ReadString(r)
	viewDistance, err := r.ReadByte()
	chatMode, err := r.ReadByte()
	chatColor, err := r.ReadByte()
	displayedSkinPart, err := r.ReadByte()
	mainHand, err := r.ReadByte()
	textFiltering, err := r.ReadByte()
	serverListings, err := r.ReadByte()

	if err != nil {
		return err
	}

	p.Locale = locale
	p.ViewDistance = viewDistance
	p.ChatMode = ClientChatMode(chatMode)
	p.ChatColors = chatColor == 1
	p.DisplayedSkinPart = displayedSkinPart
	p.MainHand = ClientMainHand(mainHand)
	p.TextFiltering = textFiltering == 1
	p.ServerListings = serverListings == 1

	return nil
}

func (p *PluginMessage) Raw() []byte {
	rawData := new(bytes.Buffer)

	ctypes.WriteString(rawData, p.Identifier)
	rawData.Write(p.Data)

	return rawData.Bytes()
}

func (p *PluginMessage) FillFromRaw(raw []byte) error {
	r := bytes.NewReader(raw)

	identifier, err := ctypes.ReadString(r)
	data, err := io.ReadAll(r)

	if err != nil {
		return err
	}

	p.Identifier = identifier
	p.Data = data

	return nil
}
