ALTER TABLE "nfts" ADD COLUMN tsv TSVECTOR;

-- https://stackoverflow.com/questions/45680936/how-to-implement-full-text-search-on-complex-nested-jsonb-in-postgresql
create or replace function t_nfts_tsvector_trigger() returns trigger as $$
declare
    dict regconfig;
    part_a text;
begin
    dict := 'simple';
    select into part_a string_agg(coalesce(a, ''), ' ') || ' ' || string_agg(coalesce(b, ''), ' ')
    from (
             select
                     new.metadata->>'name',
                     new.metadata->>'description'
         ) as _ (a, b);
    new.tsv := setweight(to_tsvector(dict, part_a), 'A');
    return new;
end;
$$ language plpgsql immutable;

create trigger t_nfts_tsvector_trigger
    before insert or update on nfts for each row execute procedure t_nfts_tsvector_trigger();

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