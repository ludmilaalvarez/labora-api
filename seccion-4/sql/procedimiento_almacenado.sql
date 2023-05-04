CREATE OR REPLACE FUNCTION calculate_total_price(item_id INTEGER)
RETURNS INTEGER AS $$
BEGIN 
	RETURN (SELECT quantity*price from items where id =item_id);
END;
$$ LANGUAGE plpgsql;

SELECT items.id, items.product, calculate_total_price(items.id) FROM items;
