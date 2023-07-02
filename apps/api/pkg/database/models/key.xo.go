package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// Key represents a row from 'unkey.keys'.
type Key struct {
	ID                      string         `json:"id"`                        // id
	APIID                   string         `json:"api_id"`                    // api_id
	Hash                    string         `json:"hash"`                      // hash
	Start                   string         `json:"start"`                     // start
	OwnerID                 sql.NullString `json:"owner_id"`                  // owner_id
	Meta                    sql.NullString `json:"meta"`                      // meta
	CreatedAt               time.Time      `json:"created_at"`                // created_at
	Expires                 sql.NullTime   `json:"expires"`                   // expires
	RatelimitType           sql.NullString `json:"ratelimit_type"`            // ratelimit_type
	RatelimitLimit          sql.NullInt64  `json:"ratelimit_limit"`           // ratelimit_limit
	RatelimitRefillRate     sql.NullInt64  `json:"ratelimit_refill_rate"`     // ratelimit_refill_rate
	RatelimitRefillInterval sql.NullInt64  `json:"ratelimit_refill_interval"` // ratelimit_refill_interval
	WorkspaceID             string         `json:"workspace_id"`              // workspace_id
	ForWorkspaceID          sql.NullString `json:"for_workspace_id"`          // for_workspace_id
	Name                    sql.NullString `json:"name"`                      // name
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Key] exists in the database.
func (k *Key) Exists() bool {
	return k._exists
}

// Deleted returns true when the [Key] has been marked for deletion
// from the database.
func (k *Key) Deleted() bool {
	return k._deleted
}

// Insert inserts the [Key] to the database.
func (k *Key) Insert(ctx context.Context, db DB) error {
	switch {
	case k._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case k._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO unkey.keys (` +
		`id, api_id, hash, start, owner_id, meta, created_at, expires, ratelimit_type, ratelimit_limit, ratelimit_refill_rate, ratelimit_refill_interval, workspace_id, for_workspace_id, name` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, k.ID, k.APIID, k.Hash, k.Start, k.OwnerID, k.Meta, k.CreatedAt, k.Expires, k.RatelimitType, k.RatelimitLimit, k.RatelimitRefillRate, k.RatelimitRefillInterval, k.WorkspaceID, k.ForWorkspaceID, k.Name)
	if _, err := db.ExecContext(ctx, sqlstr, k.ID, k.APIID, k.Hash, k.Start, k.OwnerID, k.Meta, k.CreatedAt, k.Expires, k.RatelimitType, k.RatelimitLimit, k.RatelimitRefillRate, k.RatelimitRefillInterval, k.WorkspaceID, k.ForWorkspaceID, k.Name); err != nil {
		return logerror(err)
	}
	// set exists
	k._exists = true
	return nil
}

// Update updates a [Key] in the database.
func (k *Key) Update(ctx context.Context, db DB) error {
	switch {
	case !k._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case k._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE unkey.keys SET ` +
		`api_id = ?, hash = ?, start = ?, owner_id = ?, meta = ?, created_at = ?, expires = ?, ratelimit_type = ?, ratelimit_limit = ?, ratelimit_refill_rate = ?, ratelimit_refill_interval = ?, workspace_id = ?, for_workspace_id = ?, name = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, k.APIID, k.Hash, k.Start, k.OwnerID, k.Meta, k.CreatedAt, k.Expires, k.RatelimitType, k.RatelimitLimit, k.RatelimitRefillRate, k.RatelimitRefillInterval, k.WorkspaceID, k.ForWorkspaceID, k.Name, k.ID)
	if _, err := db.ExecContext(ctx, sqlstr, k.APIID, k.Hash, k.Start, k.OwnerID, k.Meta, k.CreatedAt, k.Expires, k.RatelimitType, k.RatelimitLimit, k.RatelimitRefillRate, k.RatelimitRefillInterval, k.WorkspaceID, k.ForWorkspaceID, k.Name, k.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Key] to the database.
func (k *Key) Save(ctx context.Context, db DB) error {
	if k.Exists() {
		return k.Update(ctx, db)
	}
	return k.Insert(ctx, db)
}

// Upsert performs an upsert for [Key].
func (k *Key) Upsert(ctx context.Context, db DB) error {
	switch {
	case k._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO unkey.keys (` +
		`id, api_id, hash, start, owner_id, meta, created_at, expires, ratelimit_type, ratelimit_limit, ratelimit_refill_rate, ratelimit_refill_interval, workspace_id, for_workspace_id, name` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`id = VALUES(id), api_id = VALUES(api_id), hash = VALUES(hash), start = VALUES(start), owner_id = VALUES(owner_id), meta = VALUES(meta), created_at = VALUES(created_at), expires = VALUES(expires), ratelimit_type = VALUES(ratelimit_type), ratelimit_limit = VALUES(ratelimit_limit), ratelimit_refill_rate = VALUES(ratelimit_refill_rate), ratelimit_refill_interval = VALUES(ratelimit_refill_interval), workspace_id = VALUES(workspace_id), for_workspace_id = VALUES(for_workspace_id), name = VALUES(name)`
	// run
	logf(sqlstr, k.ID, k.APIID, k.Hash, k.Start, k.OwnerID, k.Meta, k.CreatedAt, k.Expires, k.RatelimitType, k.RatelimitLimit, k.RatelimitRefillRate, k.RatelimitRefillInterval, k.WorkspaceID, k.ForWorkspaceID, k.Name)
	if _, err := db.ExecContext(ctx, sqlstr, k.ID, k.APIID, k.Hash, k.Start, k.OwnerID, k.Meta, k.CreatedAt, k.Expires, k.RatelimitType, k.RatelimitLimit, k.RatelimitRefillRate, k.RatelimitRefillInterval, k.WorkspaceID, k.ForWorkspaceID, k.Name); err != nil {
		return logerror(err)
	}
	// set exists
	k._exists = true
	return nil
}

// Delete deletes the [Key] from the database.
func (k *Key) Delete(ctx context.Context, db DB) error {
	switch {
	case !k._exists: // doesn't exist
		return nil
	case k._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM unkey.keys ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, k.ID)
	if _, err := db.ExecContext(ctx, sqlstr, k.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	k._deleted = true
	return nil
}

// KeyByHash retrieves a row from 'unkey.keys' as a [Key].
//
// Generated from index 'hash_idx'.
func KeyByHash(ctx context.Context, db DB, hash string) (*Key, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, api_id, hash, start, owner_id, meta, created_at, expires, ratelimit_type, ratelimit_limit, ratelimit_refill_rate, ratelimit_refill_interval, workspace_id, for_workspace_id, name ` +
		`FROM unkey.keys ` +
		`WHERE hash = ?`
	// run
	logf(sqlstr, hash)
	k := Key{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, hash).Scan(&k.ID, &k.APIID, &k.Hash, &k.Start, &k.OwnerID, &k.Meta, &k.CreatedAt, &k.Expires, &k.RatelimitType, &k.RatelimitLimit, &k.RatelimitRefillRate, &k.RatelimitRefillInterval, &k.WorkspaceID, &k.ForWorkspaceID, &k.Name); err != nil {
		return nil, logerror(err)
	}
	return &k, nil
}

// KeyByID retrieves a row from 'unkey.keys' as a [Key].
//
// Generated from index 'keys_id_pkey'.
func KeyByID(ctx context.Context, db DB, id string) (*Key, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, api_id, hash, start, owner_id, meta, created_at, expires, ratelimit_type, ratelimit_limit, ratelimit_refill_rate, ratelimit_refill_interval, workspace_id, for_workspace_id, name ` +
		`FROM unkey.keys ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	k := Key{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&k.ID, &k.APIID, &k.Hash, &k.Start, &k.OwnerID, &k.Meta, &k.CreatedAt, &k.Expires, &k.RatelimitType, &k.RatelimitLimit, &k.RatelimitRefillRate, &k.RatelimitRefillInterval, &k.WorkspaceID, &k.ForWorkspaceID, &k.Name); err != nil {
		return nil, logerror(err)
	}
	return &k, nil
}