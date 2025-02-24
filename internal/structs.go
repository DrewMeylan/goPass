// (Probably) a temporary location for storing structs; eventually move them to the package where the structs are consumed

package internals

import {
	"net"
}

type CPDU struct { //Cluster Protocol Data Unit --> Used in cluster bootstrapping, state convergence
	ClusterID	string
	LocalIP		net.IPAddr
	SeedIPs		[]net.IPAddr
	LocalRevision	int
	LocalHash	string
	Priority	int
	MasterState	string
}

type KeepAlive struct {
	
}
