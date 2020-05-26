CREATE TABLE "Culture" (
	"Id" UUID PRIMARY KEY,
	"Name" VARCHAR(250)
);
CREATE TABLE "Country" (
	"Id" UUID PRIMARY KEY,
    "Name" VARCHAR (250)
);
CREATE TABLE "City" (
	"Id" UUID PRIMARY KEY,
    "Name" VARCHAR (250)
);
CREATE TABLE "Address" (
	"Id" UUID PRIMARY KEY,
    "CountryId" UUID REFERENCES "Country"("Id"),
	"CityId" UUID REFERENCES "City"("Id"),
	"Street" VARCHAR (250),
	"Building" VARCHAR (250),
	"CoordinateX" DECIMAL,
	"CoordinateY" DECIMAL
);
CREATE TABLE "Image" (
	"Id" UUID PRIMARY KEY,
    "Url" VARCHAR (1000)
);
CREATE TABLE "Category" (
	"Id" UUID PRIMARY KEY,
	"Name" VARCHAR(250)
);
CREATE TABLE "CategoryLcz" (
	"Id" UUID PRIMARY KEY,
	"Name" VARCHAR(250),
	"RecordId" UUID REFERENCES "Category"("Id") ON DELETE CASCADE,
	"CultureId" UUID REFERENCES "Culture"("Id")
);
CREATE TABLE "Company" (
	"Id" UUID PRIMARY KEY,
    "Name" VARCHAR (250),
	"AddressId" UUID REFERENCES "Address"("Id"),
	"IconId" UUID REFERENCES "Image"("Id"),
	"OwnerId" UUID REFERENCES "User"("Id"),
	"CategoryId" UUID REFERENCES "Category"("Id")
);
CREATE TABLE "CompanyImage" (
	"CompanyId" UUID REFERENCES "Company"("Id"),
    "ImageId" UUID REFERENCES "Image"("Id"),
	PRIMARY KEY ("CompanyId", "ImageId")
);