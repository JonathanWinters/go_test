package database

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID
}

/*
// creates initial user
func CreateUser(ctx context.Context, user *User) (err error) {
	err = DB.ExecOne(ctx,
		`INSERT INTO
		users (id, roundid)
		VALUES ($1, $2)
		ON CONFLICT (id) DO
		UPDATE SET roundid=$2`,
		user.ID, user.CurrentRoundID)
	return
}

// gets existing user
func GetUserByID(ctx context.Context, id pb.UserID) (user *User, err error) {
	user = &User{}
	err = DB.QueryRow(ctx,
		`SELECT id, roundid
		FROM users
		WHERE id=$1`,
		id).Scan(&user.ID, &user.CurrentRoundID)
	return
}

// updates user's current round ID, after wager accepted by pb
func UpdateUserCurrentRound(ctx context.Context, tx *postgresqldb.Transaction, userID pb.UserID, currentRoundID *pb.RoundID) (err error) {
	err = tx.ExecOne(ctx,
		`UPDATE users
	 	SET roundid=$2
	 	WHERE id=$1`,
		userID, currentRoundID)
	return
}
*/
