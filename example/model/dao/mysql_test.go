package dao

import (
	"database/sql"
	"fmt"
	"runtime"
	"testing"

	"github.com/extra-time-zone/xgin/example/config"

	_ "github.com/go-sql-driver/mysql"
)

func TestMySqlConnectionDev(t *testing.T) {
	config := &config.MySqlConfig{
		Host:      "54.69.237.139",
		Port:      8096,
		Dbname:    "test",
		Username:  "root",
		Password:  "123456",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}

	conn, err := connectMySql(config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("conn: %+v\n", conn)

	rows, err := conn.Query("select id,name,age from employees where id=?", 2)
	fmt.Printf("err: %+v\n", err)
	for rows.Next() {
		var id, age int
		var name string
		if err := rows.Scan(&id, &name, &age); err != nil {
			panic(err)
		}
		fmt.Printf("id: %d, name:%s, age: %d\n", id, name, age)
	}

	conn.Close()
}

func TestMySqlConnectionTest(t *testing.T) {
	config := &config.MySqlConfig{
		Host:      "localhost",
		Port:      13306,
		Dbname:    "test",
		Username:  "root",
		Password:  "r74pqyYtgdjlYB41jmWA",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}

	conn, err := connectMySql(config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("conn: %+v\n", conn)

	rows, err := conn.Query("select uid, item, expire, itime from bag_0001 where uid=?", 1)
	fmt.Printf("err: %+v\n", err)
	for rows.Next() {
		var uid, itime int
		var item, expire string

		if err := rows.Scan(&uid, &item, &expire, &itime); err != nil {
			panic(err)
		}
		fmt.Printf("uid: %d, item: %s, expire:%v, time: %d\n", uid, item, expire, itime)
	}

	conn.Close()
}

func connectMySql(config *config.MySqlConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s`, config.Username, config.Password, config.Host, config.Port, config.Dbname, config.Charset, config.Collation)
	fmt.Println(dsn)
	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf(`sql.Open error: %w`, err)
	}
	if dbConn == nil {
		return nil, fmt.Errorf(`%v`, "db conn is nil")
	}
	if err = dbConn.Ping(); err != nil {
		return nil, fmt.Errorf(`ping error: %w`, err)
	}

	//Finalizer
	runtime.SetFinalizer(dbConn, func(conn *sql.DB) {
		conn.Close()
	})
	//保证传入的参数在这个方法被调用之前不被垃圾回收器回收掉
	runtime.KeepAlive(dbConn)

	//return
	return dbConn, nil
}
