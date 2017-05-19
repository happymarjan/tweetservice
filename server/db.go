package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
	"log"
	"sync"
)

//----------------------------------------------

type IAutoUUID interface {
	GenerateUUID() (string, error)
}

type Object struct{}

type AutoUUID struct {
	Object
	UUID string `sql:"size:37;unique_index"`
}

func (s *AutoUUID) GenerateUUID() (string, error) {
	s.UUID = uuid.NewV4().String()
	return s.UUID, nil
}

func GenerateUUIDIfSupported(ormObject interface{}) (string, error) {
	if autoUUIDObject, ok := ormObject.(IAutoUUID); ok {
		return autoUUIDObject.GenerateUUID()
	}
	return "", nil
}

//----------------------------------------------

type Tweet struct {
	AutoUUID
	Id          int    `gorm:"primary_key"`
	TweetID     string `sql:"size:35;unique_index"`
	TweetAuthor string `sql:"size:25"`
	TweetDate   string `sql:"size:35"`
	TweetText   string `sql:"size:256"`
}

//----------------------------------------------

type IDB interface {
	DB() *gorm.DB
}

type DB struct {
	db *gorm.DB
}

func (s *DB) DB() *gorm.DB {
	return s.db
}

var dbInstance *DB
var once sync.Once

type DB_out struct {
	db *gorm.DB
}

//----------------------------------------------

func getDB() *DB {
	once.Do(func() {
		dbInstance = &DB{}
	})
	return dbInstance
}

//----------------------------------------------

func (s *DB) Close() error {
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			log.Printf("Error in closing GORM DB: %v\n", err)
			return err
		}
		s.db = nil
	}
	return nil
}

//----------------------------------------------

func (s *DB) Open(connStr string) error {
	if s.db != nil {
		return fmt.Errorf("GORM DB has been opened before")
	}
	var err error
	if s.db, err = gorm.Open("postgres", connStr); err != nil {
		return err
	}
	s.db.LogMode(true)
	return nil
}

//----------------------------------------------

func (s *DB) CreateTable() error { //returns a pointer to the gorb DB
	if s.db.HasTable(&Tweet{}) {
		return fmt.Errorf("Model Tweet's table already exists")
	}

	s.db.DropTableIfExists(&Tweet{})
	//if err := s.db.CreateTable(&Tweet{}).Error; err != nil {
	if err := s.db.CreateTable(&Tweet{}).Error; err != nil {
		//if err := s.db.Table("second_tweets").CreateTable(&Tweet{}).Error; err != nil {
		log.Printf("Error in creating table with Tweet's model %v\n ", err)
		return err
	}
	return nil
}

//----------------------------------------------

func (s *DB) DeleteTable() error {
	if s.db != nil {
		if err := s.db.DropTableIfExists(&Tweet{}).Error; err != nil {
			fmt.Errorf("Error in dropping table with Tweet's model%v\n", err)
			return err
		}
	}
	return nil
}

//----------------------------------------------

func (s *DB) InitTables(force bool) error {
	if force {
		if err := s.DeleteTable(); err != nil {
			return err
		}
	}
	return s.CreateTable()
}

//----------------------------------------------

func (s *DB) populateTable(newsID string, newsAuthor string, newsDate string, newsText string) {
	//tweetObj := Tweet{tweetID: "12806767", tweetAuthor: "testAuth", tweetDate: "011-07-14T19:43:37+0100", tweetText: "test tweet"}
	tweetObj := Tweet{TweetID: newsID, TweetAuthor: newsAuthor, TweetDate: newsDate, TweetText: newsText}
	GenerateUUIDIfSupported(&tweetObj)
	s.DB().Create(&tweetObj)
}

//----------------------------------------------

func (s *DB) queryTable() []string {
	var tweets []Tweet
	var tweetTexts []string

	//if err := s.db.Where("tweet_id = ?", newsid).First(&Tweet{}).Error; err != nil {  //newsid being string argument
	//if err := s.db.Where("tweet_id = ?", newsid).Find(&tweets).Error; err != nil {
	//if err := s.db.Find(&tweets).Error; err != nil {

	if err := s.db.Find(&tweets).Pluck("tweet_text", &tweetTexts).Error; err != nil {
		fmt.Errorf("Error in excecuting query %v\n", err)
	}
	for _, text := range tweetTexts {
		fmt.Println(text)
	}
	return tweetTexts
}

//----------------------------------------------
