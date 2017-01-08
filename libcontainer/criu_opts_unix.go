// +build linux freebsd

package libcontainer

// cgroup restoring strategy provided by criu
type cgMode uint32

// CRIU cgroup restoring modes.
const (
	CriuCgModeSoft    cgMode = 3 + iota // restore cgroup properties if only dir created by criu
	CriuCgModeFull                      // always restore all cgroups and their properties
	CriuCgModeStrict                    // restore all, requiring them to not present in the system
	CriuCgModeDefault                   // the same as CriuCgModeSoft
)

// CriuPageServerInfo contains information of CRIU page server.
type CriuPageServerInfo struct {
	Address string // IP address of CRIU page server
	Port    int32  // port number of CRIU page server
}

// VethPairName contains veth pair names on host and container.
type VethPairName struct {
	ContainerInterfaceName string
	HostInterfaceName      string
}

// CriuOpts contains information for CRIU operations.
type CriuOpts struct {
	ImagesDirectory         string             // directory for storing image files
	WorkDirectory           string             // directory to cd and write logs/pidfiles/stats to
	LeaveRunning            bool               // leave container in running state after checkpoint
	TcpEstablished          bool               // checkpoint/restore established TCP connections
	ExternalUnixConnections bool               // allow external unix connections
	ShellJob                bool               // allow to dump and restore shell jobs
	FileLocks               bool               // handle file locks, for safety
	PageServer              CriuPageServerInfo // allow to dump to criu page server
	VethPairs               []VethPairName     // pass the veth to criu when restore
	ManageCgroupsMode       cgMode             // dump or restore cgroup mode
	EmptyNs                 uint32             // don't c/r properties for namespace from this mask
}
