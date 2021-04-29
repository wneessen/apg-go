package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func checkHibp(p string) (bool, error) {
	shaSum := fmt.Sprintf("%x", sha1.Sum([]byte(p)))
	firstPart := shaSum[0:5]
	secondPart := shaSum[5:]
	isPwned := false

	httpClient := &http.Client{Timeout: time.Second * 2}
	httpRes, err := httpClient.Get("https://api.pwnedpasswords.com/range/" + firstPart)
	if err != nil {
		return false, err
	}
	defer func() {
		err := httpRes.Body.Close()
		if err != nil {
			log.Printf("error while closing HTTP response body: %v", err)
		}
	}()

	scanObj := bufio.NewScanner(httpRes.Body)
	for scanObj.Scan() {
		scanLine := strings.SplitN(scanObj.Text(), ":", 2)
		if strings.ToLower(scanLine[0]) == secondPart {
			isPwned = true
			break
		}
	}
	if err := scanObj.Err(); err != nil {
		return isPwned, err
	}

	return isPwned, nil
}
