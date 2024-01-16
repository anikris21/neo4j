package main

import (
	"fmt"
	"context"
    "github.com/neo4j/neo4j-go-driver/v5/neo4j"

)

func main() {
    fmt.Println("Hello, world.")
	ctx := context.Background()
    // URI examples: "neo4j://localhost", "neo4j+s://xxx.databases.neo4j.io"
    dbUri := "neo4j://localhost:7687"
    dbUser := "neo4j"
    dbPassword := "Wal&2021"
    driver, err := neo4j.NewDriverWithContext(
        dbUri,
        neo4j.BasicAuth(dbUser, dbPassword, ""))
    if err != nil {
        panic(err)
    }
    defer driver.Close(ctx)

    err = driver.VerifyConnectivity(ctx)
    if err != nil {
        panic(err)
    }
    fmt.Println("Connection established.")

    result, _ := neo4j.ExecuteQuery(ctx, driver,
        "MATCH (p:App) RETURN p.name AS name",
        nil,
         neo4j.EagerResultTransformer,
        neo4j.ExecuteQueryWithDatabase("autorca"))
    
    // Loop through results and do something with them
    for _, record := range result.Records {
        fmt.Println(record.AsMap())
    }
    
    // Summary information
    fmt.Printf("The query `%v` returned %v records in %+v.\n",
        result.Summary.Query().Text(), len(result.Records),
        result.Summary.ResultAvailableAfter()) 
}