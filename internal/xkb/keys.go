package xkb

type keyTable struct {
	Return uint32
	Esc    uint32
	Space  uint32

	LeftPointer  uint32
	RightPointer uint32

	A, a uint32
	B, b uint32
	C, c uint32
	D, d uint32
	E, e uint32
	F, f uint32
	G, g uint32
	H, h uint32
	I, i uint32
	J, j uint32
	K, k uint32
	L, l uint32
	M, m uint32
	N, n uint32
	O, o uint32
	P, p uint32
	Q, q uint32
	R, r uint32
	S, s uint32
	T, t uint32
	U, u uint32
	V, v uint32
	W, w uint32
	X, x uint32
	Y, y uint32
	Z, z uint32
}

var Key = keyTable{
	Return: 0xff0d,
	Esc:    0xff1b,
	Space:  0x0020,

	LeftPointer:  0x110,
	RightPointer: 0x111,

	A: 0x0041, a: 0x0061,
	B: 0x0042, b: 0x0062,
	C: 0x0043, c: 0x0063,
	D: 0x0044, d: 0x0064,
	E: 0x0045, e: 0x0065,
	F: 0x0046, f: 0x0066,
	G: 0x0047, g: 0x0067,
	H: 0x0048, h: 0x0068,
	I: 0x0049, i: 0x0069,
	J: 0x004a, j: 0x006a,
	K: 0x004b, k: 0x006b,
	L: 0x004c, l: 0x006c,
	M: 0x004d, m: 0x006d,
	N: 0x004e, n: 0x006e,
	O: 0x004f, o: 0x006f,
	P: 0x0050, p: 0x0070,
	Q: 0x0051, q: 0x0071,
	R: 0x0052, r: 0x0072,
	S: 0x0053, s: 0x0073,
	T: 0x0054, t: 0x0074,
	U: 0x0055, u: 0x0075,
	V: 0x0056, v: 0x0076,
	W: 0x0057, w: 0x0077,
	X: 0x0058, x: 0x0078,
	Y: 0x0059, y: 0x0079,
	Z: 0x005a, z: 0x007a,
}
