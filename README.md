```markdown
# Finite State Machine

This project implements a Finite State Machine (FSM) using the State design pattern in Go. The State pattern allows an object to change its behavior when its internal state changes. This implementation provides a flexible and scalable way to manage state transitions and actions associated with each state.

## Key Components

- **State Interface**: Defines the methods that each state must implement.
  ```go
  type State interface {
      Name() string
      Next() (int, error)
  }
  ```

- **StateMac Interface**: Defines the methods for starting the state machine, posting events, registering states, and setting the maximum count of transitions.
  ```go
  type StateMac interface {
      Start() error
      Post(int)
      Register(State, int) error
      SetMaxCount(int)
  }
  ```

- **StateMacCfg Struct**: Configuration for the state machine, including options to stop on error and stop after a certain number of transitions.
  ```go
  type StateMacCfg struct {
      StopOnError bool
      StopOnCount int
      Lg          *logrus.Logger
  }
  ```

## Usage

To create a new state machine, initialize it with a configuration and a list of states (cfg can be nil):
```go
func NewMaquinaEstado(cfg *StateMacCfg, states ...State) (StateMac, error) 
```

This implementation provides a robust framework for managing complex state transitions and behaviors in a clean and maintainable way.

## Error Handling

The following errors are defined for handling various state machine issues:
```go
var (
    ErrorStateNotFound = fmt.Errorf("state not found")
    ErrorStateReg      = fmt.Errorf("state already registered")
    ErrorMaxCount      = fmt.Errorf("max count reached")
)
```

## Logging

The state machine uses `logrus` for logging. Ensure you have it installed:
```sh
go get github.com/sirupsen/logrus
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
```