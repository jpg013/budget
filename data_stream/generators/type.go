package generators

type Chunk interface{}

type Generator interface {
	Next() (Chunk, error)
}
