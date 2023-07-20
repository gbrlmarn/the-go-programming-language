package wconv

// KToP converts kilograms weight to pounds
func KToP(k Kilogram) Pound { return Pound(k *2.2) }

// PToK converts pounds weight to kilograms
func PToK(p Pound) Kilogram { return Kilogram(p / 2.2) }
