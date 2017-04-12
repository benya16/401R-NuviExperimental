package pgdatabase

import (
	"database/sql"
	"../models"
	"fmt"
)

type DAO struct {
	connection *Connection
}

func NewDAO() *DAO {
	dao := new(DAO)
	dao.connection = new(Connection)

	return dao
}

func (this *DAO) AddRawPost(id string, post []byte) {
	this.connection.Connect()
	stmt := this.connection.Prepare(readSQLFile("resources/sql/insertRawPost.sql"))
	_, err :=stmt.Exec(id, post)
	sqlError(err, "Error in AddRawPost()")
	this.connection.Close()
}

func (this *DAO) AddProcessedPost(id string, post *models.ProcessedPost) {
	fmt.Println(post)
	this.connection.Connect()
	stmt := this.connection.Prepare(readSQLFile("resources/sql/insertProcessedPost.sql"))
	_, err := stmt.Exec(id, post.PostLength, post.LikeCount, post.FollowersCount, post.FriendCount, post.HashtagCount, post.RetweetCount, post.IsRetweet, post.KloutScore, post.ExclaimationCount,
		post.ActiveShooter, post.Attack, post.Attacker, post.Bomb, post.BombThreat, post.Breaking, post.Danger, post.Dead, post.Gunman, post.Killing, post.Rape, post.Shooter, post.Shooting, post.Stabbing, post.Terrorist, post.Warning)
	sqlError(err, "Error in AddProcessedPost()")
	this.connection.Close()
}

func (this *DAO) LabelThreat(id string, threat bool) {
	this.connection.Connect()
	stmt := this.connection.Prepare(readSQLFile("resources/sql/updateThreat.sql"))
	_, err := stmt.Exec(id, threat)
	sqlError(err, "Error in LabelThreat()")
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