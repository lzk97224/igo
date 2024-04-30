package iperfor

import (
	log "github.com/sirupsen/logrus"
	"time"
)

var defaultPerfor = NewPerfor()

func InitPerfor() {
	defaultPerfor = NewPerfor()
}

func DotBegin(name string) {
	defaultPerfor.DotBegin(name)
}
func DotEnd(name string) {
	defaultPerfor.DotEnd(name)
}
func Print() {
	defaultPerfor.Print()
}

type Perfor struct {
	nameList []string
	timeMap  map[string]*timePair
}

func (p *Perfor) DotBegin(name string) {
	if pair, ok := p.timeMap[name]; ok {
		pair.Begin = time.Now()
	} else {
		p.nameList = append(p.nameList, name)
		p.timeMap[name] = &timePair{
			Name:  name,
			Begin: time.Now(),
		}
	}
}

func (p *Perfor) DotEnd(name string) {
	if pair, ok := p.timeMap[name]; ok {
		pair.End = time.Now()
	} else {
		p.nameList = append(p.nameList, name)
		p.timeMap[name] = &timePair{
			Name: name,
			End:  time.Now(),
		}
	}
}

func (p *Perfor) Print() {
	for _, name := range p.nameList {
		log.Infof("---------------%v---------------", name)
		log.Infof("[性能测试][%v][开始时间]%v", name, p.timeMap[name].Begin)
		log.Infof("[性能测试][%v][结束时间]%v", name, p.timeMap[name].End)
		log.Infof("[性能测试][%v][耗时]%v", name, p.timeMap[name].End.Sub(p.timeMap[name].Begin))
	}
}

func NewPerfor() *Perfor {
	return &Perfor{nameList: make([]string, 0), timeMap: make(map[string]*timePair)}
}

type timePair struct {
	Name  string
	Begin time.Time
	End   time.Time
}
