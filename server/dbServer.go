package server

import (
	"context"
	"log"
	"prodcat/config"
	"prodcat/ent"
	"prodcat/ent/migrate"
	"prodcat/test/fixtures"

	"github.com/spf13/viper"
)

func InitDatabase(cfg *viper.Viper) *ent.Client {
	connectionString := cfg.GetString("database.connection_string")
	driverName := cfg.GetString("database.driver_name")

	if connectionString == "" {
		log.Fatalf("Database connection string is missing")
	}

	client, err := ent.Open(driverName, connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if config.State == "DEV" && config.RenewDB {
		log.Println("CLEAR DATABASE AND LOAD FIXTURES")
		ClearDB(context.Background(), client)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if config.State == "DEV" && config.RenewDB {
		fixtures.ProductFixtures(context.Background(), client)
	}

	return client
}

func ClearDB(ctx context.Context, c *ent.Client) {
	query := `DROP TABLE public.attribute_value_bools CASCADE;
	DROP TABLE public.attribute_value_nums CASCADE;
	DROP TABLE public.attribute_value_strings CASCADE;
	DROP TABLE public.attribute_variant_nums CASCADE;
	DROP TABLE public.attribute_variant_strings CASCADE;
	DROP TABLE public."attributes" CASCADE;
	DROP TABLE public.products CASCADE;
	DROP TABLE public.user_settings CASCADE;
	DROP TABLE public.users CASCADE;`
	c.ExecContext(ctx, query)
}
