// Code generated DO NOT EDIT

package mesh



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus"
)

var Attention1Interface = "org.bluez.mesh.Attention1"


// NewAttention1 create a new instance of Attention1
//
// Args:
// - servicePath: unique name
// - objectPath: freely definable
func NewAttention1(servicePath string, objectPath dbus.ObjectPath) (*Attention1, error) {
	a := new(Attention1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  servicePath,
			Iface: Attention1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(Attention1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
Attention1 Mesh Attention Hierarchy

*/
type Attention1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*Attention1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// Attention1Properties contains the exposed properties of an interface
type Attention1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

}

//Lock access to properties
func (p *Attention1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *Attention1Properties) Unlock() {
	p.lock.Unlock()
}



// Close the connection
func (a *Attention1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return Attention1 object path
func (a *Attention1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return Attention1 dbus client
func (a *Attention1) Client() *bluez.Client {
	return a.client
}

// Interface return Attention1 interface
func (a *Attention1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *Attention1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a Attention1Properties to map
func (a *Attention1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an Attention1Properties
func (a *Attention1Properties) FromMap(props map[string]interface{}) (*Attention1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an Attention1Properties
func (a *Attention1Properties) FromDBusMap(props map[string]dbus.Variant) (*Attention1Properties, error) {
	s := new(Attention1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *Attention1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *Attention1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *Attention1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *Attention1) GetProperties() (*Attention1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *Attention1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *Attention1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *Attention1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *Attention1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *Attention1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *Attention1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
SetTimer 
		The element_index parameter is the element's index within the
		node where the health server model is hosted.

		The time parameter indicates how many seconds the attention
		state shall be on.

		PossibleErrors:
			org.bluez.mesh.Error.NotSupported


*/
func (a *Attention1) SetTimer(element_index uint8, time uint16) error {
	
	return a.client.Call("SetTimer", 0, element_index, time).Store()
	
}

/*
GetTimer 
		The element parameter is the unicast address within the node
		where the health server model is hosted.

		Returns the number of seconds for how long the attention action
		remains staying on.

		PossibleErrors:
			org.bluez.mesh.Error.NotSupported



*/
func (a *Attention1) GetTimer(element uint16) (uint16, error) {
	
	var val0 uint16
	err := a.client.Call("GetTimer", 0, element).Store(&val0)
	return val0, err	
}

