package override

var ExposePropertiesInterface = map[string]bool{
	"org.bluez.AgentManager1": false,
}

// ExposeProperties expose Properties interface to the struct
func ExposeProperties(iface string) bool {
	if val, ok := ExposePropertiesInterface[iface]; ok {
		return val
	}
	return true
}
