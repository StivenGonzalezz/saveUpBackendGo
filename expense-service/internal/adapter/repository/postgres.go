package repository

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo() ports.ExpenseRepository{
	dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}
	
	db.AutoMigrate(&model.expense{})

	return &PostgresRepo{db: db}
}

func (p *PostgresRepo) CreateExpense(expense *model.expense) error {
	return p.db.Create(expense).Error
}

func (p *PostgresRepo) DeleteExpense(id uint) error {
	var expense model.expense
	if err := p.db.First(&expense, id).Error; err != nil {
		return err
	}
	return p.db.Delete(&expense).Error
}

func (p *PostgresRepo) GetExpensesByFinanceId(financeId uint) ([]model.expense, error) {
	var expenses []model.expense
	if err := p.db.Where("finance_id = ?", financeId).Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}
	