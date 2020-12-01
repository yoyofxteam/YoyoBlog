package domain

type Page struct {
	PageIndex  int
	PageSize   int
	TotalCount int
	Data       interface{}
}
