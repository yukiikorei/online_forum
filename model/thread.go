/*
create_time: 2020-9-9
author: korei
*/
package model

import (
	"time"
)

type Thread struct {
	ID 			int
	Tittle		string
	UpdateTime 	time.Time

	AtBlock 	Block
	Theme		string			//r
	Author 		User			//TODO: reference
	Comments 	[]Comment		//TODO: related
}

type Comment struct{
	Author 		User
	CreateTime	time.Time
	Content 	string
}

