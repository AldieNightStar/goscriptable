package goscriptable

import "strings"

// ============================

type Params struct {
	m map[string]*ParamValue
}

func newParams() *Params {
	return &Params{m: make(map[string]*ParamValue, 32)}
}

func (p *Params) append(name string, val string) {
	v, ok := p.m[name]
	if !ok {
		v = newParamVal()
		p.m[name] = v
	}
	v.values = append(v.values, val)
}

func (p *Params) Get(name string) *ParamValue {
	v, ok := p.m[name]
	if !ok {
		return nil
	}
	return v
}

func (p *Params) IsPresent(name string) bool {
	return p.Get(name) != nil
}

func (p *Params) Len() int {
	return len(p.m)
}

// ============================

type ParamValue struct {
	values []string
}

func newParamVal() *ParamValue {
	return &ParamValue{values: make([]string, 0, 8)}
}

func (v *ParamValue) First() string {
	if len(v.values) < 1 {
		return ""
	}
	return v.values[0]
}

func (v *ParamValue) Len() int {
	return len(v.values)
}

func (v *ParamValue) All() []string {
	return v.values
}

// ============================

// Parses arguments in format: -a arg1 -b arg2 -c -d -e -arr a b c d
//		Example:
//	 		-f filename -t title -z zip2 -link http://google.com/ -o -p -s -q
func ParseArgs(args []string) *Params {
	params := newParams()
	name := ""
	paramCount := 0
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if name != "" && paramCount == 0 {
				params.append(name, "true")
			}
			name = arg[1:]
			paramCount = 0
			continue
		}
		if name == "" {
			continue
		}
		params.append(name, arg)
		paramCount += 1
	}
	if name != "" && paramCount == 0 {
		params.append(name, "true")
	}
	return params
}
