/*
create_time: 2020-9-9
author: korei
*/
package model

type Subforum struct {
	Name	string
	Blocks 	[]Block
}

type Block struct {
	Name 		string
	Master 		User
}
