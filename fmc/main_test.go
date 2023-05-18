package fmc

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	FMC_HOST := os.Getenv("FMC_HOST")
	FMC_USERNAME := os.Getenv("FMC_USERNAME")
	FMC_PASSWORD := os.Getenv("FMC_PASSWORD")
	FMC_INSECURE_SKIP_VERIFY := os.Getenv("FMC_INSECURE_SKIP_VERIFY") == "true" // convert string to bool

	c := NewClient(FMC_HOST, FMC_USERNAME, FMC_PASSWORD, FMC_INSECURE_SKIP_VERIFY)
	err := c.Login()
	if c.Login(); err != nil {
		panic("Unable to login")
	}

	err = setup(c)
	if err != nil {
		panic("Unable to register device")
	}

	m.Run()

	err = teardown(c)
	if err != nil {
		panic("Unable to de-register device")
	}
}

// Register device
func setup(c *Client) error {
	ctx := context.Background()

	// Create access policy
	defaultAction := AccessPolicyDefaultAction{
		Type:            access_policy_default_action_type,
		Logend:          true,
		Sendeventstofmc: true,
		Action:          "BLOCK",
	}

	res, err := c.CreateFmcAccessPolicy(ctx, &AccessPolicy{
		Name:          "Test-ACP",
		Type:          "AccessPolicy",
		Defaultaction: defaultAction,
	})

	if err != nil {
		return fmt.Errorf("error creating access policy: %s", err.Error())
	}

	// Register device
	FTD_HOST := os.Getenv("FTD_HOST")
	_, err = c.CreateFmcDevice(ctx, &Device{
		Name:         "ftd.adyah.cisco",
		HostName:     FTD_HOST,
		RegKey:       "cisco",
		Type:         "Device",
		AccessPolicy: &AccessPolicyItem{ID: res.ID, Type: res.Type},
	})

	if err != nil {
		return fmt.Errorf("error registering device: %s", err.Error())
	}

	return nil
}

// De-Register device
func teardown(c *Client) error {
	ctx := context.Background()

	// Delete Device
	device, err := c.GetFmcDeviceByName(ctx, "ftd.adyah.cisco")
	if err != nil {
		return fmt.Errorf("error getting device: %s", err.Error())
	}

	err = c.DeleteFmcDevice(ctx, device.ID)
	if err != nil {
		return fmt.Errorf("error deleting device: %s", err.Error())
	}

	// Delete Access Policy
	acp, err := c.GetFmcAccessPolicyByName(ctx, "Test-ACP")
	if err != nil {
		return fmt.Errorf("error getting access policy: %s", err.Error())
	}

	err = c.DeleteFmcAccessPolicy(ctx, acp.ID)
	if err != nil {
		return fmt.Errorf("error deleting access policy: %s", err.Error())
	}

	return nil
}
