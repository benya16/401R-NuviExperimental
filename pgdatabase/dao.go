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
	_, err := stmt.Exec(id, post.PostLength, post.LikeCount, post.FollowersCount, post.FriendCount, post.HashtagCount, post.RetweetCount, post.IsRetweet, post.KloutScore, post.ExclaimationCount)
	sqlError(err, "Error in AddProcessedPost()")
	if post.ActiveShooter {
		this.UpdateDangerWord(id, "active_shooter", true)
	}
	if post.Attack {
		this.UpdateDangerWord(id, "attack", true)
	}
	if post.Attacker {
		this.UpdateDangerWord(id, "attacker", true)
	}
	if post.Bomb {
		this.UpdateDangerWord(id, "bomb", true)
	}
	if post.BombThreat {
		this.UpdateDangerWord(id, "bomb_threat", true)
	}
	if post.Breaking {
		this.UpdateDangerWord(id, "breaking", true)
	}
	if post.Danger {
		this.UpdateDangerWord(id, "danger", true)
	}
	if post.Dead{
		this.UpdateDangerWord(id, "dead", true)
	}
	if post.Gunman {
		this.UpdateDangerWord(id, "gunman", true)
	}
	if post.Killing {
		this.UpdateDangerWord(id, "killing", true)
	}
	if post.Rape {
		this.UpdateDangerWord(id, "rape", true)
	}
	if post.Shooter {
		this.UpdateDangerWord(id, "shooter", true)
	}
	if post.Shooting {
		this.UpdateDangerWord(id, "shooting", true)
	}
	if post.Stabbing {
		this.UpdateDangerWord(id, "stabbing", true)
	}
	if post.Terrorist {
		this.UpdateDangerWord(id, "terrorist", true)
	}
	if post.Warning {
		this.UpdateDangerWord(id, "warning", true)
	}
	this.connection.Close()
}

func (this *DAO) UpdateDangerWord(id string, word string, state bool) {
	var transaction bool = false
	if this.connection.IsConnected() {
		transaction = true
	}
	if !transaction {
		this.connection.Connect()
	}
	stmt := this.connection.Prepare(readSQLFile("resources/sql/updateDangerWord.sql"))
	_, err := stmt.Exec(id, word, state)
	sqlError(err, "Error in ")
	if !transaction {
		this.connection.Close()
	}
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