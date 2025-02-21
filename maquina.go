package statemaquina

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	ErrorStateNotFound = fmt.Errorf("state not found")
	ErrorStateReg      = fmt.Errorf("state already registered")
	ErrorMaxCount      = fmt.Errorf("max count reached")
)

type State interface {
	Name() string
	Next() (int, error)
}

type StateMac interface {
	Start() error
	Post(int)
	Register(State, int) error
	SetMaxCount(int)
}

type StateMacCfg struct {
	StopOnError bool
	StopOnCount int
	Lg          *logrus.Logger
}

type xMaquinaDelMal struct {
	transit    []int
	transitPtr int
	table      map[int]State
	cfg        *StateMacCfg
}

func NewMaquinaEstado(cfg *StateMacCfg, states ...State) (StateMac, error) {

	var newMaquina xMaquinaDelMal

	if cfg == nil {
		newMaquina.cfg = defaultStateMacCfg()
	} else {
		newMaquina.cfg = cfg
	}

	newMaquina.transit = make([]int, 0)
	newMaquina.table = make(map[int]State)
	for i, state := range states {
		if err := newMaquina.Register(state, i+1); err != nil {
			newMaquina.print(err.Error())
			if newMaquina.cfg.StopOnError {
				return nil, err
			}
		}
	}

	return &newMaquina, nil
}

func (x *xMaquinaDelMal) Start() error {

	transitCounter := 0
	for {
		if transitCounter >= x.transitPtr {
			break
		}

		nextTransition := x.transit[transitCounter]
		transitCounter++
		if x.cfg.StopOnCount > 0 && transitCounter > x.cfg.StopOnCount {
			return ErrorMaxCount
		}

		if x.table[nextTransition] == nil {
			if x.cfg.StopOnError {
				return ErrorStateNotFound
			}

			x.print(fmt.Sprintf("Next state '%v' not found", nextTransition))
			continue
		}

		nexState, err := x.table[nextTransition].Next()
		if err != nil {
			if x.cfg.StopOnError {
				return err
			}

			x.print(fmt.Sprintf("transition error: %v -> %v", nextTransition, err))
			continue
		}

		if nexState > 0 {
			x.Post(nexState)
		}
	}

	return nil
}

func (x *xMaquinaDelMal) Post(event int) {
	x.transit = append(x.transit, event)
	x.transitPtr++
}

// State '0' is reserved for the initial state
func (x *xMaquinaDelMal) Register(state State, id int) error {

	if state == nil {
		return fmt.Errorf("nil state registry attempt")
	}

	if id <= 0 {
		return fmt.Errorf("invalid id '%v' for state '%s'", id, state.Name())
	}

	if _, ok := x.table[id]; ok {
		return ErrorStateReg
	}

	x.table[id] = state
	x.print(fmt.Sprintf("State %s registered with id %d", state.Name(), id))
	return nil
}

func (x *xMaquinaDelMal) SetMaxCount(count int) {
	x.cfg.StopOnCount = count
}

func (m *xMaquinaDelMal) print(msg string) {

	if m.cfg.Lg != nil {
		m.cfg.Lg.Info(msg)
	} else {
		fmt.Println(msg)
	}
}

func defaultStateMacCfg() *StateMacCfg {

	var cfg StateMacCfg

	cfg.StopOnError = true
	return &cfg
}
