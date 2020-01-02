package frontends

import (
	"log"

	pipanel "github.com/BenJetson/pipanel/go"

	"github.com/BenJetson/pipanel/go/frontends/alerters/gtkttsalerter"
	"github.com/BenJetson/pipanel/go/frontends/audio_players/beeper"
	"github.com/BenJetson/pipanel/go/frontends/display_managers/pitouch"
	"github.com/BenJetson/pipanel/go/frontends/power_managers/systemdpwr"
)

// NewPiPanelGTK creates a pipanel.Frontend that supports the RPi official
// touch display, includes GTK/TTS alerts, and systemd power management.
func NewPiPanelGTK(log *log.Logger) *pipanel.Frontend {
	return &pipanel.Frontend{
		Alerter:        gtkttsalerter.New(log),
		AudioPlayer:    beeper.New(log),
		DisplayManager: pitouch.New(log),
		PowerManager:   systemdpwr.New(log),
	}
}
