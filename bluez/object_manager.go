package bluez

import (
	"github.com/godbus/dbus"
)

// NewObjectManager create a new ObjectManager client
func NewObjectManager(name string, path string) (*ObjectManager, error) {
	om := new(ObjectManager)
	om.client = NewClient(
		&Config{
			Name:  name,
			Iface: "org.freedesktop.DBus.ObjectManager",
			Path:  path,
			Bus:   SystemBus,
		},
	)
	return om, nil
}

// ObjectManager manges the list of all available objects
type ObjectManager struct {
	client *Client
}

// Close the connection
func (o *ObjectManager) Close() {
	o.client.Disconnect()
}

// GetManagedObjects return a list of all available objects registered
func (o *ObjectManager) GetManagedObjects() (map[dbus.ObjectPath]map[string]map[string]dbus.Variant, error) {
	var objs map[dbus.ObjectPath]map[string]map[string]dbus.Variant
	err := o.client.Call("GetManagedObjects", 0).Store(&objs)
	return objs, err
}

//Register watch for signal events
func (o *ObjectManager) Register() (chan *dbus.Signal, error) {
	path := o.client.Config.Path
	iface := o.client.Config.Iface
	return o.client.Register(path, iface)
}

//Unregister watch for signal events
func (o *ObjectManager) Unregister(signal chan *dbus.Signal) error {
	path := o.client.Config.Path
	iface := o.client.Config.Iface
	return o.client.Unregister(path, iface, signal)
}

// SignalAdded notify of interfaces being added
func (o *ObjectManager) SignalAdded(path dbus.ObjectPath, props map[string]map[string]dbus.Variant) error {
	return o.client.Emit(path, InterfacesAdded, props)
}

// SignalRemoved notify of interfaces being removed
func (o *ObjectManager) SignalRemoved(path dbus.ObjectPath, ifaces []string) error {
	if ifaces == nil {
		ifaces = make([]string, 0)
	}
	return o.client.Emit(path, InterfacesRemoved, ifaces)
}

// GetManagedObject return an up to date view of a single object state.
// object is nil if the object path is not found
func (o *ObjectManager) GetManagedObject(objpath dbus.ObjectPath) (map[string]map[string]dbus.Variant, error) {
	objects, err := o.GetManagedObjects()
	if err != nil {
		return nil, err
	}
	if p, ok := objects[objpath]; ok {
		return p, nil
	}
	// return nil, errors.New("Object not found")
	return nil, nil
}