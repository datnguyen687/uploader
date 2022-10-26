DO $$
BEGIN
   FOR cnt IN 1..100 loop
    INSERT INTO products(name, brand, price, created_at)
	VALUES ('test', 'test brand', random(), now());
   END loop;
END; $$