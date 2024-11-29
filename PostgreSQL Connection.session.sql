SELECT p.product_name, c.category_name FROM tbl_products p JOIN tbl_categories c ON
p.category_id = c.category_id;