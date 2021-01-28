package postgres

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/shyshlakov/go-http-server/config"
	"github.com/shyshlakov/go-http-server/persistence/model"
	"github.com/shyshlakov/go-http-server/persistence/repo"
)

type postgreRepo struct {
	cfg     *config.Config
	Article []model.Article
	db      *pg.DB
}

func NewRepo(cfg *config.Config) repo.Repo {
	return &postgreRepo{
		cfg: cfg,
	}
}

func (r *postgreRepo) Connect() error {
	dsn, err := r.cfg.GetDSN("postgres")
	if err != nil {
		fmt.Printf("Can not get DSN: %v", err)
		return err
	}
	opt, err := pg.ParseURL(dsn)
	if err != nil {
		fmt.Printf("Can not parse DSN: %v", err)
		return err
	}
	opt.PoolSize = r.cfg.PostgresMaxOpenConns
	opt.ReadTimeout = r.cfg.PostgresMaxConnLifetime
	opt.WriteTimeout = r.cfg.PostgresMaxConnLifetime
	db := pg.Connect(opt)
	r.db = db
	return nil
}

func (r *postgreRepo) Close() error {
	return r.db.Close()
}
