package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DeLuci/inventory-system/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertUser inserts a reservation into the database
func (m *postgresDBRepo) InsertUser(res models.User) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	// m.DB.ExecContext(ctx, smt, ...)

	stmt := `insert into users (email, email_confirmation, password, access_level, created_at, updated_at) 
			values ($1, $2, $3, $4, $5, $6)`
	_, err := m.DB.Exec(stmt,
		res.Email,
		res.EmailConfirmation,
		res.Password,
		res.AccessLevel,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

// FindUser find the user in the database
func (m *postgresDBRepo) Authenticate(email, password string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var hashedPwd string
	row := m.DB.QueryRowContext(ctx, "select id, password from users where email = $1", email)
	err := row.Scan(&id, &hashedPwd)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPwd, nil
}

// InsertFinishedProduct one time function to insert the products in the DB
func (m *postgresDBRepo) InsertFinishedProduct(res models.Product) error {
	return nil
}

// InsertNewProduct adding any new product in the db
func (m *postgresDBRepo) InsertNewProduct(res models.ScanProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `Update sizes set $1 = $1 + 1 where id =$2`
	_, err := m.DB.QueryContext(ctx, query, res.Size, res.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) SearchProduct(searchName string) ([]models.SearchBoot, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var items []models.SearchBoot
	rows, err := m.DB.QueryContext(ctx, `select  p.description, s."24", s."24.5" , s."25" , s."25.5" , s."26" ,
       s."26.5" , s."27" , s."27.5" ,
	    s."28", s."28.5" , s."29" , s."29.5" , s."30" , s."30.5", s."31" from sizes as s
		inner join products as p
		on s.id = p.id where to_tsvector(p.description) @@ plainto_tsquery($1) `, searchName)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	for rows.Next() {
		var item models.SearchBoot
		err = rows.Scan(&item.Description, &item.Size24, &item.Size24and5, &item.Size25, &item.Size25and5, &item.Size26,
			&item.Size26and5, &item.Size27, &item.Size27and5, &item.Size28, &item.Size28and5, &item.Size29,
			&item.Size29and5, &item.Size30, &item.Size30and5, &item.Size31)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
