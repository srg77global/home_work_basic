package hw04

type Book struct {
	id     uint32
	title  string
	author string
	year   uint16
	size   uint16
	rate   float32
}

func (s *Book) WriteID(x uint32) {
	s.id = x
}

func (s *Book) WriteTitle(x string) {
	s.title = x
}

func (s *Book) WriteAuthor(x string) {
	s.author = x
}

func (s *Book) WriteYear(x uint16) {
	s.year = x
}

func (s *Book) WriteSize(x uint16) {
	s.size = x
}

func (s *Book) WriteRate(x float32) {
	s.rate = x
}

func (s *Book) GetID() uint32 {
	return s.id
}

func (s *Book) GetTitle() string {
	return s.title
}

func (s *Book) GetAuthor() string {
	return s.author
}

func (s *Book) GetYear() uint16 {
	return s.year
}

func (s *Book) GetSize() uint16 {
	return s.size
}

func (s *Book) GetRate() float32 {
	return s.rate
}

const (
	year = iota
	size
	rate
)

type CBooks struct {
	enum uint8
}

func (s CBooks) Comparator(a, b Book) bool {
	switch s.enum {
	case 0:
		return a.year > b.year
	case 1:
		return a.size > b.size
	case 2:
		return a.rate > b.rate
	default:
		return false
	}
}

func createStruct(x uint8) *CBooks {
	return &CBooks{enum: x}
}
