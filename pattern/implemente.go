

type Flyer interface {
	Fly() error
}

type Wing struct {
	color string
	height int
}

func (w* Wing) Fly() err error {
	// fly
	return err
}


type Bird struct {
	Wing
	kind string
	weight int
}

type Plane struct {
	Wing
	company string
}


func HighLevelFunc(f Flyer) {
	f.Fly()
}

// ============================================================================
// upper convert




// ============================================================================
// down convert




