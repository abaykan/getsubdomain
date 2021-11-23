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
    (_(
    ('')
  _  "\ )>,_     .-->
  _>--w/((_ >,_.'
        ///
        "'"     akbar.kustirama.id	
------------------------------------
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
			fmt.Println("Domain: ", scanner.Text())
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
			fmt.Println(sb)
			os.Exit(0)
		}

		pecah := strings.Split(sb, "<br>")
		pecah = append(pecah, scanner.Text())

		if len(os.Args) < 2 {
			fmt.Printf("Subdomain: " + strconv.Itoa(len(pecah)-1) + "\n\n")
		}

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
		fmt.Println(err)
	}
	location, err := resp.Location()
	if err != nil {
		location = resp.Request.URL
	}

	if strings.Contains(location.String(), domain) {
		fmt.Printf("[%s] %s\n", strconv.Itoa(resp.StatusCode), location)
	} else {
		fmt.Printf("[%s] %s → %s\n", strconv.Itoa(resp.StatusCode), domain, location)
	}
}
