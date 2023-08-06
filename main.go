package main

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	ctx := context.Background()
	// URI examples: "neo4j://localhost", "neo4j+s://xxx.databases.neo4j.io"
	dbUri := "neo4j://localhost"
	dbUser := "neo4j"
	dbPassword := "1qaz2wsx"
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))

	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MERGE (p:Person {name: $name}) RETURN p",
		map[string]any{
			"name": "Alice",
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created %v nodes in %+v.\n",
		result.Summary.Counters().NodesCreated(),
		result.Summary.ResultAvailableAfter())

	// Get the name of all 42 year-olds
	result, err = neo4j.ExecuteQuery(ctx, driver,
		"MATCH (p:Person WHERE age = $age) RETURN p.name AS name",
		map[string]any{
			"age": "42",
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		log.Fatal(err)
	}
	// Loop through results and do something with them
	for _, record := range result.Records {
		fmt.Println(record.AsMap())
	}

	// Summary information
	fmt.Printf("The query `%v` returned %v records in %+v.\n",
		result.Summary.Query().Text(), len(result.Records),
		result.Summary.ResultAvailableAfter())
}
