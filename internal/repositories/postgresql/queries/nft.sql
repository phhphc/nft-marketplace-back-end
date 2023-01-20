-- name: UpsertNft :exec
INSERT INTO "nfts" (token_id, contract_addr, owner, is_burned, block_number, tx_index)
VALUES ($1,$2,$3,$4, $5, $6)
ON CONFLICT (token_id, contract_addr) DO UPDATE
SET owner=$3,is_burned=$4,block_number=$5, tx_index=$6
WHERE $5 > nfts.block_number or ($5 = nfts.block_number and $6 > nfts.tx_index);

-- name: GetNftByCollection :many
select n.token_id, n.contract_addr, n.owner, n.is_burned, l.listing_id, l.seller, l.price
from "nfts" n
left join "listings" l on n.token_id = l.token_id and n.contract_addr = l.collection
where l.status = 'Open' or l.listing_id is null;