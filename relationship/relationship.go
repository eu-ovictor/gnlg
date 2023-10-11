package relationship

type Members struct {
	Parent int64 `json:"parent"`
	Child  int64 `json:"child"`
}

type NamedMembers struct {
	Parent string
	Child  string
}

type Kinship string

const (
	Parent Kinship = "parent"
	Child  Kinship = "child"
)

type Relationship struct {
    Member  string `json:"name"`
    Kinship Kinship `json:"relationship"`
}

type Relationships map[string][]Relationship
