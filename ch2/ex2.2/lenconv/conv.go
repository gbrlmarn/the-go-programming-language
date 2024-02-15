package lenconv

// MToF converts a meter length to feet
func MToF(m Meter) Foot { return Foot(m * 0.3048) }

// FToM converts a feet length to meter
func FToM(f Foot) Meter { return Meter(f / 0.3048) }
