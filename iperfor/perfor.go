package iperfor

import (
	"fmt"
	"strings"
	"sync"
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
func GetDot(name string) string {
	return defaultPerfor.GetDot(name)
}
func Print() string {
	return defaultPerfor.Print()
}

type Perfor struct {
	nameList []string
	timeMap  map[string]*timePair
	lock     sync.Mutex
}

func (p *Perfor) DotBegin(name string) {
	p.lock.Lock()
	defer p.lock.Unlock()
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

func (p *Perfor) GetDot(name string) string {
	strList := make([]string, 0, 4)
	strList = append(strList, fmt.Sprintf("---------------%v---------------", name))
	strList = append(strList, fmt.Sprintf("[性能测试][%v][开始时间]%v", name, p.timeMap[name].Begin))
	strList = append(strList, fmt.Sprintf("[性能测试][%v][结束时间]%v", name, p.timeMap[name].End))
	duration := p.timeMap[name].End.Sub(p.timeMap[name].Begin)
	durationFormat := "%v"
	if duration > time.Millisecond*500 {
		durationFormat = "\033[1;31m%v\033[0m"
	}
	strList = append(strList, fmt.Sprintf("[性能测试][%v][耗时]%v", name, fmt.Sprintf(durationFormat, duration)))
	return strings.Join(strList, "\n")
}

func (p *Perfor) Print() string {
	strList := make([]string, 0, len(p.nameList))
	for _, name := range p.nameList {
		strList = append(strList, p.GetDot(name))
	}
	return strings.Join(strList, "\n")
}

func NewPerfor() *Perfor {
	return &Perfor{nameList: make([]string, 0), timeMap: make(map[string]*timePair)}
}

type timePair struct {
	Name  string
	Begin time.Time
	End   time.Time
}
