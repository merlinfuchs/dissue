package db

import (
	"fmt"
	"strconv"

	"github.com/dgraph-io/badger/v4"
)

func (db *KVDatabase) EnableSync(channelID string, repository string) error {
	return nil
}

func (db *KVDatabase) CreateIssueThreadMapping(issueID int, threadID string) error {
	err := db.kv.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(fmt.Sprintf("mappings:issue:%s", issueID)), []byte(threadID))
		if err != nil {
			return err
		}

		return txn.Set([]byte(fmt.Sprintf("mappings:thread:%s", threadID)), []byte(fmt.Sprintf("%d", issueID)))
	})
	return err
}

func (db *KVDatabase) GetIssueIDForThreadID(threadID string) (int, error) {
	var rawIssueID string
	err := db.kv.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fmt.Sprintf("mappings:thread:%s", threadID)))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			rawIssueID = string(val)
			return nil
		})
		return err
	})

	if err != nil {
		return 0, err
	}

	issueID, err := strconv.Atoi(rawIssueID)
	return issueID, err
}

func (db *KVDatabase) GetThreadIDForIssueID(issueID int) (string, error) {
	var threadID string
	err := db.kv.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fmt.Sprintf("mappings:issue:%d", issueID)))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			threadID = string(val)
			return nil
		})
		return err
	})
	return threadID, err
}
