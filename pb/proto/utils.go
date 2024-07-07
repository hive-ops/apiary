package pb

import (
	"fmt"
)

func NewNamespace(org, collection string) *Namespace {
	return &Namespace{
		Org:        org,
		Collection: collection,
	}
}

func NewKeyspace(namespace *Namespace, name string) *Keyspace {
	return &Keyspace{
		Namespace: namespace,
		Name:      name,
	}
}

func NewSetEntriesCommand(keyspace *Keyspace, entries []*Entry) *SetEntriesCommand {
	return &SetEntriesCommand{
		Keyspace: keyspace,
		Entries:  entries,
	}
}

func NewGetEntriesCommand(keyspace *Keyspace, keys []string) *GetEntriesCommand {
	return &GetEntriesCommand{
		Keyspace: keyspace,
		Keys:     keys,
	}
}

func NewDeleteEntriesCommand(keyspace *Keyspace, keys []string) *DeleteEntriesCommand {
	return &DeleteEntriesCommand{
		Keyspace: keyspace,
		Keys:     keys,
	}
}

func (k *Keyspace) GetKeyspaceRef() string {
	return fmt.Sprintf("%s__%s__%s", k.Namespace.Org, k.Namespace.Collection, k.Name)
}
