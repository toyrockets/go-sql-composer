package sqlcomposer

type JoinType int

const (
	InnerJoin      JoinType = iota
	LeftOuterJoin  JoinType = iota
	RightOuterJoin JoinType = iota
	FullOuterJoin  JoinType = iota
	CrossJoin      JoinType = iota
)
