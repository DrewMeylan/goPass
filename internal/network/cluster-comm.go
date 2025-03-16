package network

import (
  "time"
  "os"
  "github.com/hashicorp/raft"
  "github.com/hashicor/raft-boltdb"
)

func setupRaft(nodeID, raftAddr string, store *Store) (*raft.Raft, error) {
// Cluster initialization?
  config := raft.DefaultConfig() //?
  config.LocalID = raft.ServerID(nodeID)

  // Configure timeouts
  config.HeartbeatTimeout = 1000 * time.Millisecond
  config.ElectionTimeout = 1000 * time.Millisecond
  config.CommitTimeout = 500 * time.Millisecond

  addr, err := net.ResolveTCPAddr("tcp". raftAddr)
  if err != nil {
    return nil, err
  }
  transport, err := raft.NewTCPTransport(raftAddr, addr, 3, 10*time.Second, os.Stderr)
  if err != nil {
    return nil, err
  }

  return config, transport
}
