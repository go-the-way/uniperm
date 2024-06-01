// Copyright 2024 uniperm Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	_ "embed"

	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	logger      = log.New(os.Stdout, "[permctl] ", log.LstdFlags|log.Lshortfile)
	generateCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen", "g"},
		Short:   "generate permission data",
		Long:    "generate permission data to db from json file",
		Run:     generateRun,
	}
	table             = "uniperm_permissions"
	dbHost            = "127.0.0.1"
	dbPort            = "22"
	dbUser            = "root"
	dbPasswd          = "passwd"
	dbName            = "uniperm"
	dbDSNFromEnv      = false
	dbDSNEnvName      = "PERMCTL_DSN"
	jsonFile          = "perm.json"
	truncate          = false
	create            = false
	skipEmptyName     = true
	verbose           = false
	printInsertSqlTpl = false

	db *sql.DB
)

func init() {
	generateCmd.PersistentFlags().StringVarP(&table, "table", "t", table, "The uniperm permission table name")
	generateCmd.PersistentFlags().StringVarP(&dbHost, "dbHost", "H", dbHost, "The database host")
	generateCmd.PersistentFlags().StringVarP(&dbPort, "dbPort", "p", dbPort, "The database port")
	generateCmd.PersistentFlags().StringVarP(&dbUser, "dbUser", "u", dbUser, "The database user")
	generateCmd.PersistentFlags().StringVarP(&dbPasswd, "dbPasswd", "P", dbPasswd, "The database password")
	generateCmd.PersistentFlags().StringVarP(&dbName, "dbName", "d", dbName, "The database name")
	generateCmd.PersistentFlags().BoolVarP(&dbDSNFromEnv, "dbDSNFromEnv", "e", dbDSNFromEnv, "The database dns read from env")
	generateCmd.PersistentFlags().StringVarP(&dbDSNEnvName, "dbDSNEnvName", "E", dbDSNEnvName, "The database dns env variable name")
	generateCmd.PersistentFlags().StringVarP(&jsonFile, "jsonFile", "j", jsonFile, "The perm json file")
	generateCmd.PersistentFlags().BoolVarP(&create, "create", "c", create, "Create uniperm permission table?")
	generateCmd.PersistentFlags().BoolVarP(&truncate, "truncate", "C", truncate, "Truncate uniperm permission table?")
	generateCmd.PersistentFlags().BoolVarP(&skipEmptyName, "skipEmptyName", "s", skipEmptyName, "Skip empty perm name?")
	generateCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", verbose, "Show verbose info")
	generateCmd.PersistentFlags().BoolVar(&printInsertSqlTpl, "printInsertSqlTpl", printInsertSqlTpl, "Print insert sql template, then exit")

	rootCmd.AddCommand(generateCmd)
}

func generateRun(_ *cobra.Command, _ []string) {
	if printInsertSqlTpl {
		logShow(insertSql())
		return
	}
	check()
	initDb()
	createTable()
	truncateTable()
	insert()
	logShow("finish done.")
}

func check() {
	if _, err := os.ReadFile(jsonFile); err != nil {
		logShow("read json file err:", err)
		os.Exit(1)
	}
}

func initDb() {
	dbDSN := getDbDSN()
	if dbDSNFromEnv {
		dbDSN = os.Getenv(dbDSNEnvName)
	}
	db, _ = sql.Open("mysql", dbDSN)
	if verbose {
		verboseShow("current db dsn:", dbDSN)
	}
	if err := db.Ping(); err != nil {
		logShow("ping db err:", err)
		os.Exit(1)
	}
}

var (
	//go:embed files/ddl1.txt
	sql1Fs string
	//go:embed files/ddl2.txt
	sql2Fs string
)

func createTable() {
	if create {
		if _, err := db.Exec(sql1Fs); err != nil {
			logShow("create table err:", err)
			os.Exit(1)
		}
		if _, err := db.Exec(sql2Fs); err != nil {
			logShow("create index err:", err)
			os.Exit(1)
		}
	}
}

func truncateTable() {
	if !create && truncate {
		if _, err := db.Exec("truncate table " + table); err != nil {
			logShow("truncate table err:", err)
			os.Exit(1)
		}
	}
}

func getDbDSN() string {
	var dbParam = "charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbUser, dbPasswd, dbHost, dbPort, dbName, dbParam)
}

func insert() {
	buf, _ := os.ReadFile(jsonFile)
	var ps []perm
	if err := json.Unmarshal(buf, &ps); err != nil {
		logShow("unmarshal json err:", err)
		return
	}
	if len(ps) > 0 {
		for _, pm := range ps {
			if !(skipEmptyName && pm.ignore()) {
				if err := treeInsert(pm); err != nil {
					logShow("insert one err:", err)
					return
				}
			}
		}
	}
}

func treeInsert(pm perm) (err error) {
	var (
		result     sql.Result
		insertedId int64
	)

	sqlStr := getInsertSqlFromPerm(pm, pm.ParentId, pm.IsButton)
	if result, err = db.Exec(sqlStr); err != nil {
		return
	}

	verboseShow(sqlStr)

	if insertedId, err = result.LastInsertId(); err != nil {
		return
	}

	verboseShow("last inserted id:", insertedId)

	if len(pm.Routes) > 0 {
		for _, pmm := range pm.Routes {
			if !(skipEmptyName && pm.ignore()) {
				pmm.ParentId = insertedId
				if err = treeInsert(pmm); err != nil {
					return
				}
			}
		}
	}

	return
}

func logShow(v ...any) {
	logger.Println(v...)
}

func verboseShow(v ...any) {
	if verbose {
		logger.Println(v...)
	}
}
func insertSql() string {
	return "insert into " + table + "(name, route, parent_id, is_button) values ('%s', '%s', %d, %d)"
}

func getInsertSql(name, route string, parentId int64, isButton bool) string {
	// 是否按钮 1是 2否
	m := map[bool]byte{true: 1, false: 2}
	return fmt.Sprintf(insertSql(), name, route, parentId, m[isButton])
}

func getInsertSqlFromPerm(perm perm, parentId int64, isButton bool) string {
	return getInsertSql(perm.Name, perm.Path, parentId, isButton)
}
