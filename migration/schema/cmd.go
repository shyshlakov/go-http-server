package main

// import (
// 	"time"

// 	"github.com/golang-migrate/migrate/v4"
// 	"github.com/golang-migrate/migrate/v4/database/postgres"
// 	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/rs/zerolog/log"
// 	"github.com/shyshlakov/go-http-server/config"
// 	"github.com/spf13/cobra"
// )

// var (
// 	migration = &cobra.Command{
// 		Use: "kontomatik-migrate",
// 		RunE: func(_ *cobra.Command, _ []string) error {
// 			log.Info().
// 				Msg("Starting to migrate party...")

// 			binInstance, err := bindata.WithInstance(bindata.Resource(AssetNames(), Asset))
// 			if err != nil {
// 				log.Error().Err(err).Msg("failed to init db instance")
// 				return err
// 			}
// 			cfg, err := config.FromEnv()
// 			if err != nil {
// 				return err
// 			}
// 			dsn, err := cfg.GetDSN("postgres")
// 			if err != nil {
// 				return err
// 			}

// 			var conn *sqlx.DB
// 			//nolint:gomnd
// 			for tries := 0; tries <= 20; tries++ {
// 				conn, err = sqlx.Connect("postgres", dsn)
// 				if err != nil {
// 					log.Error().Err(err).Msg("failed to connect to DB. Retying in 2 seconds")
// 					time.Sleep(2000 * time.Millisecond)
// 					if tries == 20 {
// 						return err
// 					}
// 					continue
// 				}
// 				break
// 			}
// 			defer func() {
// 				if err = conn.Close(); err != nil {
// 					log.Error().Err(err).Msg("failed to close postgres connection")
// 				}
// 			}()

// 			targetInstance, err := postgres.WithInstance(conn.DB, new(postgres.Config))
// 			if err != nil {
// 				return err
// 			}
// 			m, err := migrate.NewWithInstance("go-bindata", binInstance, "postgres", targetInstance)
// 			if err != nil {
// 				return err
// 			}
// 			err = m.Migrate(cfg.MigrationVersion)
// 			if err != nil && err == migrate.ErrNoChange {
// 				log.Debug().Msg("No new migrations found")
// 				return nil
// 			} else if err != nil {
// 				log.Error().Err(err)
// 				return err
// 			} else {
// 				log.Debug().Msgf("Migrations to revision %d run.", cfg.MigrationVersion)
// 				return nil
// 			}
// 		},
// 	}
// )

// func Factory() *cobra.Command { return migration }
