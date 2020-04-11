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
CREATE TABLE "Company" (
	"Id" UUID PRIMARY KEY,
    "Name" VARCHAR (250),
	"AddressId" UUID REFERENCES "Address"("Id"),
	"UserId" UUID REFERENCES "User"("Id")
);