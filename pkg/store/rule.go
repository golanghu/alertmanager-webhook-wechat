// TODO memery Cache
package store

import "sync"

type RuleInformation interface {
	GetApp(name string) (bool, *tb_app)
	GetRule(id int) (bool, *tb_rule)
	SetApp(name string, app *tb_app)
	SetRule(id int, rule *tb_rule)
	DelApp(name string)
	DelRule(id int)
}

type tb_rule struct {
}
type tb_app struct {
}
type RuleInfo struct {
	mu    sync.Mutex
	Rules map[int]*tb_rule
	Apps  map[string]*tb_app
}

func (r *RuleInfo) GetApp(name string) (bool, *tb_app) {
	r.mu.Lock()
	defer r.mu.Unlock()
	app, ok := r.Apps[name]
	if ok {
		return ok, app
	}
	return false, nil
}
func (r *RuleInfo) GetRule(id int) (bool, *tb_rule) {
	r.mu.Lock()
	defer r.mu.Unlock()
	rule, ok := r.Rules[id]
	if ok {
		return ok, rule
	}
	return false, nil
}

func (r *RuleInfo) SetApp(name string, app *tb_app) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Apps[name] = app
}
func (r *RuleInfo) SetRule(id int, rule *tb_rule) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Rules[id] = rule
}
func (r *RuleInfo) DelApp(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.Apps, name)
}
func (r *RuleInfo) DelRule(id int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.Rules, id)
}
func (r *RuleInfo) LoadData() error {
	return nil
}
