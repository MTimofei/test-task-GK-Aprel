package postgres

const (
	queryCheck = `
	DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_tables WHERE tablename = 'users') THEN
        CREATE TABLE users (
            id UUID PRIMARY KEY,
            login VARCHAR(255) NOT NULL UNIQUE,
            password_hash VARCHAR(255) NOT NULL,
            login_attempts INT NOT NULL DEFAULT 0
        );
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_tables WHERE tablename = 'sessions') THEN
        CREATE TABLE sessions (
            user_id UUID NOT NULL REFERENCES users(id),
            rev_token TEXT NOT NULL UNIQUE,
            generated_at TIMESTAMP NOT NULL,
            expiration_interval INTERVAL NOT NULL,
            expiration_at TIMESTAMP NOT NULL
        );
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_tables WHERE tablename = 'auth_audit') THEN
        CREATE TABLE auth_audit (
            user_id UUID NOT NULL REFERENCES users(id),
            event_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            event_type VARCHAR(255) NOT NULL
        );
    END IF;
	END $$;`

	queryAuth = `
	WITH updated_users AS (
		UPDATE users 
		SET login_attempts = CASE 
			WHEN login_attempts >= 5 THEN 5
			WHEN password_hash != md5('%s') THEN login_attempts + 1 
			ELSE 0 
			END
		WHERE login = '%s'
		RETURNING id ,CASE 
			WHEN login_attempts >= 5 THEN 'block'
			WHEN password_hash != md5('%s') THEN 'invalid password'
			ELSE  'successful login'
			END AS result
	)
	INSERT INTO auth_audit (user_id, event_type)
	SELECT id, result
	FROM updated_users
	RETURNING user_id, event_type;
	`

	queryTransaction = `
	BEGIN;
	SELECT * FROM users WHERE login = '%s' FOR UPDATE;
	`
	queryCommit = `COMMIT;`

	queryROLLBACK = `ROLLBACK;`

	querySetToken = `
		INSERT INTO sessions (user_id, rev_token, generated_at, expiration_interval, expiration_at)
		VALUES ('%s', '%s', 
		CURRENT_TIMESTAMP, 
		INTERVAL '1 hour', 
		CURRENT_TIMESTAMP + INTERVAL '1 hour'
		);`

	queryAudit = `
	SELECT aa.event_time, aa.event_type
	FROM auth_audit aa
	JOIN sessions s ON aa.user_id = s.user_id
	WHERE EXISTS (
		SELECT 1
		FROM sessions
		WHERE rev_token = '%s'
			AND expiration_at > CURRENT_TIMESTAMP
		);`

	queryClearAudit = ` 
	DELETE FROM auth_audit
    WHERE user_id IN (
        SELECT aa.user_id
        FROM auth_audit aa
        JOIN sessions s ON aa.user_id = s.user_id
        WHERE EXISTS (
            SELECT 1
            FROM sessions
            WHERE rev_token = '%s'
                AND expiration_at > CURRENT_TIMESTAMP
        )
    )
    returning	event_time	,event_type ;`
)
