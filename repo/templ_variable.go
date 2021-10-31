package repo

import (
	"sync"

	"fyne.io/fyne/v2/widget"
)

type emailVar struct {
	Id       string
	KeyEntry *widget.Entry
	ValEntry *widget.Entry
}

type emailVarListRepo struct {
	variableList []*emailVar
	rwMu         sync.RWMutex
}

var (
	emailVariablesOnce sync.Once
	emailVarList       *emailVarListRepo
)

func EmailVarList() *emailVarListRepo {
	emailVariablesOnce.Do(func() {
		emailVarList = &emailVarListRepo{
			variableList: make([]*emailVar, 0, 5),
		}
	})
	return emailVarList
}

func (thiz *emailVarListRepo) Append(id string, key, val *widget.Entry) {
	thiz.rwMu.Lock()
	defer thiz.rwMu.Unlock()
	thiz.variableList = append(thiz.variableList,
		&emailVar{Id: id, KeyEntry: key, ValEntry: val})
}

func (thiz *emailVarListRepo) Delete(idx int) {
	thiz.rwMu.Lock()
	defer thiz.rwMu.Unlock()
	if len(thiz.variableList) == 1 {
		thiz.variableList = make([]*emailVar, 0)
		return
	}
	thiz.variableList = append(thiz.variableList[:idx], thiz.variableList[idx+1:]...)
}

func (thiz *emailVarListRepo) Len() int {
	thiz.rwMu.RLock()
	defer thiz.rwMu.RUnlock()
	return len(thiz.variableList)
}

func (thiz *emailVarListRepo) Map() map[string]interface{} {
	thiz.rwMu.RLock()
	defer thiz.rwMu.RUnlock()
	dest := make(map[string]interface{}, len(thiz.variableList))
	for _, data := range thiz.variableList {
		dest[data.KeyEntry.Text] = data.ValEntry.Text
	}
	return dest
}

func (thiz *emailVarListRepo) List() []*emailVar {
	thiz.rwMu.RLock()
	defer thiz.rwMu.RUnlock()
	dest := make([]*emailVar, len(thiz.variableList))
	copy(dest, thiz.variableList)
	return dest
}
