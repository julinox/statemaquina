```markdown
# Finite State Machine

This project implements a Finite State Machine (FSM) using the State design pattern in Go.  
The State pattern allows an object to change its behavior when its internal state changes.  
This implementation provides a flexible and scalable way to manage state transitions and  
actions associated with each state.

## Key Components

- **State Interface**: Defines the methods that each state must implement.
  ```go
  type State interface {
      Name() string
      Next() (int, error)
  }
  ```
## License

This project is licensed under the MIT License - see the LICENSE file for details.
```