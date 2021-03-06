package pipanel

import (
	"log"
)

// Frontend provides abstraction for processing PiPanel events.
type Frontend struct {
	Alerter
	AudioPlayer
	PowerManager
	DisplayManager
}

// Init initializes all components of the Frontend.
func (f *Frontend) Init(log *log.Logger, cfg *FrontendConfig) error {
	if f.Alerter != nil {
		if err := f.Alerter.Init(log, cfg.AlerterConfig); err != nil {
			return err
		}
	}

	if f.AudioPlayer != nil {
		if err := f.AudioPlayer.Init(log, cfg.AudioPlayerConfig); err != nil {
			return err
		}
	}

	if f.PowerManager != nil {
		if err := f.PowerManager.Init(log, cfg.PowerManagerConfig); err != nil {
			return err
		}
	}

	if f.DisplayManager != nil {
		if err := f.DisplayManager.Init(log, cfg.DisplayManagerConfig); err != nil {
			return err
		}
	}

	return nil
}

// Cleanup tears down all components of the Frontend.
func (f *Frontend) Cleanup() error {
	if f.Alerter != nil {
		if err := f.Alerter.Cleanup(); err != nil {
			return err
		}
	}

	if f.AudioPlayer != nil {
		if err := f.AudioPlayer.Cleanup(); err != nil {
			return err
		}
	}

	if f.PowerManager != nil {
		if err := f.PowerManager.Cleanup(); err != nil {
			return err
		}
	}

	if f.DisplayManager != nil {
		if err := f.DisplayManager.Cleanup(); err != nil {
			return err
		}
	}

	return nil
}
