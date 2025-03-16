package structs

import (
  "fmt"
  "time"
  "sync"

  "github.com/hashicorp/raft"
)

type Updater interface {
  Updater() 
  Result()
}

type PasswordEntry struct {
  ID          string    `json:"id"`
  Username    string    `json:"username"`
  TargetType  string    `json:"target_type"` // "api", "ssh", etc // convert to interface?
  TargetAddr  string    `json:"target_addr"` // move from string to IP.Addr?
  LastModified time.time `json:"last_modified"`
  Version     int       `json:"version"`
}

type Store struct {
  mu        sync.RWMutex
  Entries   map[string]PasswordEntry `json:"entries"`
  Version   int                      `json:"entry"`
  EncKey    []byte                   `json:"-"`
  dataFile  string
}

type Command struct {
  Type    string           `json:"type"`
  Entry   PasswordEntry    `json:"entry"`
  Applied bool             `json:"applied`
}

type ClusterNode struct {
  Store       *Store
  raft        *raft.Raft
  nodeName    string
  nodeAddr    string //IP.Addr?
  dataDir     string
  isLeader    bool
  updateChan  chan Command
}

type UpdateResult struct { 
  // The way we check whether or not a change was successful depends on the Target_Type and vendors. Decorators?
  Success  bool
  Message  string
  Error    error
}


