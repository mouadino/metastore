package storage

func Init(opts *Options) (DB, error) {
	db, err := opts.DB()
	if err != nil {
		return nil, err
	}
	err = db.Open(opts.DBPath())
	if err != nil {
		return nil, err
	}
	return db, nil
}
