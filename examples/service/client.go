package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile"
	"github.com/muka/go-bluetooth/emitter"
	"github.com/muka/go-bluetooth/linux/btmgmt"
	log "github.com/sirupsen/logrus"
)

func createClient(adapterID, hwaddr, serviceID string) (err error) {

	log.Infof("Discovering devices from %s, looking for %s (serviceID:%s)", serviceAdapterID, hwaddr, serviceID)

	// serviceAdapterInfo, err := hcitool.GetAdapter(serviceAdapterID)
	// if err != nil {
	// 	return err
	// }

	btmgmt := btmgmt.NewBtMgmt(adapterID)

	// turn off/on
	err = btmgmt.Reset()
	if err != nil {
		return err
	}

	adapter := profile.NewAdapter1(clientAdapterID)
	err = adapter.StartDiscovery()
	if err != nil {
		log.Errorf("Failed to start discovery: %s", err.Error())
		return err
	}

	devices, err := api.GetDevices()
	if err != nil {
		return err
	}

	for _, d := range devices {
		err = adapter.RemoveDevice(d.Path)
		if err != nil {
			log.Warnf("Cannot remove %s : %s", d.Path, err.Error())
		}
	}

	err = api.On("discovery", emitter.NewCallback(func(ev emitter.Event) {

		discoveryEvent := ev.GetData().(api.DiscoveredDeviceEvent)
		if discoveryEvent.Status == api.DeviceAdded {
			err := showDeviceInfo(discoveryEvent.Device, hwaddr, serviceID)
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
		}

	}))

	return err
}

func showDeviceInfo(dev *api.Device, hwaddr, serviceID string) error {

	if dev == nil {
		return errors.New("Device is nil")
	}

	props, err := dev.GetProperties()
	if err != nil {
		return fmt.Errorf("%s: Failed to get properties: %s", dev.Path, err.Error())
	}

	if hwaddr != props.Address {
		return nil
	}

	serviceID = strings.ToLower(serviceID)

	log.Infof("Found device name=%s addr=%s rssi=%d", props.Name, props.Address, props.RSSI)

	var found bool
	for _, uuid := range props.UUIDs {
		if strings.ToLower(uuid) == serviceID {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Service UUID %s not found on %s", serviceID, hwaddr)
	}

	// log.Debugf("Connecting to %s...", props.Name)
	// err = dev.Connect()
	// if err != nil {
	// 	return fmt.Errorf("Connect: %s", err)
	// }
	//
	// log.Debugf("Pairing to %s...", props.Name)
	// err = dev.Pair()
	// if err != nil {
	// 	return fmt.Errorf("Pair: %s", err)
	// }

	log.Infof("Found UUID %s", serviceID)

	chars, err := dev.GetCharsList()
	if err != nil {
		return fmt.Errorf("Service UUID %s not found on %s", serviceID, hwaddr)
	}

	log.Infof("CHARS %++v", chars)

	return nil
}
