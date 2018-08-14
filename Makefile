.PHONY: z_gen_tables.go
z_gen_tables.go: gen_test.go latlong.go dist/combined-shapefile.shp
	go test --tags=latlong_gen --generate --scale=128 -v