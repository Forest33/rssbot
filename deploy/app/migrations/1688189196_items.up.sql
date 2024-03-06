CREATE TABLE IF NOT EXISTS feed_items
(
    id         UUID        NOT NULL        DEFAULT gen_random_uuid() PRIMARY KEY,
    feed_id    UUID        NOT NULL,
    item_hash  VARCHAR(32) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "feed_items_item_uindex" ON feed_items (feed_id, item_hash);
CREATE INDEX IF NOT EXISTS "feed_items_created_index" ON feed_items (created_at);

ALTER TABLE feed_items
    OWNER TO rssbot;
