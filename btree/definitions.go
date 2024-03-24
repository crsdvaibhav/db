package btree

const (
	BNODE_INTERNAL = 1
	BNODE_LEAF     = 2
	HEADER = 4
	BTREE_PAGE_SIZE = 4096 //4 kb
	BTREE_MAX_KEY_SIZE = 1000
	BTREE_MAX_VAL_SIZE = 3000
)

type BNode struct {
	data []byte //array of bytes
}

type BTree struct {
	// pointer (a nonzero page number)
	root uint64

	// callbacks for managing on-disk pages
	get func(uint64) BNode // dereference a pointer
	new func(BNode) uint64 // allocate a new page
	del func(uint64)       // deallocate a page
}