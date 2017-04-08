package pgdatabase

import (
	"database/sql"
	"../models"
)

type DAO struct {
	connection *Connection
}

func NewDAO() *DAO {
	dao := new(DAO)
	dao.connection = new(Connection)

	return dao
}

func (this *DAO) AddRawPost(post []byte) {
	//var transaction bool = false
	//if this.connection.IsConnected() {
	//	transaction = true
	//}
	//if !transaction {
	//	this.connection.Connect()
	//}
	this.connection.Connect()
	stmt := this.connection.Prepare(readSQLFile("resources/sql/insertRawPost.sql"))
	stmt.Exec(post)
	this.connection.Close()
	//if !transaction {
	//	this.connection.Close()
	//}
}

func (this *DAO) AddProcessedPost(post *models.ProcessedPost) {
	this.connection.Connect()
	stmt := this.connection.Prepare(readSQLFile("resources/sql/insertProcessedPost.sql"))
	_, err := stmt.Exec(post.PostLength, post.LikeCount, post.FollowersCount, post.FriendCount, post.HashtagCount, post.RetweetCount, post.IsRetweet, post.KloutScore, post.ExclaimationCount)
	sqlError(err, "Error in AddProcessedPost")
	this.connection.Close()
}

func (this *DAO) GetGeoHash(geohash string) *sql.Rows{
	var transaction bool
	if this.connection.IsConnected() {
		transaction = true
	}
	if !transaction {
		this.connection.Connect()
	}
	stmt := this.connection.Prepare(readSQLFile("resources/sql/getGeoHash.sql"))
	result, err := stmt.Query(geohash)
	sqlError(err, "Error in GetGeoHash")

	return result
}

func (this *DAO) Test() {
	this.connection.Test()
}