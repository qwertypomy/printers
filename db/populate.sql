USE printer_db;
DELETE FROM printer;
DELETE FROM brand;
DELETE FROM printer;
DELETE FROM printing_technology;
DELETE FROM function_type;
DELETE FROM print_size;
DELETE FROM print_resolution;
DELETE FROM printer;
DELETE FROM connectivity_type;
DELETE FROM printer_has_connectivity_type;


insert into printer_db.brand (name, description) values ("Samsung", "Samsung Group — южнокорейская группа компаний, один из крупнейших чеболей, основанный в 1938 году.");
