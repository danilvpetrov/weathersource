package deps

import "github.com/google/wire"

// CommonDepsSet is the wire set that contains common dependencies for the
// executable files in cmd/* packages.
var CommonDepsSet = wire.NewSet(ProvideDatabase, ProvideLogger)
