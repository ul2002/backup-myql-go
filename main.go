package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"time"
	"flag"
	"fmt"

)

func main() {
	host := flag.String("host", "0.0.0.0", "Database address")
	port := flag.String("port", "3306", "Port")
	user := flag.String("user", "root", "Username")
	password := flag.String("pass", "root", "Password")
	database := flag.String("db", "", "The name of the database to be divided")
	sqlPath := flag.String("sqlpath", "./backup/", "Backup SQL storage path:")
	flag.Parse()

	if(*database == ""){
      fmt.Println("the database param should not be empty")
	}
	BackupMySqlDb(*host, *port, *user, *password, *database, *sqlPath)
	fmt.Println("Backup is successfully done")
}

/**
 *
 * Backup MySql database
   * @param host: Database address: localhost
   * @param port: Port: 3306
   * @param user: Username: root
   * @param password: password: root
   * @param databaseName: The name of the database to be divided: test
   * @param sqlPath: Backup SQL storage path: ./backup
 * @return 	backupPath
 *
 */
func BackupMySqlDb(host, port, user, password, databaseName, sqlPath string) (error,string)  {
	var cmd *exec.Cmd

	cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return err,""
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return err,""
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
		return err,""
	}
	now := time.Now().Format("20060102150405")
	backupPath := sqlPath+databaseName+"_"+now+".sql"
	
	err = ioutil.WriteFile(backupPath, bytes, 0644)

	if err != nil {
		panic(err)
		return err,""
	}
	return nil,backupPath
}

