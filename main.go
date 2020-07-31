package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type user map[string]map[string]string
type table map[string]map[string]map[string]string
type data map[string][]map[string]string

func main() {
	res := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	users := user{}
	tables := table{}
	data := data{}
	for scanner.Scan() && scanner.Text() != "done" {
		command := scanner.Text()
		c1, _ := regexp.MatchString("create user*", command)
		c2, _ := regexp.MatchString("create table*", command)
		c3, _ := regexp.MatchString("delete table*", command)
		c4, _ := regexp.MatchString("add column*", command)
		c5, _ := regexp.MatchString("remove column*", command)
		c6, _ := regexp.MatchString("add row*", command)
		c7, _ := regexp.MatchString("remove row*", command)
		c8, _ := regexp.MatchString("change*", command)
		c9, _ := regexp.MatchString("print*", command)
		c10, _ := regexp.MatchString("search*", command)
		switch {
		case c1:
			slices := strings.Split(command, " ")
			users[slices[2]] = map[string]string{}
			users[slices[2]]["rule"] = slices[3]
		case c2:
			// add table
			slices := strings.Split(command, " ")
			if users[slices[3]]["rule"] == "editor" {
				tables[slices[2]] = map[string]map[string]string{}
			} else {
				res = append(res, "access denied")
			}
		case c3:
			// delete table
			slices := strings.Split(command, " ")
			if users[slices[3]]["rule"] == "editor" {
				delete(tables, slices[2])
			} else {
				res = append(res, "access denied")
			}
		case c4:
			// add col
			slices := strings.Split(command, " ")
			if users[slices[5]]["rule"] == "editor" {
				tables[slices[2]] = map[string]map[string]string{}
				tables[slices[2]][slices[3]] = map[string]string{}
				tables[slices[2]][slices[3]]["type"] = slices[4]
				for k := range data[slices[2]][len(data[slices[2]])-1] {
					if tables[slices[2]][k]["type"] == "int" {
						data[slices[2]][len(data[slices[2]])-1][k] = "0"
					} else {
						data[slices[2]][len(data[slices[2]])-1][k] = "null"
					}
				}
			} else {
				res = append(res, "access denied")
			}

		case c5:
			//remove col
			slices := strings.Split(command, " ")
			if users[slices[4]]["rule"] == "editor" {
				delete(tables[slices[2]], slices[1])
			} else {
				res = append(res, "access denied")
			}
		case c6:
			// add row
			slices := strings.Split(command, " ")
			if users[slices[3]]["rule"] == "editor" {
				data[slices[2]] = append(data[slices[2]], map[string]string{})
			} else {
				res = append(res, "access denied")
			}
		case c7:
			// remove row
			slices := strings.Split(command, " ")
			if users[slices[4]]["rule"] == "editor" {
				i, _ := strconv.Atoi(slices[3])
				data[slices[2]][i-1] = map[string]string{}
			} else {
				res = append(res, "access denied")
			}
		case c8:
			// change
			slices := strings.Split(command, " ")
			if users[slices[5]]["rule"] == "editor" {
				i, _ := strconv.Atoi(slices[2])
				data[slices[1]][i-1][slices[3]] = slices[4]
			} else {
				res = append(res, "access denied")
			}
		case c9:
			//print
			slices := strings.Split(command, " ")
			_ = slices
		case c10:
			//search
			slices := strings.Split(command, " ")
			str := ""
			for _, v := range data[slices[1]] {
				if v[slices[2]] == slices[3] {
					for _, value := range v {
						str += value + " "
					}
				}
			}
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				res = append(res, str)
			}
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

func makeZero(d data, t table) data {
	for tName, v := range d {
		for kk, vv := range v {
			for key := range vv {
				if t[tName][key]["type"] == "int" && len(d[tName][kk][key]) == 0 {
					d[tName][kk][key] = "0"
				}
			}
		}
	}
	return d
}
