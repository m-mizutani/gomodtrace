package main

type Dependency struct {
	Src string
	Dst string
}

type Node struct {
	Name       string  `json:"pkg"`
	DependedBy []*Node `json:"depended_by"`
	Looped     bool    `json:"-"`
}

type DependMap struct {
	FwdIndex map[string][]string
	RevIndex map[string][]string
}

func NewDependMap(dependencies []*Dependency) *DependMap {
	m := &DependMap{
		FwdIndex: make(map[string][]string),
		RevIndex: make(map[string][]string),
	}
	for _, dep := range dependencies {
		m.FwdIndex[dep.Src] = append(m.FwdIndex[dep.Src], dep.Dst)
		m.RevIndex[dep.Dst] = append(m.RevIndex[dep.Dst], dep.Src)
	}

	return m
}

type traceState []string

func (x traceState) has(s string) bool {
	for _, v := range x {
		if v == s {
			return true
		}
	}
	return false
}

func (x traceState) push(s string) traceState {
	return append(x, s)
}

func (x *DependMap) Trace(target string) *Node {
	return x.trace(target, traceState{})
}

func (x *DependMap) trace(target string, state traceState) *Node {
	if state.has(target) {
		return &Node{
			Name:   target,
			Looped: true,
		}
	}

	pkgs, ok := x.RevIndex[target]

	if !ok {
		if _, has := x.FwdIndex[target]; has {
			return &Node{
				Name: target,
			}
		}
		return nil
	}

	var edges []*Node
	for _, pkg := range pkgs {
		if node := x.trace(pkg, state.push(target)); node != nil {
			edges = append(edges, node)
		}
	}

	return &Node{
		Name:       target,
		DependedBy: edges,
	}
}
