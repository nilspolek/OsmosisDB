package paser

type Paser struct{}

type Command struct {
	Type string
}

const (
	TYPEINT    = '!'
	TYPESTRING = '@'
	TYPEBOOL   = '#'
	TYPEFLOAT  = '$'
	SET        = "SET"
	GET        = "GET"
	DEL        = "DEL"
	UPT        = "UPT"
	ERR        = "ERR"
	OK         = "OK "
)

func NewPaser() *Paser {
	return &Paser{}
}

func (p *Paser) Parse() Command {
	// Parse the input
}
