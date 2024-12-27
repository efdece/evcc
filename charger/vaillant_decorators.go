package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateVaillant(base *Vaillant, battery func() (float64, error)) api.Charger {
	switch {
	case battery == nil:
		return base

	case battery != nil:
		return &struct {
			*Vaillant
			api.Battery
		}{
			Vaillant: base,
			Battery: &decorateVaillantBatteryImpl{
				battery: battery,
			},
		}
	}

	return nil
}

type decorateVaillantBatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateVaillantBatteryImpl) Soc() (float64, error) {
	return impl.battery()
}
