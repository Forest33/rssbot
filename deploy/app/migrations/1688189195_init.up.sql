CREATE TABLE IF NOT EXISTS feeds
(
    id             UUID                        NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    title          VARCHAR(255)                NOT NULL,
    site_url       VARCHAR(1024)               NOT NULL,
    feed_url       VARCHAR(1024)               NOT NULL,
    last_item_hash VARCHAR(32)                 NOT NULL DEFAULT '',
    error_count    INTEGER                     NOT NULL DEFAULT 0,
    created_at     TIMESTAMP WITHOUT TIME ZONE          DEFAULT now() NOT NULL,
    updated_at     TIMESTAMP WITHOUT TIME ZONE NULL
);

CREATE INDEX IF NOT EXISTS "feeds_updated_index" ON feeds (updated_at DESC);
CREATE UNIQUE INDEX IF NOT EXISTS "feeds_url_uindex" ON feeds (feed_url);

INSERT INTO feeds(title, site_url, feed_url)
VALUES ('Go – Компилируемый, многопоточный язык программирования', 'https://habr.com/ru/hub/go/',
        'https://habr.com/ru/rss/hub/go/all/?fl=ru');

CREATE TABLE IF NOT EXISTS users
(
    id         BIGINT                                    NOT NULL PRIMARY KEY,
    first_name VARCHAR(64)                               NOT NULL,
    last_name  VARCHAR(64)                               NOT NULL,
    username   VARCHAR(64)                               NOT NULL,
    language   CHAR(3)                                   NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE               NULL
);

CREATE TABLE IF NOT EXISTS subscriptions
(
    id         UUID   NOT NULL             DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id    BIGINT NOT NULL,
    feed_id    UUID   NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "subscriptions_user_feed_uindex" ON subscriptions (user_id, feed_id);
CREATE INDEX IF NOT EXISTS "subscriptions_feed_index" ON subscriptions (feed_id);

ALTER TABLE feeds
    OWNER TO rssbot;
ALTER TABLE users
    OWNER TO rssbot;
ALTER TABLE subscriptions
    OWNER TO rssbot;