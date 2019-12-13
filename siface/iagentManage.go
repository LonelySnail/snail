package siface

type IAgentManage interface {
	Add (agent IAgent)
	Remove(id uint32)
	Get(id uint32) (IAgent,bool)
	Clear()
}