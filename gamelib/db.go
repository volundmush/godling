package gamelib

type DbRef int

type DbObjType struct {
	name   string
	letter string
}

type DbObjFlag struct {
	name     string
	letter   string
	holders  []DbObject
	objTypes []string
}

type DbNamespace struct {
	name    string
	letter  string
	holders []DbObject
}

type LockFlag struct {
}

type AttrFlag struct {
}

type DbLockValue struct {
	name    string
	creator DbRef
	flags   []LockFlag
}

type AttrNode struct {
	name    string
	creator DbRef
	flags   []AttrFlag
	nodes   map[string]AttrNode
}

type AttrBase struct {
	nodes map[string]AttrNode
}

type DbObject struct {
	dbRef      DbRef
	objType    *DbObjType
	namespace  *DbNamespace
	name       MarkupString
	flags      []DbObjFlag
	powers     []DbObjFlag
	location   *DbObject
	contents   []DbObject
	zone       *DbObject
	zoned      []DbObject
	owner      *DbObject
	belongings []DbObject
	parent     *DbObject
	children   []DbObject
	created    int
	modified   int
	attributes AttrBase
	authority  int
}

type GameDB struct {
	dbObjects  map[DbRef]DbObject
	nextref    DbRef
	attributes AttrBase
}
