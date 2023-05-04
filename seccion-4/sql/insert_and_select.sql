INSERT INTO public.items(
	 customer_name, order_date, product, quantity, price)
	VALUES ('Mauro', Date('03-05-2023'), 'Computadora', 1, 1500),
		   ('Teo', Date('01-04-2023'), 'Sillas', 10, 600),
		   ('Agostina', Date('15-02-2023'), 'Muebles', 2, 850),
		   ('Candela', Date('30-04-2023'), 'Teclado', 1, 30),
		   ('Bautista', Date('20-03-2023'), 'Ventilador', 3, 150);
		   
		   
SELECT *
	FROM public.items
	WHERE quantity > 2 AND price > 50;