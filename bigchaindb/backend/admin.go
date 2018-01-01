package backend

import (

)

type GetConfiger interface {
	GetConfig(connection , interface{}, table)
}

type Reconfigurer interface {
	Reconfigure(connection, interface{}, table, shareds, replicas, kwargs)
}

type SetSharedser interface {
	SetShareds(connection, interface{}, shareds)
}

func SetReplicaser interface {
	SetReplicas(connection, interface{}, replicas)
}

func AddReplicaser interface {
	AddReplica(connection, replicas)
}

func RemoveReplicaser interface {
	Removereplicas(connection, replicas)
}