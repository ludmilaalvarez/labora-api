CREATE INDEX idx_items_product ON items (product);

EXPLAIN ANALYZE SELECT * FROM items WHERE product='Computadora';

