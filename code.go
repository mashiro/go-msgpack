package msgpack

// Byte returns the byte representation of the Code
func (t Code) Byte() byte {
	return byte(t)
}

// IsMapFamily returns true if the given code is equivalent to
// one of the `map` family in msgpack
func IsMapFamily(c Code) bool {
	b := c.Byte()
	return (b >= FixMap0.Byte() && b <= FixMap15.Byte()) ||
		b == Map16.Byte() ||
		b == Map32.Byte()
}

func IsArrayFamily(c Code) bool {
	b := c.Byte()
	return (b >= FixArray0.Byte() && b <= FixArray15.Byte()) ||
		b == Array16.Byte() ||
		b == Array32.Byte()
}

func IsStrFamily(c Code) bool {
	b := c.Byte()
	return (b >= FixStr0.Byte() && b <= FixStr31.Byte()) ||
		b == Str8.Byte() ||
		b == Str16.Byte() ||
		b == Str32.Byte()
}

func IsBinFamily(c Code) bool {
	b := c.Byte()
	return b == Bin8.Byte() || b == Bin16.Byte() || b == Bin32.Byte()
}

func IsExtFamily(c Code) bool {
	b := c.Byte()
	return b == Ext8.Byte() || b == Ext16.Byte() || b == Ext32.Byte() ||
		b == FixExt1.Byte() || b == FixExt2.Byte() || b == FixExt4.Byte() || b == FixExt8.Byte() || b == FixExt16.Byte()
}
