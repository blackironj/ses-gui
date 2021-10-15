package repo

import (
	"sync"
)

type repository struct {
	templateNameList []string
	rwMu             sync.RWMutex
}

var (
	once sync.Once
	r    *repository
)

func Instance() *repository {
	once.Do(func() {
		r = &repository{
			templateNameList: make([]string, 0, 30),
		}
	})
	return r
}

func (thiz *repository) Append(newTemplateName string) {
	thiz.rwMu.Lock()
	defer thiz.rwMu.Unlock()
	thiz.templateNameList = append(thiz.templateNameList, newTemplateName)
}

func (thiz *repository) Delete(idx int) {
	thiz.rwMu.Lock()
	defer thiz.rwMu.Unlock()
	if len(thiz.templateNameList) == 1 {
		thiz.templateNameList = make([]string, 0)
		return
	}
	thiz.templateNameList = append(thiz.templateNameList[:idx], thiz.templateNameList[idx+1:]...)
}

func (thiz *repository) Get(idx int) string {
	thiz.rwMu.RLock()
	defer thiz.rwMu.RUnlock()
	return thiz.templateNameList[idx]
}

func (thiz *repository) List() []string {
	thiz.rwMu.RLock()
	defer thiz.rwMu.RUnlock()
	dest := make([]string, len(thiz.templateNameList))
	copy(dest, thiz.templateNameList)
	return dest
}

func (thiz *repository) Len() int {
	thiz.rwMu.RLock()
	defer thiz.rwMu.RUnlock()
	return len(thiz.templateNameList)
}
