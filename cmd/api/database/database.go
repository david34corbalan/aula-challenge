package database

import (
	"sync"
	follow "uala/pkg/follow/domain"
	tweets "uala/pkg/tweets/domain"
	users "uala/pkg/users/domain"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type DataBaseIntance struct {
	Writer       *gorm.DB
	Reader       *gorm.DB
	Transacction *gorm.DB
}

var lock = &sync.Mutex{}
var transacction *gorm.DB

func NewDataBaseIntance() *DataBaseIntance {
	return &DataBaseIntance{}
}

func (d *DataBaseIntance) Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tweets.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	d.Reader = db
	d.Writer = db
	d.Transacction = db
	return db
}

func (d *DataBaseIntance) SingletonDB() *gorm.DB {
	if d.Writer == nil {
		lock.Lock()
		defer lock.Unlock()
		if d.Writer == nil {
			d.Writer = &gorm.DB{}
		} else {
			// fmt.Println("instancia ya creada anteriormente")
		}
	} else {
		// fmt.Println("ya estaba creada.")
	}

	return d.Writer
}

func (d *DataBaseIntance) InitTransaction() {
	d.Transacction = d.SingletonDB().Begin()
}

func (d *DataBaseIntance) CommitTransaction() {
	d.Transacction.Commit()
}

func (d *DataBaseIntance) RollbackTransaction() {
	d.Transacction.Rollback()
}

func (d *DataBaseIntance) Migrations(migrations ...interface{}) {
	d.SingletonDB().AutoMigrate(migrations...)
}

func (d *DataBaseIntance) InsertManyRows() {
	user := users.SeedUsers()
	tweet := tweets.SeedTweets()
	followers := follow.SeedFollowers()
	d.SingletonDB().Model(users.Users{}).Create(user)
	d.SingletonDB().Model(tweets.Tweets{}).CreateInBatches(&tweet, 100)
	d.SingletonDB().Model(follow.Follow{}).Create(&followers)
}
