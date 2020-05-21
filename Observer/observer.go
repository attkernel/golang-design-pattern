package Observer

type Obj interface {
	Update(o *Observer)
}

type Observer struct {
	objs []Obj
	name string
}

func NewObserver() *Observer {
	return &Observer{
		objs: make([]Obj, 0),
	}
}

func (s *Observer) Attach(o Obj) {
	s.objs = append(s.objs, o)
}

func (s *Observer) Notify() {
	for idx, _ := range s.objs {
		s.objs[idx].Update(s)
	}
}

func (s *Observer) ChangeName() {
	s.name = "change"
	s.Notify()
}
