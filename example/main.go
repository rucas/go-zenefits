package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/rucas/zenefits/zenefits"
	"golang.org/x/oauth2"
)

func main() {
	id := flag.Int("id", 0, "Zenefits Company Id")
	k := flag.String("key", "", "Zenefits Api Key")
	flag.Parse()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *k})
	tc := oauth2.NewClient(ctx, ts)
	client := zenefits.NewClient(tc)
	opt := &zenefits.PeopleQueryParams{Limit: 10}

	var allPeople []*zenefits.People
	for {
		p, resp, err := client.People.List(ctx, *id, opt)
		if err != nil {
			fmt.Printf("Error, %v \n", err)
			break
		}
		allPeople = append(allPeople, p...)
		if resp.NextPage == 0 {
			break
		}
		opt.StartingAfter = resp.NextPage
	}

	fmt.Printf("Done...Number of Employees: %v \n", len(allPeople))
}
