package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"

	"github.com/karncomsci/developer-learning-roadmap/api/internal/store"
)

// withTx runs the test inside a transaction that is always rolled back.
// Set TEST_DATABASE_URL (e.g. the local Supabase DB) to run these tests.
func withTx(t *testing.T) pgx.Tx {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	tx, err := conn.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = tx.Rollback(ctx)
		_ = conn.Close(ctx)
	})
	return tx
}

func TestListRolesProgressAndOrder(t *testing.T) {
	tx := withTx(t)
	ctx := context.Background()

	mustExec := func(sql string, args ...any) {
		t.Helper()
		if _, err := tx.Exec(ctx, sql, args...); err != nil {
			t.Fatal(err)
		}
	}
	mustExec(`insert into roles (slug, name, created_at) values
		('it-empty', 'Empty', '2000-01-01'),
		('it-mixed', 'Mixed', '2000-01-02')`)
	mustExec(`insert into topics (role_id, title, status)
		select id, t.title, t.status::learning_status from roles,
		(values ('a', 'done'), ('b', 'learning'), ('c', 'not_started'), ('d', 'done')) as t(title, status)
		where slug = 'it-mixed'`)

	app := NewApp(NewServer(store.New(tx)))
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/v1/roles", nil))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", res.StatusCode)
	}
	var roles []RoleSummary
	if err := json.NewDecoder(res.Body).Decode(&roles); err != nil {
		t.Fatal(err)
	}

	got := map[string]RoleProgress{}
	order := []string{}
	for _, r := range roles {
		got[r.Slug] = r.Progress
		order = append(order, r.Slug)
	}

	// P2: a Role without Topics is 0/0.
	if p := got["it-empty"]; p != (RoleProgress{}) {
		t.Errorf("it-empty progress = %+v, want zeros", p)
	}
	// P1/P3: counts per status.
	if p := got["it-mixed"]; p != (RoleProgress{Done: 2, Learning: 1, Total: 4}) {
		t.Errorf("it-mixed progress = %+v, want done=2 learning=1 total=4", p)
	}
	// R3: oldest first; both test Roles predate any seed data.
	if len(order) < 2 || order[0] != "it-empty" || order[1] != "it-mixed" {
		t.Errorf("order starts %v, want [it-empty it-mixed ...]", order)
	}
}
