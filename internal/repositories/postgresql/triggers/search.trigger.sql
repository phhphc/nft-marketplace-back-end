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
