package inspection

import "sync/atomic"

type Level uint32

const (
    Base Level = iota
    Detail
)

type Inspection struct {
    Level Level
}

func New() *Inspection {
    return &Inspection{
        Level: Base,
    }
}

var (
    inspection = New()
)

// SetLevel sets the standard logger level.
func SetLevel(level Level) {
    inspection.setLevel(level)
}

// GetLevel sets the standard logger level.
func GetLevel() Level {
    return inspection.getLevel()
}

func (logger *Inspection) setLevel(level Level) {
    atomic.StoreUint32((*uint32)(&logger.Level), uint32(level))
}

func (logger *Inspection) getLevel() Level {
    return Level(atomic.LoadUint32((*uint32)(&logger.Level)))
}
