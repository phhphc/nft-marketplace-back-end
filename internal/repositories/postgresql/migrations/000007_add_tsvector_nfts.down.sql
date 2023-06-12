ALTER TABLE "nfts"
DROP COLUMN IF EXISTS "tsv";

DROP TRIGGER IF EXISTS "t_nfts_tsvector_trigger" ON "nfts";

DROP FUNCTION IF EXISTS "t_nfts_tsvector_trigger"();