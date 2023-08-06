package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	ctx := context.Background()
	var employeeThreshold int64 = 10
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

	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// Create 100 people and assign them to various organizations
	for i := 0; i < 100; i++ {
		name := "Thor" + strconv.Itoa(i)
		orgId, err := session.ExecuteWrite(ctx,
			func(tx neo4j.ManagedTransaction) (any, error) {
				var orgId string

				// Create new Person node with given name, if not exists already
				_, err := tx.Run(
					ctx,
					"MERGE (p:Person {name: $name})",
					map[string]any{
						"name": name,
					})
				if err != nil {
					return nil, err
				}

				// Obtain most recent organization ID and the number of people linked to it
				result, err := tx.Run(
					ctx, `
                    MATCH (o:Organization)
                    RETURN o.id AS id, COUNT{(p:Person)-[r:WORKS_FOR]->(o)} AS employeesN
                    ORDER BY o.createdDate DESC
                    LIMIT 1
                    `, nil)
				if err != nil {
					return nil, err
				}
				org, err := result.Single(ctx)

				// If no organization exists, create one and add Person to it
				if org == nil {
					orgId, _ = createOrganization(ctx, tx)
					fmt.Println("No orgs available, created", orgId)
					err = addPersonToOrganization(ctx, tx, name, orgId)
					if err != nil {
						return nil, errors.New("Failed to add person to new org")
						// Transaction will roll back
						// -> not even Person and/or Organization is created!
					}
				} else {
					orgId = org.AsMap()["id"].(string)
					if employeesN := org.AsMap()["employeesN"].(int64); employeesN == 0 {
						return nil, errors.New("Most recent organization is empty")
						// Transaction will roll back
						// -> not even Person is created!
					}

					// If org does not have too many employees, add this Person to it
					if employeesN := org.AsMap()["employeesN"].(int64); employeesN < employeeThreshold {
						err = addPersonToOrganization(ctx, tx, name, orgId)
						if err != nil {
							return nil, err
							// Transaction will roll back
							// -> not even Person is created!
						}
						// Otherwise, create a new Organization and link Person to it
					} else {
						orgId, err = createOrganization(ctx, tx)
						if err != nil {
							return nil, err
							// Transaction will roll back
							// -> not even Person is created!
						}
						fmt.Println("Latest org is full, created", orgId)
						err = addPersonToOrganization(ctx, tx, name, orgId)
						if err != nil {
							return nil, err
							// Transaction will roll back
							// -> not even Person and/or Organization is created!
						}
					}
				}
				// Return the Organization ID to which the new Person ends up in
				return orgId, nil
			})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("User", name, "added to organization", orgId)
		}
	}
}

func createOrganization(ctx context.Context, tx neo4j.ManagedTransaction) (string, error) {
	result, err := tx.Run(
		ctx, `
        CREATE (o:Organization {id: randomuuid(), createdDate: datetime()})
        RETURN o.id AS id
        `, nil)
	if err != nil {
		return "", err
	}
	org, err := result.Single(ctx)
	if err != nil {
		return "", err
	}
	orgId, _ := org.AsMap()["id"]
	return orgId.(string), err
}

func addPersonToOrganization(ctx context.Context, tx neo4j.ManagedTransaction, personName string, orgId string) error {
	_, err := tx.Run(
		ctx, `
        MATCH (o:Organization {id: $orgId})
        MATCH (p:Person {name: $name})
        MERGE (p)-[:WORKS_FOR]->(o)
        `, map[string]any{
			"orgId": orgId,
			"name":  personName,
		})
	return err
}
