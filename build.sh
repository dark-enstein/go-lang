#! /bin/bash

echo -------------------- BUILDING CODEGEN --------------------
cd codegen/ ; go build -v ./...
echo
echo -------------------- BUILDING EMAIL --------------------
cd ../email/ ; go build -v ./...
echo
echo -------------------- BUILDING CODEGEN --------------------
cd ../pype/ ; go build -v ./...