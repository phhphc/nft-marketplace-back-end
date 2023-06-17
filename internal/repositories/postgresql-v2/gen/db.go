// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package gen

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.deleteProfileStmt, err = db.PrepareContext(ctx, deleteProfile); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProfile: %w", err)
	}
	if q.deleteUserRoleStmt, err = db.PrepareContext(ctx, deleteUserRole); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserRole: %w", err)
	}
	if q.fullTextSearchStmt, err = db.PrepareContext(ctx, fullTextSearch); err != nil {
		return nil, fmt.Errorf("error preparing query FullTextSearch: %w", err)
	}
	if q.getAllRolesStmt, err = db.PrepareContext(ctx, getAllRoles); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllRoles: %w", err)
	}
	if q.getCategoryByNameStmt, err = db.PrepareContext(ctx, getCategoryByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetCategoryByName: %w", err)
	}
	if q.getCollectionStmt, err = db.PrepareContext(ctx, getCollection); err != nil {
		return nil, fmt.Errorf("error preparing query GetCollection: %w", err)
	}
	if q.getCollectionLastSyncBlockStmt, err = db.PrepareContext(ctx, getCollectionLastSyncBlock); err != nil {
		return nil, fmt.Errorf("error preparing query GetCollectionLastSyncBlock: %w", err)
	}
	if q.getEventStmt, err = db.PrepareContext(ctx, getEvent); err != nil {
		return nil, fmt.Errorf("error preparing query GetEvent: %w", err)
	}
	if q.getExpiredOrderStmt, err = db.PrepareContext(ctx, getExpiredOrder); err != nil {
		return nil, fmt.Errorf("error preparing query GetExpiredOrder: %w", err)
	}
	if q.getMarketplaceLastSyncBlockStmt, err = db.PrepareContext(ctx, getMarketplaceLastSyncBlock); err != nil {
		return nil, fmt.Errorf("error preparing query GetMarketplaceLastSyncBlock: %w", err)
	}
	if q.getMarketplaceSettingsStmt, err = db.PrepareContext(ctx, getMarketplaceSettings); err != nil {
		return nil, fmt.Errorf("error preparing query GetMarketplaceSettings: %w", err)
	}
	if q.getNftStmt, err = db.PrepareContext(ctx, getNft); err != nil {
		return nil, fmt.Errorf("error preparing query GetNft: %w", err)
	}
	if q.getNotificationStmt, err = db.PrepareContext(ctx, getNotification); err != nil {
		return nil, fmt.Errorf("error preparing query GetNotification: %w", err)
	}
	if q.getOfferStmt, err = db.PrepareContext(ctx, getOffer); err != nil {
		return nil, fmt.Errorf("error preparing query GetOffer: %w", err)
	}
	if q.getOrderStmt, err = db.PrepareContext(ctx, getOrder); err != nil {
		return nil, fmt.Errorf("error preparing query GetOrder: %w", err)
	}
	if q.getProfileStmt, err = db.PrepareContext(ctx, getProfile); err != nil {
		return nil, fmt.Errorf("error preparing query GetProfile: %w", err)
	}
	if q.getUserByAddressStmt, err = db.PrepareContext(ctx, getUserByAddress); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByAddress: %w", err)
	}
	if q.getUsersStmt, err = db.PrepareContext(ctx, getUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsers: %w", err)
	}
	if q.getValidMarketplaceSettingsStmt, err = db.PrepareContext(ctx, getValidMarketplaceSettings); err != nil {
		return nil, fmt.Errorf("error preparing query GetValidMarketplaceSettings: %w", err)
	}
	if q.insertCategoryStmt, err = db.PrepareContext(ctx, insertCategory); err != nil {
		return nil, fmt.Errorf("error preparing query InsertCategory: %w", err)
	}
	if q.insertCollectionStmt, err = db.PrepareContext(ctx, insertCollection); err != nil {
		return nil, fmt.Errorf("error preparing query InsertCollection: %w", err)
	}
	if q.insertEventStmt, err = db.PrepareContext(ctx, insertEvent); err != nil {
		return nil, fmt.Errorf("error preparing query InsertEvent: %w", err)
	}
	if q.insertMarketplaceSettingsStmt, err = db.PrepareContext(ctx, insertMarketplaceSettings); err != nil {
		return nil, fmt.Errorf("error preparing query InsertMarketplaceSettings: %w", err)
	}
	if q.insertNotificationStmt, err = db.PrepareContext(ctx, insertNotification); err != nil {
		return nil, fmt.Errorf("error preparing query InsertNotification: %w", err)
	}
	if q.insertOrderStmt, err = db.PrepareContext(ctx, insertOrder); err != nil {
		return nil, fmt.Errorf("error preparing query InsertOrder: %w", err)
	}
	if q.insertOrderConsiderationItemStmt, err = db.PrepareContext(ctx, insertOrderConsiderationItem); err != nil {
		return nil, fmt.Errorf("error preparing query InsertOrderConsiderationItem: %w", err)
	}
	if q.insertOrderOfferItemStmt, err = db.PrepareContext(ctx, insertOrderOfferItem); err != nil {
		return nil, fmt.Errorf("error preparing query InsertOrderOfferItem: %w", err)
	}
	if q.insertUserStmt, err = db.PrepareContext(ctx, insertUser); err != nil {
		return nil, fmt.Errorf("error preparing query InsertUser: %w", err)
	}
	if q.insertUserRoleStmt, err = db.PrepareContext(ctx, insertUserRole); err != nil {
		return nil, fmt.Errorf("error preparing query InsertUserRole: %w", err)
	}
	if q.listNftWithListingStmt, err = db.PrepareContext(ctx, listNftWithListing); err != nil {
		return nil, fmt.Errorf("error preparing query ListNftWithListing: %w", err)
	}
	if q.updateCollectionLastSyncBlockStmt, err = db.PrepareContext(ctx, updateCollectionLastSyncBlock); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCollectionLastSyncBlock: %w", err)
	}
	if q.updateMarketplaceLastSyncBlockStmt, err = db.PrepareContext(ctx, updateMarketplaceLastSyncBlock); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMarketplaceLastSyncBlock: %w", err)
	}
	if q.updateNftStmt, err = db.PrepareContext(ctx, updateNft); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateNft: %w", err)
	}
	if q.updateNonceStmt, err = db.PrepareContext(ctx, updateNonce); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateNonce: %w", err)
	}
	if q.updateNotificationStmt, err = db.PrepareContext(ctx, updateNotification); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateNotification: %w", err)
	}
	if q.updateOrderStatusStmt, err = db.PrepareContext(ctx, updateOrderStatus); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateOrderStatus: %w", err)
	}
	if q.updateOrderStatusByOfferStmt, err = db.PrepareContext(ctx, updateOrderStatusByOffer); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateOrderStatusByOffer: %w", err)
	}
	if q.updateUserBlockStateStmt, err = db.PrepareContext(ctx, updateUserBlockState); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserBlockState: %w", err)
	}
	if q.upsertNftLatestStmt, err = db.PrepareContext(ctx, upsertNftLatest); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertNftLatest: %w", err)
	}
	if q.upsertProfileStmt, err = db.PrepareContext(ctx, upsertProfile); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertProfile: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.deleteProfileStmt != nil {
		if cerr := q.deleteProfileStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProfileStmt: %w", cerr)
		}
	}
	if q.deleteUserRoleStmt != nil {
		if cerr := q.deleteUserRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserRoleStmt: %w", cerr)
		}
	}
	if q.fullTextSearchStmt != nil {
		if cerr := q.fullTextSearchStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing fullTextSearchStmt: %w", cerr)
		}
	}
	if q.getAllRolesStmt != nil {
		if cerr := q.getAllRolesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllRolesStmt: %w", cerr)
		}
	}
	if q.getCategoryByNameStmt != nil {
		if cerr := q.getCategoryByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCategoryByNameStmt: %w", cerr)
		}
	}
	if q.getCollectionStmt != nil {
		if cerr := q.getCollectionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCollectionStmt: %w", cerr)
		}
	}
	if q.getCollectionLastSyncBlockStmt != nil {
		if cerr := q.getCollectionLastSyncBlockStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCollectionLastSyncBlockStmt: %w", cerr)
		}
	}
	if q.getEventStmt != nil {
		if cerr := q.getEventStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEventStmt: %w", cerr)
		}
	}
	if q.getExpiredOrderStmt != nil {
		if cerr := q.getExpiredOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getExpiredOrderStmt: %w", cerr)
		}
	}
	if q.getMarketplaceLastSyncBlockStmt != nil {
		if cerr := q.getMarketplaceLastSyncBlockStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMarketplaceLastSyncBlockStmt: %w", cerr)
		}
	}
	if q.getMarketplaceSettingsStmt != nil {
		if cerr := q.getMarketplaceSettingsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMarketplaceSettingsStmt: %w", cerr)
		}
	}
	if q.getNftStmt != nil {
		if cerr := q.getNftStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNftStmt: %w", cerr)
		}
	}
	if q.getNotificationStmt != nil {
		if cerr := q.getNotificationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNotificationStmt: %w", cerr)
		}
	}
	if q.getOfferStmt != nil {
		if cerr := q.getOfferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOfferStmt: %w", cerr)
		}
	}
	if q.getOrderStmt != nil {
		if cerr := q.getOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOrderStmt: %w", cerr)
		}
	}
	if q.getProfileStmt != nil {
		if cerr := q.getProfileStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProfileStmt: %w", cerr)
		}
	}
	if q.getUserByAddressStmt != nil {
		if cerr := q.getUserByAddressStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByAddressStmt: %w", cerr)
		}
	}
	if q.getUsersStmt != nil {
		if cerr := q.getUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersStmt: %w", cerr)
		}
	}
	if q.getValidMarketplaceSettingsStmt != nil {
		if cerr := q.getValidMarketplaceSettingsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getValidMarketplaceSettingsStmt: %w", cerr)
		}
	}
	if q.insertCategoryStmt != nil {
		if cerr := q.insertCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertCategoryStmt: %w", cerr)
		}
	}
	if q.insertCollectionStmt != nil {
		if cerr := q.insertCollectionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertCollectionStmt: %w", cerr)
		}
	}
	if q.insertEventStmt != nil {
		if cerr := q.insertEventStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertEventStmt: %w", cerr)
		}
	}
	if q.insertMarketplaceSettingsStmt != nil {
		if cerr := q.insertMarketplaceSettingsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertMarketplaceSettingsStmt: %w", cerr)
		}
	}
	if q.insertNotificationStmt != nil {
		if cerr := q.insertNotificationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertNotificationStmt: %w", cerr)
		}
	}
	if q.insertOrderStmt != nil {
		if cerr := q.insertOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertOrderStmt: %w", cerr)
		}
	}
	if q.insertOrderConsiderationItemStmt != nil {
		if cerr := q.insertOrderConsiderationItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertOrderConsiderationItemStmt: %w", cerr)
		}
	}
	if q.insertOrderOfferItemStmt != nil {
		if cerr := q.insertOrderOfferItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertOrderOfferItemStmt: %w", cerr)
		}
	}
	if q.insertUserStmt != nil {
		if cerr := q.insertUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertUserStmt: %w", cerr)
		}
	}
	if q.insertUserRoleStmt != nil {
		if cerr := q.insertUserRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertUserRoleStmt: %w", cerr)
		}
	}
	if q.listNftWithListingStmt != nil {
		if cerr := q.listNftWithListingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listNftWithListingStmt: %w", cerr)
		}
	}
	if q.updateCollectionLastSyncBlockStmt != nil {
		if cerr := q.updateCollectionLastSyncBlockStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCollectionLastSyncBlockStmt: %w", cerr)
		}
	}
	if q.updateMarketplaceLastSyncBlockStmt != nil {
		if cerr := q.updateMarketplaceLastSyncBlockStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMarketplaceLastSyncBlockStmt: %w", cerr)
		}
	}
	if q.updateNftStmt != nil {
		if cerr := q.updateNftStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateNftStmt: %w", cerr)
		}
	}
	if q.updateNonceStmt != nil {
		if cerr := q.updateNonceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateNonceStmt: %w", cerr)
		}
	}
	if q.updateNotificationStmt != nil {
		if cerr := q.updateNotificationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateNotificationStmt: %w", cerr)
		}
	}
	if q.updateOrderStatusStmt != nil {
		if cerr := q.updateOrderStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateOrderStatusStmt: %w", cerr)
		}
	}
	if q.updateOrderStatusByOfferStmt != nil {
		if cerr := q.updateOrderStatusByOfferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateOrderStatusByOfferStmt: %w", cerr)
		}
	}
	if q.updateUserBlockStateStmt != nil {
		if cerr := q.updateUserBlockStateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserBlockStateStmt: %w", cerr)
		}
	}
	if q.upsertNftLatestStmt != nil {
		if cerr := q.upsertNftLatestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertNftLatestStmt: %w", cerr)
		}
	}
	if q.upsertProfileStmt != nil {
		if cerr := q.upsertProfileStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertProfileStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                 DBTX
	tx                                 *sql.Tx
	deleteProfileStmt                  *sql.Stmt
	deleteUserRoleStmt                 *sql.Stmt
	fullTextSearchStmt                 *sql.Stmt
	getAllRolesStmt                    *sql.Stmt
	getCategoryByNameStmt              *sql.Stmt
	getCollectionStmt                  *sql.Stmt
	getCollectionLastSyncBlockStmt     *sql.Stmt
	getEventStmt                       *sql.Stmt
	getExpiredOrderStmt                *sql.Stmt
	getMarketplaceLastSyncBlockStmt    *sql.Stmt
	getMarketplaceSettingsStmt         *sql.Stmt
	getNftStmt                         *sql.Stmt
	getNotificationStmt                *sql.Stmt
	getOfferStmt                       *sql.Stmt
	getOrderStmt                       *sql.Stmt
	getProfileStmt                     *sql.Stmt
	getUserByAddressStmt               *sql.Stmt
	getUsersStmt                       *sql.Stmt
	getValidMarketplaceSettingsStmt    *sql.Stmt
	insertCategoryStmt                 *sql.Stmt
	insertCollectionStmt               *sql.Stmt
	insertEventStmt                    *sql.Stmt
	insertMarketplaceSettingsStmt      *sql.Stmt
	insertNotificationStmt             *sql.Stmt
	insertOrderStmt                    *sql.Stmt
	insertOrderConsiderationItemStmt   *sql.Stmt
	insertOrderOfferItemStmt           *sql.Stmt
	insertUserStmt                     *sql.Stmt
	insertUserRoleStmt                 *sql.Stmt
	listNftWithListingStmt             *sql.Stmt
	updateCollectionLastSyncBlockStmt  *sql.Stmt
	updateMarketplaceLastSyncBlockStmt *sql.Stmt
	updateNftStmt                      *sql.Stmt
	updateNonceStmt                    *sql.Stmt
	updateNotificationStmt             *sql.Stmt
	updateOrderStatusStmt              *sql.Stmt
	updateOrderStatusByOfferStmt       *sql.Stmt
	updateUserBlockStateStmt           *sql.Stmt
	upsertNftLatestStmt                *sql.Stmt
	upsertProfileStmt                  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                 tx,
		tx:                                 tx,
		deleteProfileStmt:                  q.deleteProfileStmt,
		deleteUserRoleStmt:                 q.deleteUserRoleStmt,
		fullTextSearchStmt:                 q.fullTextSearchStmt,
		getAllRolesStmt:                    q.getAllRolesStmt,
		getCategoryByNameStmt:              q.getCategoryByNameStmt,
		getCollectionStmt:                  q.getCollectionStmt,
		getCollectionLastSyncBlockStmt:     q.getCollectionLastSyncBlockStmt,
		getEventStmt:                       q.getEventStmt,
		getExpiredOrderStmt:                q.getExpiredOrderStmt,
		getMarketplaceLastSyncBlockStmt:    q.getMarketplaceLastSyncBlockStmt,
		getMarketplaceSettingsStmt:         q.getMarketplaceSettingsStmt,
		getNftStmt:                         q.getNftStmt,
		getNotificationStmt:                q.getNotificationStmt,
		getOfferStmt:                       q.getOfferStmt,
		getOrderStmt:                       q.getOrderStmt,
		getProfileStmt:                     q.getProfileStmt,
		getUserByAddressStmt:               q.getUserByAddressStmt,
		getUsersStmt:                       q.getUsersStmt,
		getValidMarketplaceSettingsStmt:    q.getValidMarketplaceSettingsStmt,
		insertCategoryStmt:                 q.insertCategoryStmt,
		insertCollectionStmt:               q.insertCollectionStmt,
		insertEventStmt:                    q.insertEventStmt,
		insertMarketplaceSettingsStmt:      q.insertMarketplaceSettingsStmt,
		insertNotificationStmt:             q.insertNotificationStmt,
		insertOrderStmt:                    q.insertOrderStmt,
		insertOrderConsiderationItemStmt:   q.insertOrderConsiderationItemStmt,
		insertOrderOfferItemStmt:           q.insertOrderOfferItemStmt,
		insertUserStmt:                     q.insertUserStmt,
		insertUserRoleStmt:                 q.insertUserRoleStmt,
		listNftWithListingStmt:             q.listNftWithListingStmt,
		updateCollectionLastSyncBlockStmt:  q.updateCollectionLastSyncBlockStmt,
		updateMarketplaceLastSyncBlockStmt: q.updateMarketplaceLastSyncBlockStmt,
		updateNftStmt:                      q.updateNftStmt,
		updateNonceStmt:                    q.updateNonceStmt,
		updateNotificationStmt:             q.updateNotificationStmt,
		updateOrderStatusStmt:              q.updateOrderStatusStmt,
		updateOrderStatusByOfferStmt:       q.updateOrderStatusByOfferStmt,
		updateUserBlockStateStmt:           q.updateUserBlockStateStmt,
		upsertNftLatestStmt:                q.upsertNftLatestStmt,
		upsertProfileStmt:                  q.upsertProfileStmt,
	}
}
