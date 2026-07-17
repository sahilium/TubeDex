CREATE TABLE IF NOT EXISTS collection_channels (
    collection_id BIGINT NOT NULL REFERENCES collections(id) ON DELETE CASCADE,
    channel_id    BIGINT NOT NULL REFERENCES channels(id) ON DELETE CASCADE,
    PRIMARY KEY (collection_id, channel_id)
);

CREATE INDEX IF NOT EXISTS idx_collection_channels_collection_id ON collection_channels(collection_id);
CREATE INDEX IF NOT EXISTS idx_collection_channels_channel_id ON collection_channels(channel_id);
