package daemon

import (
	"math/rand"
	"poc/daemon/client"
)

func (d *Daemon) pickClient() client.INodeClient {
	return d.nodeClients[rand.Intn(len(d.nodeClients))] //take random client
}