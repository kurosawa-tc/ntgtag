package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Songmu/flextime"
	"github.com/itchyny/timefmt-go"
)

var nowFunc func() time.Time

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	flextime.NowFunc(func() time.Time {
		return time.Now().In(jst)
	})
}

func main() {
	var tagList []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		s := strings.TrimRight(sc.Text(), "\n")
		s = strings.TrimRight(s, "\r")
		tagList = append(tagList, s)
	}
	var tag string
	if len(os.Args) > 1 {
		tag = GenGitTag(tagList, os.Args[1])
	} else {
		tag = GenGitTag(tagList, "")
	}
	fmt.Println(tag)
}

func GenGitTag(tagList []string, branchName string) string {
	tagSearchBaseFmt := `%Y%m%d-(\d+)`
	tagOutputBaseFmt := `%Y%m%d-%%02d.%%s`
	branchSearchFmt := `/(\d+)|(\d+)$`
	var ticketNumberSrc string
	if len(branchName) > 0 {
		ticketNumberSrc = branchName
	}
	now := flextime.Now()
	tagSearchFmt := timefmt.Format(now, tagSearchBaseFmt)
	reTagSearch := regexp.MustCompile(tagSearchFmt)
	reBranch := regexp.MustCompile(branchSearchFmt)
	countNumber := 1

	for _, line := range tagList {
		str := reTagSearch.FindStringSubmatch(line)
		if len(str) == 2 {
			if n, err := strconv.Atoi(strings.TrimLeft(str[1], "0")); err == nil {
				countNumber = n + 1
			}
		}
	}
	var ticketNumber string
	str := reBranch.FindStringSubmatch(ticketNumberSrc)
	if len(str) > 0 && len(str[1]) > 0 {
		ticketNumber = str[1]
	} else if len(str) > 0 && len(str[2]) > 0 {
		ticketNumber = str[2]
	} else {
		ticketNumber = "000"
	}

	outputFmt := timefmt.Format(now, tagOutputBaseFmt)
	return fmt.Sprintf(outputFmt, countNumber, ticketNumber)
}
