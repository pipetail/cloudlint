package inspection

import "sync/atomic"

// Level basic structure to define level of inspection
type Level uint32

// Level iota
const (
    BASE   Level = iota // BASE == 0
    DETAIL              // DETAIL == 1
)

// Inspection structure
type Inspection struct {
    Level Level
}

// New creates new instance of Inspection structure
func New() *Inspection {
    return &Inspection{
        Level: BASE,
    }
}

var (
    inspection = New()
)

// SetLevel sets the standard logger level.
func SetLevel(level Level) {
    inspection.setLevel(level)
}

// CheckDetail is true if level of check is DETAIL.
func CheckDetail() bool {
    return inspection.getLevel() == DETAIL
}

func (logger *Inspection) setLevel(level Level) {
    atomic.StoreUint32((*uint32)(&logger.Level), uint32(level))
}

func (logger *Inspection) getLevel() Level {
    return Level(atomic.LoadUint32((*uint32)(&logger.Level)))
}
