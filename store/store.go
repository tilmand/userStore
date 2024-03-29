package store

import (
	"database/sql"
	"fmt"
	"log"
	"userStore/config"

	_ "github.com/go-sql-driver/mysql"
)

type obj map[string]interface{}

type Store struct {
	UsersRepository *UsersRepository
	AuthRepository  *AuthRepository
	db              *sql.DB
}

func NewSqlStore(conf *config.Config) (*Store, error) {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	initializeDB(db)
	store := &Store{
		db: db,
	}
	store.UsersRepository = store.Users()
	store.AuthRepository = store.Auth()

	return store, nil
}

func (s *Store) Users() *UsersRepository {
	if s.UsersRepository == nil {
		s.UsersRepository = NewUsersRepository(s)
	}

	return s.UsersRepository
}

func (s *Store) Auth() *AuthRepository {
	if s.AuthRepository == nil {
		s.AuthRepository = NewAuthRepository(s)
	}

	return s.AuthRepository
}

func initializeDB(db *sql.DB) {
	createTableQueries := []string{
		`CREATE TABLE IF NOT EXISTS auth (
			id bigint(20) NOT NULL AUTO_INCREMENT,
			api_key varchar(32) NOT NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		`CREATE TABLE IF NOT EXISTS user (
			id bigint(20) NOT NULL AUTO_INCREMENT,
			username varchar(64) NOT NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		`CREATE TABLE IF NOT EXISTS user_profile (
			user_id bigint(20) NOT NULL,
			first_name varchar(32) NOT NULL,
			last_name varchar(64) NOT NULL,
			phone varchar(64) NOT NULL,
			address varchar(64) NOT NULL,
			city varchar(64) NOT NULL,
			PRIMARY KEY (user_id),
			FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		`CREATE TABLE IF NOT EXISTS user_data (
			user_id bigint(20) NOT NULL,
			school varchar(32) NOT NULL,
			PRIMARY KEY (user_id),
			FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
	}

	for _, query := range createTableQueries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	insertDataQueries := []string{
		`INSERT INTO auth VALUES (1,'www-dfq92-sqfwf'),(2,'ffff-2918-xcas');`,
		`INSERT INTO user VALUES (1,'test'),(2,'admin'),(3,'guest');`,
		`INSERT INTO user_data VALUES (1,'гімназія №179 міста Києва'),(2,'ліцей №227'),(3,'Медична гімназія №33 міста Києва');`,
		`INSERT INTO user_profile VALUES (1,'Александр','Школьный','+38050123455','ул. Сибирская 2','Киев'),(2,'Дмитрий','Арбузов','+38065133223','ул. Белая 4','Харьков'),(3,'Василий','Шпак','+38055221166','ул. Северная 5','Житомир');`,
	}

	for _, query := range insertDataQueries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
}
