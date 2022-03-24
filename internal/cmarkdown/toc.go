package cmarkdown

type CMImage struct {
	URL string
}

type CMTOC struct {
	ID    string
	Name  string
	Level CMTOCLevel
}

type CMTOCLevel int

const (
	CMTOCLevelH1 CMTOCLevel = 1
	CMTOCLevelH2 CMTOCLevel = 2
	CMTOCLevelH3 CMTOCLevel = 3
	CMTOCLevelH4 CMTOCLevel = 4
	CMTOCLevelH5 CMTOCLevel = 5
)

func NewTOCLevel(tag string) CMTOCLevel {
	switch tag {
	case "h1":
		return CMTOCLevelH1
	case "h2":
		return CMTOCLevelH2
	case "h3":
		return CMTOCLevelH3
	case "h4":
		return CMTOCLevelH4
	case "h5":
		return CMTOCLevelH5
	}
	return CMTOCLevelH5
}
