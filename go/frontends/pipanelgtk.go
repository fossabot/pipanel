package frontends

import (
	pipanel "github.com/BenJetson/pipanel/go"

	"github.com/BenJetson/pipanel/go/frontends/alerters/gtkttsalerter"
	"github.com/BenJetson/pipanel/go/frontends/audio_players/beeper"
	"github.com/BenJetson/pipanel/go/frontends/display_managers/pitouch"
	"github.com/BenJetson/pipanel/go/frontends/power_managers/systemdpwr"
)

// The line below will cause a Makefile target, go build tag, and main frontend
// factory to be generated for this type of frontend.
// go:generate go run gen.go PiPanelGTK

// NewPiPanelGTK creates a pipanel.Frontend that supports the RPi official
// touch display, includes GTK/TTS alerts, and systemd power management.
func NewPiPanelGTK() *pipanel.Frontend {
	return &pipanel.Frontend{
		Alerter:        gtkttsalerter.New(),
		AudioPlayer:    beeper.New(),
		DisplayManager: pitouch.New(),
		PowerManager:   systemdpwr.New(),
	}
}
