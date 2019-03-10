package golife

const (
	t = true
	f = false
)

// Glider
var Glider = [][]bool{
	{f, f, t},
	{t, f, t},
	{f, t, t},
}

// Glider Gun
var GliderGun = [][]bool{
	{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, t, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, t, f, t, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, f, t, t, f, f, f, f, f, f, t, t, f, f, f, f, f, f, f, f, f, f, f, f, t, t, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, t, f, f, f, t, f, f, f, f, t, t, f, f, f, f, f, f, f, f, f, f, f, f, t, t, f},
	{f, t, t, f, f, f, f, f, f, f, f, t, f, f, f, f, f, t, f, f, f, t, t, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, t, t, f, f, f, f, f, f, f, f, t, f, f, f, t, f, t, t, f, f, f, f, t, f, t, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, t, f, f, f, f, f, t, f, f, f, f, f, f, f, t, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, t, f, f, f, t, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, f, t, t, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
}

// Galaxy
var Galaxy = [][]bool{
	{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, t, t, f, t, t, t, t, t, t, f, f, f},
	{f, f, f, t, t, f, t, t, t, t, t, t, f, f, f},
	{f, f, f, t, t, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, t, t, f, f, f, f, f, t, t, f, f, f},
	{f, f, f, t, t, f, f, f, f, f, t, t, f, f, f},
	{f, f, f, t, t, f, f, f, f, f, t, t, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, t, t, f, f, f},
	{f, f, f, t, t, t, t, t, t, f, t, t, f, f, f},
	{f, f, f, t, t, t, t, t, t, f, t, t, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
	{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
}
