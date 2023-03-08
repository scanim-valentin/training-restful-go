package database

import (
	"database/sql"
	"fmt"
	"log"
	"service/utils"
	"time"

	"github.com/lib/pq"
)

// MessageID Unique message identifier
type MessageID int

// MessageContent This could change in the future
type MessageContent string

// Message A message contains the source user, the destination user and a message body
type Message struct {
	ID                  MessageID
	Source, Destination UserID
	Content             MessageContent
	Time                time.Time
}

// Conversation is returned after selecting user to talk to
type Conversation struct {
	Messages []Message
}

// GetMessages retrieves a conversation between two users
func GetMessages(user UserID, other UserID) []Message {
	rows, err := DB.Query("SELECT * FROM messages WHERE source = $1 AND destination = $2 OR source = $2 AND destination = $1 ORDER BY time ASC", user, other)
	if err != nil {
		log.Panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	var messages []Message

	// Reading rows
	for rows.Next() {
		var message Message
		fmt.Print(rows)
		if err := rows.Scan(&message.ID, &message.Source, &message.Destination, &message.Content, &message.Time); err != nil {
			break
		}
		messages = append(messages, message)
	}
	return messages
}

// NewMessage adds a new message to the messages table
func NewMessage(source UserID, destination UserID, content MessageContent, time time.Time) (MessageID, error) {
	var id MessageID
	err := DB.QueryRow("INSERT INTO messages (source, destination, content, time) VALUES ($1, $2, $3, $4) RETURNING id",
		fmt.Sprint(source), fmt.Sprint(destination), content, string(pq.FormatTimestamp(time))).Scan(&id)
	return id, err
}

/**/

func RandomMessage(id MessageID, source UserID, destination UserID, nbMaxChar int) Message {
	return Message{
		ID:          id,
		Source:      source,
		Destination: destination,
		Content:     MessageContent(utils.RandomString(nbMaxChar)),
		Time:        time.Now().UTC().Round(time.Millisecond),
	}
}
