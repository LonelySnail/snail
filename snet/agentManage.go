package snet

import (
	"github.com/snail/siface"
	"sync"
)

type agentManage struct {
	agents    map[uint32]siface.IAgent
	lock      sync.RWMutex
}

func NewAgentManage()  siface.IAgentManage {
	return &agentManage{
		agents: map[uint32]siface.IAgent{},
	}

}

func (am *agentManage)Add(agent siface.IAgent)  {
	am.lock.Lock()
	am.agents[agent.GetAgentId()]= agent
	am.lock.Unlock()
}

func (am *agentManage)Remove(id uint32)  {
	am.lock.Lock()
	delete(am.agents,id)
	am.lock.Unlock()
}

func (am *agentManage)Get(id uint32) (siface.IAgent,bool) {
	am.lock.RLock()
	agent,ok :=  am.agents[id]
	am.lock.RUnlock()
	return agent,ok
}

func (am *agentManage)Clear()  {
	am.agents = map[uint32]siface.IAgent{}
}