CREATE TABLE IF NOT EXISTS channels (
    id                    BIGSERIAL PRIMARY KEY,
    youtube_channel_id    TEXT NOT NULL UNIQUE,
    name                  TEXT NOT NULL,
    handle                TEXT NOT NULL DEFAULT '',
    description           TEXT NOT NULL DEFAULT '',
    avatar                TEXT NOT NULL DEFAULT '',
    banner                TEXT NOT NULL DEFAULT '',
    subscriber_count      BIGINT NOT NULL DEFAULT 0,
    uploads_playlist_id   TEXT NOT NULL DEFAULT '',
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_channels_youtube_id ON channels(youtube_channel_id);
CREATE INDEX IF NOT EXISTS idx_channels_name ON channels(name);
CREATE INDEX IF NOT EXISTS idx_channels_handle ON channels(handle);
