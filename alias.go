// Package core provides convenient access to foundational, central go-libp2p primitives via type aliases.
package core

import (
	"github.com/RTradeLtd/libp2px-core/host"
	"github.com/RTradeLtd/libp2px-core/network"
	"github.com/RTradeLtd/libp2px-core/peer"
	"github.com/RTradeLtd/libp2px-core/protocol"
	"github.com/multiformats/go-multiaddr"
)

// Multiaddr aliases the Multiaddr type from github.com/multiformats/go-multiaddr.
//
// Refer to the docs on that type for more info.
type Multiaddr = multiaddr.Multiaddr

// PeerID aliases peer.ID.
//
// Refer to the docs on that type for more info.
type PeerID = peer.ID

// PeerID aliases protocol.ID.
//
// Refer to the docs on that type for more info.
type ProtocolID = protocol.ID

// PeerAddrInfo aliases peer.AddrInfo.
//
// Refer to the docs on that type for more info.
type PeerAddrInfo = peer.AddrInfo

// Host aliases host.Host.
//
// Refer to the docs on that type for more info.
type Host = host.Host

// Network aliases network.Network.
//
// Refer to the docs on that type for more info.
type Network = network.Network

// Conn aliases network.Conn.
//
// Refer to the docs on that type for more info.
type Conn = network.Conn

// Stream aliases network.Stream.
//
// Refer to the docs on that type for more info.
type Stream = network.Stream
