ALTER TABLE "nfts" ADD COLUMN tsv TSVECTOR;

UPDATE nfts n1 SET tsv = (
        setweight(to_tsvector(
                          coalesce(n1.metadata->>'name', '')
                      ), 'A') ||
        setweight(to_tsvector(
                          coalesce(n1.metadata->>'description', '')
                      ), 'B')
    )
FROM nfts n2;

DROP INDEX i_nfts_tsv;

CREATE INDEX i_nfts_tsv ON nfts USING gin(tsv);