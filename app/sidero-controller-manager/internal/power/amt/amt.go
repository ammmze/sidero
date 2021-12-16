package amt

import (
	amt "github.com/ammmze/go-amt"
	metalv1alpha1 "github.com/talos-systems/sidero/app/sidero-controller-manager/api/v1alpha1"
	"github.com/talos-systems/sidero/app/sidero-controller-manager/internal/power/metal"
)

type Client struct {
	AMTClient *amt.Client
}

// NewClient creates an amt client to use.
func NewClient(amtInfo metalv1alpha1.AMT) (*Client, error) {
	connection := &amt.Connection{
		Host: amtInfo.Endpoint,
		Port: amtInfo.Port,
		User: amtInfo.User,
		Pass: amtInfo.Pass,
	}
	amtClient, err := amt.NewClient(*connection)
	if err != nil {
		return nil, err
	}
	return &Client{
		AMTClient: amtClient,
	}, nil
}

// Close the client.
func (c *Client) Close() error {
	return nil
}

// PowerOn will power on a given machine.
func (c *Client) PowerOn() error {
	return c.AMTClient.PowerOn()
}

// PowerOff will power off a given machine.
func (c *Client) PowerOff() error {
	return c.AMTClient.PowerOff()
}

// PowerCycle will power cycle a given machine.
func (c *Client) PowerCycle() error {
	return c.AMTClient.PowerCycle()
}

// SetPXE makes sure the node will pxe boot next time.
func (c *Client) SetPXE(mode metal.PXEMode) error {
	// TODO: not sure how/if we can use the pxe mode
	return c.AMTClient.SetPXE()
}

// IsPoweredOn checks current power state.
func (c *Client) IsPoweredOn() (bool, error) {
	return c.AMTClient.IsPoweredOn()
}

// IsFake returns false.
func (c *Client) IsFake() bool {
	return false
}
