create table public.items (
	id  bigserial NOT NULL primary key,
	ChrtID     int,
	TrackNumber  varchar(128), 
	Price      int,    
	Rid        varchar(256), 
	Name       varchar(128), 
	Sale       int,    
	Size       varchar(128), 
	TotalPrice int,    
	NmID       int,    
	Brand      varchar(128), 
	Status     int
);

create table public.payment (
	id bigserial NOT NULL primary key,
	Transaction  varchar(256),
	RequestID	 varchar(128),
	Currency     varchar(128), 
	Provider     varchar(128),
	Amount       int    ,
	PaymentDt    int  ,  
	Bank         varchar(128),
	DeliveryCost int   ,
	GoodsTotal   int ,
	CustomFee 	 int
);

create table public.delivery (
	id      bigserial NOT NULL primary key,
	Name       varchar(128),
	Phone 	   varchar(128),
	Zip 		varchar(128),
	City 		varchar(128),
	Address varchar(128),
	Region  varchar(128),
	Email  varchar(128)
);

create table public.orders (
	id bigserial NOT NULL primary key,
	OrderUID          varchar(128), 
	TrackNumber 	  varchar(128),
	Entry             varchar(128),
	delivery_id_fk   	bigserial, 
	payment_id_fk       bigserial,
	items_id_fk         int[],  
	Locale            varchar(128), 
	InternalSignature varchar(128),
	CustomerID        varchar(128),   
	DeliveryService   varchar(128), 
	Shardkey          varchar(128),  
	SmID              int,
	DateCreated		  varchar(128),
	OofShard          varchar(128)
);

create table public.order_items (
	id  bigserial NOT NULL primary key, 
	order_id_fk        bigserial,
	item_id_fk         bigserial
);

create table public.cache (
	id  bigserial NOT NULL primary key, 
	order_id  int, 
	app_key        varchar(128)
);

ALTER TABLE public.orders ADD CONSTRAINT payment_id_fkey FOREIGN KEY (payment_id_fk) REFERENCES public.payment(id) on update no action on delete no action not valid;
ALTER TABLE public.orders ADD CONSTRAINT items_id_fkey FOREIGN KEY (items_id_fk) REFERENCES public.items(id) on update no action on delete no action not valid;
ALTER TABLE public.orders ADD CONSTRAINT delivery_id_fkey FOREIGN KEY (delivery_id_fk) REFERENCES public.delivery(id) on update no action on delete no action not valid;
ALTER TABLE public.order_items ADD CONSTRAINT order_id_fkey FOREIGN KEY (order_id_fk) REFERENCES public.orders(id) match simple on update no action on delete no action not valid;

