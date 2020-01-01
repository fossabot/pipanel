package pitouch

import (
	"errors"
	"log"
	"os"
	"strconv"

	pipanel "github.com/BenJetson/pipanel/go"
)

const brightFile string = "/sys/class/backlight/rpi_backlight/brightness"

// TouchDisplayManager implements pipanel.DisplayManager for the Raspberry Pi
// official 7" touchscreen device.
type TouchDisplayManager struct {
	log *log.Logger
}

// New creates a TouchDisplayManager instance.
func New(log *log.Logger) *TouchDisplayManager {
	return &TouchDisplayManager{log}
}

// SetBrightness handles pipanel brightness events.
func (t *TouchDisplayManager) SetBrightness(e pipanel.BrightnessEvent) error {
	// Setting the brightness less than ten will cause the screen to blank.
	if e.Level < 10 {
		return errors.New("device does not support brigtness values < 10")
	}

	f, err := os.OpenFile(brightFile, os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	f.WriteString(strconv.Itoa(int(e.Level)))
	t.log.Printf("Setting RPi touchscreen brightness to %d.", e.Level)

	return f.Close()
}
