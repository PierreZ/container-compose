package templating 

import (
	"net"
)

// Context is the context used during templating
type Context struct {

	CurrentGroup string
	CurrentNumber int 
	CurentIP net.IP

	Groups map[string]Group

	Others map[string]string
}

// Group is representing a group of containers
type Group struct {
	Name string
	DockerImage string 

}


