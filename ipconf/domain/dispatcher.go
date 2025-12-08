package domain

import (
	"github.com/Nuyoahch/conflux/ipconf/source"
	"sort"
	"sync"
)

// Dispatcher 调度模块
type Dispatcher struct {
	candidateTable map[string]*Endport
	sync.RWMutex
}

var dp *Dispatcher

// Init 初始化
func Init() {
	dp = &Dispatcher{}
	dp.candidateTable = make(map[string]*Endport)
	go func() {
		for event := range source.EventChan() {
			switch event.Type {
			case source.AddNodeEvent:
				dp.addNode(event)
			case source.DelNodeEvent:
				dp.delNode(event)
			}
		}
	}()
}

// Dispatch 实现调度方法
func Dispatch(ctx *IpConfContext) []*Endport {
	// Step1: 获得候选endport
	eds := dp.getCandidateEndport(ctx)
	// Step2: 逐一计算得分
	for _, ed := range eds {
		ed.CalculateScore(ctx)
	}
	// Step3: 全局排序，动静结合的排序策略。
	sort.Slice(eds, func(i, j int) bool {
		// 优先基于活跃分数进行排序
		if eds[i].ActiveSorce > eds[j].ActiveSorce {
			return true
		}
		// 如果活跃分数相同，则使用静态分数排序
		if eds[i].ActiveSorce == eds[j].ActiveSorce {
			if eds[i].StaticSorce > eds[j].StaticSorce {
				return true
			}
			return false
		}
		return false
	})
	return eds
}

// 获取候选端口号
func (dp *Dispatcher) getCandidateEndport(ctx *IpConfContext) []*Endport {
	dp.RLock()
	defer dp.RUnlock()
	candidateList := make([]*Endport, 0, len(dp.candidateTable))
	for _, ed := range dp.candidateTable {
		candidateList = append(candidateList, ed)
	}
	return candidateList
}

// 删除 node 节点
func (dp *Dispatcher) delNode(event *source.Event) {
	dp.Lock()
	defer dp.Unlock()
	delete(dp.candidateTable, event.Key())
}

// 添加 node 节点
func (dp *Dispatcher) addNode(event *source.Event) {
	dp.Lock()
	defer dp.Unlock()
	ed := NewEndport(event.IP, event.Port)
	ed.UpdateStat(&Stat{
		ConnectNum:   event.ConnectNum,
		MessageBytes: event.MessageBytes,
	})
	dp.candidateTable[event.Key()] = ed
}
