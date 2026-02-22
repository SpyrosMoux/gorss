-- Create "feeds" table
CREATE TABLE "public"."feeds" (
  "id" text NOT NULL,
  "name" text NOT NULL,
  "link" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_feeds_link" UNIQUE ("link"),
  CONSTRAINT "uni_feeds_name" UNIQUE ("name")
);
-- Create "articles" table
CREATE TABLE "public"."articles" (
  "id" text NOT NULL,
  "external_id" text NOT NULL,
  "title" text NOT NULL,
  "content" text NOT NULL,
  "link" text NOT NULL,
  "feed_id" text NULL,
  "image_url" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_articles_external_id" UNIQUE ("external_id"),
  CONSTRAINT "uni_articles_link" UNIQUE ("link"),
  CONSTRAINT "fk_feeds_articles" FOREIGN KEY ("feed_id") REFERENCES "public"."feeds" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_articles_feed_id" to table: "articles"
CREATE INDEX "idx_articles_feed_id" ON "public"."articles" ("feed_id");
