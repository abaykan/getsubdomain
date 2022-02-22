package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func main() {
	banner := `
----------------------------
             .
    ________o=o_______
    akbar.kustirama.id	
----------------------------
	`

	if len(os.Args) < 2 {
		fmt.Println(banner)
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	file := usr.HomeDir + "/.config/c99.txt"
	API_KEY, err := os.ReadFile(file)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if len(os.Args) < 2 {
			fmt.Println("▶ Domain: ", scanner.Text())
		}

		resp, err := http.Get("https://api.c99.nl/subdomainfinder?key=" + strings.TrimRight(string(API_KEY), "\n") + "&domain=" + scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		sb := string(body)

		if strings.Contains(sb, "No subdomains found.") {
			fmt.Printf("No subdomains found.\n\n")
			continue
		}

		pecah := strings.Split(sb, "<br>")
		pecah = append(pecah, scanner.Text())

		if len(os.Args) < 2 {
			fmt.Printf("Subdomain(s): " + strconv.Itoa(len(pecah)-1) + "\n\n")
		}

		fmt.Printf("Checking ...\n")
		for _, subdo := range removeEmptyStrings(pecah) {
			subdo = strings.TrimLeft(subdo, "\r\n")
			getSc(subdo)
		}
		fmt.Printf("\n")
	}
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func getSc(domain string) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	resp, err := client.Get("http://" + domain)
	if err != nil {
		log.Fatalln(err)
	}

	location, err := resp.Location()
	if err != nil {
		location = resp.Request.URL
	}

	if strings.Contains(location.String(), domain) {
		resp2, err := client.Get(location.String())
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("[%s] %s\n", strconv.Itoa(resp2.StatusCode), location.String())
	} else {
		fmt.Printf("[%s] %s → %s\n", strconv.Itoa(resp.StatusCode), domain, location)
	}
}
