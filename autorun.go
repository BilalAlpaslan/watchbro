package main

import (
	"errors"
	"os"

	"golang.org/x/sys/windows/registry"
)

const (
	AutoRunName = `WatchBro`
	regKey      = `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
)

func self() string {
	return os.Args[0]
}

func AutoRunEnabled() (bool, error) {
	rk, err := registry.OpenKey(registry.CURRENT_USER, regKey, registry.QUERY_VALUE)
	if err != nil {
		return false, err
	}
	defer rk.Close()

	v, _, err := rk.GetStringValue(AutoRunName)
	if errors.Is(err, registry.ErrNotExist) || errors.Is(err, registry.ErrUnexpectedType) {
		return false, nil
	}
	return v == self(), err
}

func AutoRunDisable() error {
	rk, err := registry.OpenKey(registry.CURRENT_USER, regKey, registry.WRITE)
	if err != nil {
		return err
	}
	defer rk.Close()

	err = rk.DeleteValue(AutoRunName)
	if errors.Is(err, registry.ErrNotExist) {
		return nil
	}
	return err
}

func AutoRunEnable() error {
	rk, err := registry.OpenKey(registry.CURRENT_USER, regKey, registry.WRITE)
	if err != nil {
		return err
	}
	defer rk.Close()
	return rk.SetStringValue(AutoRunName, self())
}
