package db

type OrderDirection int

const (
	ORDER_DESCENDING = iota
	ORDER_ASCENDING
)

var orderDirection = map[OrderDirection]string{
	ORDER_DESCENDING: "desc",
	ORDER_ASCENDING:  "asc",
}

func (od OrderDirection) String() string {
	return orderDirection[od]
}
