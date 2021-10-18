package repo

import (
	"sync"
)

type templateListRepo struct {
	templateNameList []string
	rwMu             sync.RWMutex
}

var (
	templateListOnce sync.Once
	templList        *templateListRepo
)

func TemplateList() *templateListRepo {
	templateListOnce.Do(func() {
		templList = &templateListRepo{
			templateNameList: make([]string, 0, 30),
		}
	})
	return templList
}

func (thiz *templateListRepo) Append(newTemplateName string) {
	thiz.rwMu.Lock()
	defer thiz.rwMu.Unlock()
	thiz.templateNameList = append(thiz.templateNameList, newTemplateName)
}

func (thiz *templateListRepo) Delete(idx int) {
	thiz.rwMu.Lock()
	defer thiz.rwMu.Unlock()
	if len(thiz.templateNameList) == 1 {
		thiz.templateNameList = make([]string, 0)
		return
	}
	thiz.templateNameList = append(thiz.templateNameList[:idx], thiz.templateNameList[idx+1:]...)
}

func (thiz *templateListRepo) Get(idx int) string {
	thiz.rwMu.RLock()
	defer thiz.rwMu.RUnlock()
	return thiz.templateNameList[idx]
}

func (thiz *templateListRepo) List() []string {
	thiz.rwMu.RLock()
	defer thiz.rwMu.RUnlock()
	dest := make([]string, len(thiz.templateNameList))
	copy(dest, thiz.templateNameList)
	return dest
}

func (thiz *templateListRepo) Len() int {
	thiz.rwMu.RLock()
	defer thiz.rwMu.RUnlock()
	return len(thiz.templateNameList)
}
