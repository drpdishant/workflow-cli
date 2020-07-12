package cmd

import (
	"github.com/teamhephy/controller-sdk-go/api"
	"github.com/teamhephy/controller-sdk-go/appsettings"
)

// RoutingInfo provides information about the status of app routing.
func (d *HephyCmd) RoutingInfo(appID string) error {
	s, appID, err := load(d.ConfigFile, appID)

	if err != nil {
		return err
	}

	appSettings, err := appsettings.List(s.Client, appID)
	if d.checkAPICompatibility(s.Client, err) != nil {
		return err
	}

	if appSettings.Routable == nil || *appSettings.Routable {
		d.Println("Routing is enabled.")
	} else {
		d.Println("Routing is disabled.")
	}
	return nil
}

// RoutingEnable enables an app from being exposed by the router.
func (d *HephyCmd) RoutingEnable(appID string) error {
	s, appID, err := load(d.ConfigFile, appID)

	if err != nil {
		return err
	}

	d.Printf("Enabling routing for %s... ", appID)

	quit := progress(d.WOut)
	appSettings := api.AppSettings{Routable: api.NewRoutable()}
	_, err = appsettings.Set(s.Client, appID, appSettings)

	quit <- true
	<-quit

	if err != nil {
		return err
	}

	d.Println("done")
	return nil
}

// RoutingDisable disables an app from being exposed by the router.
func (d *HephyCmd) RoutingDisable(appID string) error {
	s, appID, err := load(d.ConfigFile, appID)

	if err != nil {
		return err
	}

	d.Printf("Disabling routing for %s... ", appID)

	quit := progress(d.WOut)
	appSettings := api.AppSettings{Routable: api.NewRoutable()}
	*appSettings.Routable = false
	_, err = appsettings.Set(s.Client, appID, appSettings)

	quit <- true
	<-quit

	if err != nil {
		return err
	}

	d.Println("done")
	return nil
}
