package state_machine

// StatePersistenceManager interface defines the signature of a persistence manager that state machine
// uses to persist state machine and recover from persisted state. Any type of persistence manager
// need to implement this interface to be operable by state machine.
type StatePersistenceManager interface {
	// Save the state data supplied as byte slice to the persistence storage, it returns error in case of persistence is not possible
	Save([]byte) error

	// Load the state data from persistence storage, it should return byte slice and error in case fo persisted data can not be recovered
	Load() ([]byte, error)

	// Delete the persisted data if nay
	Delete()
}
