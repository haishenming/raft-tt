package lib

import (
	"github.com/hashicorp/raft"
	"io"
)

type MyFSM struct {}

func (m MyFSM) Apply(log *raft.Log) interface{} {
	panic("implement me")
}

func (m MyFSM) Snapshot() (raft.FSMSnapshot, error) {
	panic("implement me")
}

func (m MyFSM) Restore(closer io.ReadCloser) error {
	panic("implement me")
}

