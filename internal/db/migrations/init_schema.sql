CREATE TABLE "products" (
    "id" serial NOT NULL,
    "name" serial NOT NULL,
    "type" TEXT NOT NULL,
    "description" TEXT,
    "price" float8 NOT NULL,
    "img_url" TEXT NOT NULL,
    "additional_info" TEXT,
    CONSTRAINT "products_pk" PRIMARY KEY ("id")
) WITH (
    OIDS=FALSE
);

CREATE TABLE "users" (
     "id" serial NOT NULL,
     "name" TEXT NOT NULL,
     "phone" TEXT NOT NULL,
     "password" TEXT NOT NULL,
     "address" TEXT NOT NULL,
     "sex" TEXT NOT NULL,
     "role" TEXT NOT NULL,
     CONSTRAINT "users_pk" PRIMARY KEY ("id")
) WITH (
     OIDS=FALSE
);

CREATE TABLE "stores" (
      "id" serial NOT NULL,
      "name" TEXT NOT NULL,
      "address" TEXT NOT NULL,
      CONSTRAINT "stores_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
);

CREATE TABLE "orders" (
      "id" serial NOT NULL,
      "courier_id" bigint NOT NULL,
      CONSTRAINT "orders_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
);

CREATE TABLE "order_user" (
      "order_id" bigint NOT NULL,
      "client_id" bigint NOT NULL
) WITH (
      OIDS=FALSE
);

CREATE TABLE "order_product" (
     "order_id" bigint NOT NULL,
     "product_id" bigint NOT NULL
) WITH (
    OIDS=FALSE
);

ALTER TABLE "orders" ADD CONSTRAINT "orders_fk0" FOREIGN KEY ("courier_id") REFERENCES "users"("id");

ALTER TABLE "order_user" ADD CONSTRAINT "order_user_fk0" FOREIGN KEY ("order_id") REFERENCES "orders"("id");
ALTER TABLE "order_user" ADD CONSTRAINT "order_user_fk1" FOREIGN KEY ("client_id") REFERENCES "users"("id");

ALTER TABLE "order_product" ADD CONSTRAINT "order_product_fk0" FOREIGN KEY ("order_id") REFERENCES "orders"("id");
ALTER TABLE "order_product" ADD CONSTRAINT "order_product_fk1" FOREIGN KEY ("product_id") REFERENCES "products"("id");
