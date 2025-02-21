# Finite State Machine (State Pattern)

Implements the State design pattern in Go. This pattern allows an object to change its behavior when its internal state changes.  
This implementation provides a flexible and scalable way to manage state transitions and actions associated with each state.

## Key Components
  ```go
  type State interface {
      Name() string
      Next() (int, error) // returns next state or error
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

  var (
    ErrorStateNotFound = fmt.Errorf("state not found")
    ErrorStateReg      = fmt.Errorf("state already registered")
    ErrorMaxCount      = fmt.Errorf("max count reached")
  )
  ```
## Usage

To create a new state machine, initialize it with a configuration and a list of states (cfg can be nil):  
Also can use 'Register()' function instead of 'states...'
```go
func NewMaquinaEstado(cfg *StateMacCfg, states ...State) (StateMac, error)
```
## Example
```go
machine, _ := NewMaquinaEstado(nil, &StateA{}, &StateB{}, &StateC{})
machine.Post(1) // Start on 'StateA'
machine.Start()
```
## License

This project is licensed under the MIT License - see the LICENSE file for details.
