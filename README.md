# go-parking-lot
Golang Parking Lot Automation

## Pendahuluan

Intruksi ini akan berjalan pada local machine untuk development dan bertujuan untuk Testing, dan akan membantu anda untuk memperoleh hasil project tersebut.

### Kebutuhan Aplikasi

Project ini hanya berjalan pada mesin Linux dan menggunakan golang v1.8 dalam pembuatannya.

### Struktur Project Tersebut
1. main.go - Point utama project. <br />
2. parkinglot terdiri dari : 
    - fileInput.go - Contains logic for file based input & (types.go - Basic Datatype for Car - structure) and 1 txt file (sample input for file based input type)
3. command terdiri dari :
    - commandInput.go - Contains logic for the STDIN input type, (fileInput.go - Contains logic for file based input) & (types.go - Basic Datatype for Car - structure) and 1 txt file (sample input for file based input type)
4. bin terdiri dari :
    - (commandInput.go - Contains logic for the STDIN input type) , (fileInput.go - Contains logic for file based input) & (types.go - Basic Datatype for Car - structure) and 1 txt file (sample input for file based input type)
5. functional_spec terdiri dari :
    - (commandInput.go - Contains logic for the STDIN input type) , (fileInput.go - Contains logic for file based input) & (types.go - Basic Datatype for Car - structure) and 1 txt file (sample input for file based input type)
6. input.txt
7. ParkingLot-1.3.1.pdf

### Menjalankan Project (Berbasis Mesin Linuc)
Step 1 - Buka Terminal <br />
Step 2 - Menuju ke direktori project tersimpan <br />
Step 3 - Jalankan go run main.go dan enter masukkan <br />

### TODO
- [x] Development
- [ ] Write unit tests
