package database

import (
	"database/sql"
	"log"
	"service/utils"
	"time"
)

// MessageID Unique message identifier
type MessageID int

// MessageContent This could change in the future
type MessageContent string

// Message A message contains the source users, the destination users and a message body
type Message struct {
	ID      MessageID
	User    UserID
	Group   GroupID
	Content MessageContent
	Time    time.Time
}

// GetMessages retrieves a conversation between two users
func GetMessages(other GroupID) ([]Message, error) {
	// TODO handle blocked users (censored message content)
	rows, err := DB.Query("SELECT * FROM messages WHERE groupid = $1 ORDER BY time ASC", other)
	if err != nil {
		log.Panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Panic("Failed to close rows in GetMessage: ", err)
		}
	}(rows)
	var messages []Message

	// Reading rows
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.ID, &message.User, &message.Group, &message.Content, &message.Time); err != nil {
			break
		}
		messages = append(messages, message)
	}
	if err = rows.Err(); err != nil {
		log.Println("Error in rows: ", err)
	}
	return messages, err
}

// NewMessage adds a new message to the messages table
func NewMessage(message Message) (MessageID, error) {
	var id MessageID
	err := DB.QueryRow("INSERT INTO messages (userid, groupid, content, time) VALUES ($1, $2, $3, $4) RETURNING id",
		message.User, message.Group, message.Content, message.Time).Scan(&id)
	return id, err
}

func DeleteMessage(messageID MessageID) error {
	_, err := DB.Exec("DELETE FROM messages WHERE id=$1", messageID)
	return err
}

/**/

func RandomMessage(id MessageID, source UserID, destination GroupID, nbMaxChar int) Message {
	return Message{
		ID:      id,
		User:    source,
		Group:   destination,
		Content: MessageContent(utils.RandomString(nbMaxChar)),
		Time:    time.Now().UTC().Round(time.Millisecond),
	}
}
