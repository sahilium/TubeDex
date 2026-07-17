CREATE TABLE IF NOT EXISTS videos (
    id                BIGSERIAL PRIMARY KEY,
    youtube_video_id  TEXT NOT NULL UNIQUE,
    channel_id        BIGINT NOT NULL REFERENCES channels(id) ON DELETE CASCADE,
    title             TEXT NOT NULL,
    description       TEXT NOT NULL DEFAULT '',
    published_at      TIMESTAMPTZ NOT NULL,
    thumbnail         TEXT NOT NULL DEFAULT '',
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_videos_channel_id ON videos(channel_id);
CREATE INDEX IF NOT EXISTS idx_videos_published_at ON videos(published_at DESC);
CREATE INDEX IF NOT EXISTS idx_videos_youtube_id ON videos(youtube_video_id);
